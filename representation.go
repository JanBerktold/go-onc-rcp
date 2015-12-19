package rpc

import (
	"bytes"
	"github.com/davecgh/go-xdr/xdr2"
)

// Handles representing either a byte slice or a pointer to a struct which can either be
// written to or interpreted as a byte slice. Used for handling parameter and reply payload
// data within the call structure.
type dataRepresentation interface {
	Len() (int, error)
}

type bytesTarget struct {
	Target []byte
}

func (b bytesTarget) Len() (int, error) {
	return len(b.Target), nil
}

type structTarget struct {
	Target        interface{}
	SeralizedData []byte
}

func (s structTarget) Seralize() error {
	buffer := bytes.NewBuffer(nil)
	_, err := xdr.Marshal(buffer, s.Target)
	if err != nil {
		return err
	}
	s.SeralizedData = buffer.Bytes()
	return nil
}

func (s structTarget) Len() (int, error) {
	if len(s.SeralizedData) == 0 {
		err := s.Seralize()
		if err != nil {
			return 0, err
		}
	}
	return len(s.SeralizedData), nil
}
