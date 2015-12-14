package rpc

type status interface{}

type call_type uint32

const (
	type_call call_type = iota
	type_reply
)

type msg_stat uint32

const (
	msg_accepted msg_stat = iota
	msg_denied
)

type accept_stat uint32

const (
	accept_sucess accept_stat = iota
	accept_prog_unavail
	accept_prog_mismatch
	accept_proc_unavail
	accept_garbage_args
	accept_system_err
)

type success struct {
	Payload []byte
}

type programMismatch struct {
	low, high uint32
}

type programUnavailable struct{}
type processUnavailable struct{}
type garbargeArgs struct{}
type systemError struct{}

// reasons for reject of call
type reject_stat uint32

const (
	reject_rpc_mismatch reject_stat = iota
	reject_auth_error
)

type rpcMismatch struct {
	low, high uint32
}

type authError struct{}

type auth_flavor uint32

const (
	auth_none auth_flavor = iota
	auth_sys
	auth_short
	auth_dh
	auth_rpcsec_gss
)
