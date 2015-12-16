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
	stat := auth_flavor(0)

	if err := binary.Read(reader, binary.BigEndian, &stat); err != nil {
		return err
	}

	switch stat {
	case auth_none:
		log.Println("NONE")
	default:
		log.Fatalln(stat)
	}

	return nil
}

func parseReply(buffer []byte, byteStream bool) (reply, error) {
	result := reply{}
	reader := bytes.NewBuffer(buffer)

	if byteStream {
		waste := uint32(0)
		if err := binary.Read(reader, binary.BigEndian, &waste); err != nil {
			return result, err
		}
	}

	if err := binary.Read(reader, binary.BigEndian, &result.XId); err != nil {
		return result, err
	}

	mtype := call_type(0)
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

	switch reply_stat == msg_accepted {
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
		case accept_prog_unavail:
			result.Status = programUnavailable{}
		case accept_proc_unavail:
			result.Status = processUnavailable{}
		case accept_garbage_args:
			result.Status = garbargeArgs{}
		case accept_system_err:
			result.Status = systemError{}
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
