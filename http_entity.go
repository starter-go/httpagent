package httpagent

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/starter-go/vlog"
)

// Entity ...
type Entity struct {
	ContentLength int64
	ContentType   string
	Data          []byte
}

// OpenReader ...
func (inst *Entity) OpenReader() (io.ReadCloser, error) {
	r := bytes.NewReader(inst.Data)
	return io.NopCloser(r), nil
}

// ReadBinary ...
func (inst *Entity) ReadBinary() ([]byte, error) {
	r, err := inst.OpenReader()
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}

// ReadText ...
func (inst *Entity) ReadText() (string, error) {
	data, err := inst.ReadBinary()
	if err != nil {
		return "", err
	}
	return string(data), err
}

// ReadJSON ...
func (inst *Entity) ReadJSON(obj any) error {
	data, err := inst.ReadBinary()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}

// NewEntityWithBinary ...
func NewEntityWithBinary(bin []byte, contentType string) *Entity {
	ent := new(Entity)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	length := len(bin)
	ent.ContentLength = int64(length)
	ent.ContentType = contentType
	ent.Data = bin
	return ent
}

// NewEntityWithJSON ...
func NewEntityWithJSON(v any) *Entity {
	ent := new(Entity)
	data, err := json.Marshal(v)
	if err == nil {
		length := len(data)
		ent.ContentLength = int64(length)
		ent.ContentType = "application/json;charset=utf-8"
		ent.Data = data
	} else {
		vlog.Warn(err.Error())
	}
	return ent
}

// NewEntityWithText ...
func NewEntityWithText(text string, contentType string) *Entity {
	ent := new(Entity)
	if contentType == "" {
		contentType = "text/plain;charset=utf-8"
	}
	data := []byte(text)
	length := len(data)
	ent.ContentLength = int64(length)
	ent.ContentType = contentType
	ent.Data = data
	return ent
}
