package rpc

import (
	"bytes"
	"encoding/binary"
	"github.com/davecgh/go-xdr/xdr2"
	"io"
	"log"
)

type reply struct {
	XId    uint32
	Status status
}

func (r reply) Seralize(byteStream bool) ([]byte, error) {
	w := bytes.Buffer{}
	enc := xdr.NewEncoder(&w)

	if byteStream {

	}

	if _, err := enc.EncodeUint(r.XId); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func parseVerification(reader io.Reader) error {
	stat := authFlavor(0)

	if err := binary.Read(reader, binary.BigEndian, &stat); err != nil {
		return err
	}

	switch stat {
	case authNone:
		log.Println("NONE")
	default:
		log.Fatalln(stat)
	}

	return nil
}

func parseReply(buffer []byte, byteStream bool) (reply, error) {
	result := reply{}
	reader := bytes.NewBuffer(buffer)
	log.Println(buffer)

	if byteStream {
		waste := uint32(0)
		if err := binary.Read(reader, binary.BigEndian, &waste); err != nil {
			return result, err
		}
	}

	if err := binary.Read(reader, binary.BigEndian, &result.XId); err != nil {
		return result, err
	}

	mtype := callType(0)
	if err := binary.Read(reader, binary.BigEndian, &mtype); err != nil {
		return result, err
	}

	if mtype != typeReply {
		log.Println("EXPECTED REPLY")
		return result, nil
	}

	// REPLY STATUS
	replyStatus := msgStatus(0)
	if err := binary.Read(reader, binary.BigEndian, &replyStatus); err != nil {
		return result, err
	}

	log.Println("REPLY STATUS", replyStatus == msgAccepted)
	switch replyStatus == msgAccepted {
	case true:
		if err := parseVerification(reader); err != nil {
			return result, err
		}

		specifiedStatus := acceptStatus(0)
		if err := binary.Read(reader, binary.BigEndian, &specifiedStatus); err != nil {
			return result, err
		}

		switch specifiedStatus {
		case acceptSucess:
			success := success{
				Payload: make([]byte, reader.Len()),
			}
			result.Status = success
			if err := binary.Read(reader, binary.BigEndian, &success.Payload); err != nil {
				return result, err
			}
		case acceptProgramMismatch:
			result.Status = programMismatch{}
		case acceptProgramUnavailable:
			result.Status = programUnavailable{}
		case acceptProcessUnavailable:
			result.Status = processUnavailable{}
		case acceptGarbageArguments:
			result.Status = garbargeArguments{}
		case acceptSystemError:
			result.Status = systemError{}
		default:
			log.Fatal("UNKNOWN SUCCESS")
		}
	case false:
		specifiedStatus := rejectStatus(0)
		if err := binary.Read(reader, binary.BigEndian, &specifiedStatus); err != nil {
			return result, err
		}

		switch specifiedStatus {
		case rejectRPCMismatch:
			result.Status = rpcMismatch{}
		case rejectAuthenticationError:
		default:
			log.Fatal("UNKNOWN FAILURE")
		}
	}

	return result, nil
}
