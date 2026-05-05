package smolp

import (
	"fmt"
	"slices"
)

type Field struct {
	Name  string
	Type  string
	Value any
}

type Payload struct {
	Fields []Field
}

func NewPayload() (*Payload, error) {
	return nil, nil
}

func (payload *Payload) Get(name string) (any, error) {
	idx := slices.IndexFunc(payload.Fields, func(f Field) bool {
		return f.Name == name
	})

	if idx == -1 {
		return nil, fmt.Errorf("field with name %s not found", name)
	}

	return payload.Fields[idx].Value, nil
}