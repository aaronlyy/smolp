package smolp



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
