package smolp

var ValidTypes = map[string]int{
    "float32": 4,
    "float64": 8,
    "uint8":   1,
    "uint16":  2,
    "uint32":  4,
    "uint64":  8,
    "int8":    1,
    "int16":   2,
    "int32":   4,
    "int64":   8,
}

type Field struct {
	Name  string
	Type  string
	Value any
}

type Payload struct {
	Fields []Field
}
