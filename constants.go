package rpc

type CALL_TYPE uint32

const (
	CALL CALL_TYPE = iota
	REPLY
)
