package rpc

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
