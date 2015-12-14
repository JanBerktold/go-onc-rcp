package rpc

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

type reply struct {
	XId      uint32
	Accepted bool
	Status   status
}

func (r reply) Seralize() ([]byte, error) {
	return nil, nil
}

func parseVerification(reader io.Reader) error {
	stat := auth_flavor(0)

	if err := binary.Read(reader, binary.BigEndian, &stat); err != nil {
		return err
	}

	switch stat {
	case auth_none:
		log.Println("NONE")
	default:
		log.Println(stat)
	}

	return nil
}

func parseReply(buffer []byte, byteStream bool) (reply, error) {
	result := reply{}
	reader := bytes.NewBuffer(buffer)

	waste := uint32(0)
	mtype := call_type(0)

	if byteStream {
		if err := binary.Read(reader, binary.BigEndian, &waste); err != nil {
			return result, err
		}
	}

	if err := binary.Read(reader, binary.BigEndian, &result.XId); err != nil {
		return result, err
	}

	if err := binary.Read(reader, binary.BigEndian, &mtype); err != nil {
		return result, err
	}

	if mtype != type_reply {
		log.Println("EXPECTED REPLY")
		return result, nil
	}

	// REPLY STATUS
	reply_stat := msg_stat(0)
	if err := binary.Read(reader, binary.BigEndian, &reply_stat); err != nil {
		return result, err
	}
	result.Accepted = reply_stat == msg_accepted

	switch result.Accepted {
	case true:
		if err := parseVerification(reader); err != nil {
			return result, err
		}

		accept_status := accept_stat(0)
		if err := binary.Read(reader, binary.BigEndian, &accept_status); err != nil {
			return result, err
		}

		switch accept_status {
		case accept_sucess:
			success := success{
				Payload: make([]byte, reader.Len()),
			}
			result.Status = success
			if err := binary.Read(reader, binary.BigEndian, &success.Payload); err != nil {
				return result, err
			}
		case accept_prog_mismatch:
			result.Status = programMismatch{}
		}
	case false:
		reject_status := reject_stat(0)
		if err := binary.Read(reader, binary.BigEndian, &reject_status); err != nil {
			return result, err
		}

		switch reject_status {
		case reject_rpc_mismatch:
			result.Status = rpcMismatch{}
		case reject_auth_error:
		}
	}

	return result, nil
}
