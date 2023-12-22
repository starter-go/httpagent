package httpagent

import (
	"fmt"
	"io"
)

// Response ...
type Response struct {
	Request *Request
	Status  int
	Message string
	Headers Headers
	Body    io.ReadCloser
}

// GetEntity 以 Entity 的形式获取响应的 body
func (inst *Response) GetEntity() (*Entity, error) {

	src := inst.Body
	inst.Body = nil
	if src == nil {
		return nil, fmt.Errorf("no body")
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	length := len(data)

	entity := new(Entity)
	entity.Data = data
	entity.ContentType = inst.Headers.Get("Content-Type")
	entity.ContentLength = int64(length)
	return entity, nil
}
