package core

import (
	"bytes"
	"encoding/gob"
)

type Row struct {
	ID       uint32
	UserName string
	Email    string
}

func (r *Row) Serialize() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(r); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (r *Row) Deserialize(data []byte) error {
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	if err := dec.Decode(&r); err != nil {
		return err
	}
	return nil
}
