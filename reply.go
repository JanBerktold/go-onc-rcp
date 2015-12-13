package rpc

type reply struct {
	XId uint32
}

func (r reply) Seralize() ([]byte, error) {
	return nil, nil
}

func parseReply(buffer []byte) (reply, error) {
	result := reply{}

	return result
}
