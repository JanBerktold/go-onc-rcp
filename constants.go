package rpc

import (
	"errors"
)

type status interface{}

type callType uint32
type msgStatus uint32
type acceptStatus uint32
type rejectStatus uint32
type authFlavor uint32

const (
	typeCall callType = iota
	typeReply

	msgAccepted msgStatus = iota
	msgDenied

	acceptSucess acceptStatus = iota
	acceptProgramUnavailable
	acceptProgramMismatch
	acceptProcessUnavailable
	acceptGarbageArguments
	acceptSystemError

	rejectRPCMismatch rejectStatus = iota
	rejectAuthenticationError

	authNone authFlavor = iota
	authSystem
	authShort
	authDH
	authRPCSECGSS
)

type success struct {
	Payload []byte
}

type programMismatch struct {
	low, high uint32
}

type programUnavailable struct{}
type processUnavailable struct{}
type garbargeArguments struct{}
type systemError struct{}

type rpcMismatch struct {
	low, high uint32
}

type authError struct{}

var (
	ErrProgramMismatch   = errors.New("Program Version mismatch")
	ErrProgramUnavilable = errors.New("Program not avilable")
	ErrProcessUnavilable = errors.New("Process not available")
	ErrGarbargeArguments = errors.New("Arguments not interpreted correctly")
	ErrSystemError       = errors.New("Generic system error")
	ErrRPCMismatch       = errors.New("Wrong RPC version")
	ErrAuthError         = errors.New("Authentication error")
	ErrUnknownError      = errors.New("Server returned unknown error")
)

func toError(st status) error {
	switch st.(type) {
	case programMismatch:
		return ErrProgramMismatch
	case programUnavailable:
		return ErrProgramUnavilable
	case processUnavailable:
		return ErrProcessUnavilable
	case garbargeArguments:
		return ErrGarbargeArguments
	case systemError:
		return ErrSystemError
	case rpcMismatch:
		return ErrRPCMismatch
	case authError:
		return ErrAuthError
	}
	return ErrUnknownError
}
