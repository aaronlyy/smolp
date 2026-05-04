package smolp

import (
	"fmt"
	"strings"
)

// only defines how a packet should look like, it never holds any real data
// it can be used to decode bytes into header and payload structs that contain the real data

var ValidFieldTypes = map[string]int{
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

type FieldDef struct {
	Name string
	Type string
}

type Protocol struct {
	Name      string // gets turned into magic bytes
	Version   uint8  // used to match packets
	Base      string // udp or tcp
	FieldDefs []FieldDef
}

// Parses a string to a field definition slice.
// The string should be a comma seperated list of colon seperated field definitions. Example: "sensorId:uint32,sensorType:uint16,sensoData:uint64"
func FieldDefsFromString(s string) ([]FieldDef, error) {
	var fieldDefs []FieldDef
	for field := range strings.SplitSeq(s, ",") {
		fieldSplit := strings.Split(field, ":")
		if len(fieldSplit) != 2 {
			return nil, fmt.Errorf("invalid field definition %q", field)
		}
		if _, ok := ValidFieldTypes[fieldSplit[1]]; !ok {
			return nil, fmt.Errorf("unsupported type in field definition: %s:%s <--", fieldSplit[0], fieldSplit[1])
		}
		fieldDefs = append(fieldDefs, FieldDef{
			Name: fieldSplit[0],
			Type: fieldSplit[1],
		})
	}
	return fieldDefs, nil
}

// Returns a new protocol definition
func NewProtocol(name string, version uint8, base string, fieldDefs []FieldDef) *Protocol {
	return &Protocol{
		Name:      name,
		Version:   version,
		Base:      base,
		FieldDefs: fieldDefs,
	}
}

func NewProtocolFromYAMLString(s string) (*Protocol, error) {
	// handle yaml errors
	return nil, nil
}

func ProtocolToYAMLString(path string, protocol *Protocol) (string, error) {
	// handle errors while writing yaml
	return "", nil
}

// Takes in a pointer to a header and encodes it to an byte array using encoding/binary
func (p *Protocol) EncodeHeader(header *Header) ([]byte, error) {
	// check if given header matches header definition in protocol
	return nil, nil
}

// Takes in a pointer to a payload and encodes it to an byte array using encoding/binary
func (p *Protocol) EncodePayload(payload *Payload) ([]byte, error) {
	// check if given payload matches payload definition in protocol
	return nil, nil
}

// Takes in a pointer to a header and a payload and encodes it to an byte array using encoding/binary
func (p *Protocol) Encode(header *Header, payload *Payload) ([]byte, []byte, error) {
	headerBytes, err := p.EncodeHeader(header)
	if err != nil {
		return nil, nil, err
	}

	payloadBytes, err := p.EncodePayload(payload)
	if err != nil {
		return nil, nil, err
	}

	return headerBytes, payloadBytes, nil
}

// Decode header bytes and return struct with data
func (p *Protocol) DecodeHeader() (*Header, error) {
	return nil, nil
}

// Decode payload bytes and return struct with data
func (p *Protocol) DecodePayload() (*Payload, error) {
	return nil, nil
}

// Decodes both the header and the payload and returns two structs
func (p *Protocol) Decode() (*Header, *Payload) {
	return nil, nil
}
