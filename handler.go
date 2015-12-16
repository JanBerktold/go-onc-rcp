package rpc

import (
	"io"
)

type Handler func(io.Writer)
