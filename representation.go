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

type bytesData struct {
	Data []byte
}

func (b bytesData) Len() (int, error) {
	return len(b.Data), nil
}

// Reprents a pointer to a struct which is used as data. Also keeps a local cache of the
// seralized struct if used as data source.
type structData struct {
	Data          interface{}
	SeralizedData []byte
}

func (s structData) Seralize() error {
	buffer := bytes.NewBuffer(nil)
	_, err := xdr.Marshal(buffer, s.Data)
	if err != nil {
		return err
	}
	s.SeralizedData = buffer.Bytes()
	return nil
}

func (s structData) Len() (int, error) {
	if len(s.SeralizedData) == 0 {
		err := s.Seralize()
		if err != nil {
			return 0, err
		}
	}
	return len(s.SeralizedData), nil
}
