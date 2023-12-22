package httpagent

import (
	"bytes"
	"encoding/json"
	"io"
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
