package rpc

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type call struct {
	XId        uint32
	RPCVersion uint32 /* must be equal to two (2) */
	Program    Program
	Process    uint32
	auth       uint32
	byteStream bool
}

func (c call) recordMarking(length uint32) uint32 {
	return (uint32(1) << 31) | length
}

func (c call) Seralize() ([]byte, error) {
	rd := bytes.NewBuffer(nil)

	if c.byteStream {
		if err := binary.Write(rd, binary.BigEndian, c.recordMarking(40)); err != nil {
			return nil, err
		}
	}

	if err := binary.Write(rd, binary.BigEndian, c.XId); err != nil {
		return nil, err
	}

	if err := binary.Write(rd, binary.BigEndian, type_call); err != nil {
		return nil, err
	}

	if err := binary.Write(rd, binary.BigEndian, c.RPCVersion); err != nil {
		return nil, err
	}

	if err := binary.Write(rd, binary.BigEndian, c.Program.Number); err != nil {
		return nil, err
	}

	if err := binary.Write(rd, binary.BigEndian, c.Program.Version); err != nil {
		return nil, err
	}

	if err := binary.Write(rd, binary.BigEndian, c.Process); err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", c)
	for i := 0; i < 4; i++ {
		if err := binary.Write(rd, binary.BigEndian, c.auth); err != nil {
			return nil, err
		}
	}

	return rd.Bytes(), nil
}
