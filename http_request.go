package httpagent

import (
	"bytes"
	"context"
	"io"
	"strconv"
)

// Request ...
type Request struct {
	Context context.Context
	Method  string
	URL     string
	Headers Headers
	Body    io.ReadCloser
}

// SetEntity 设置请求的 body
func (inst *Request) SetEntity(entity *Entity) {
	if entity == nil {
		return
	}

	length := len(entity.Data)
	body := bytes.NewReader(entity.Data)

	inst.Headers.Set("Content-Length", strconv.Itoa(length))
	inst.Headers.Set("Content-Type", entity.ContentType)
	inst.Body = io.NopCloser(body)
}
