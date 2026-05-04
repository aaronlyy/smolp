package smolp

// a protocol needs a header definition and a payload definition

type Protocol struct {
	Header  Header
	Payload Payload
}

func NewProtocol(header *Header, payload *Payload) (*Protocol, error) {
	return nil, nil
}

// decode only header bytes an return struct with data
func (p *Protocol) DecodeHeader(bytes []byte) (*Header, error) {
	return nil, nil
}

// decode bytes after header and return struct with data
func (p *Protocol) DecodePayload(bytes []byte) (*Payload, error) {
	return nil, nil
}

// decode full packet and return struct with data
func (p *Protocol) DecodePacket(bytes []byte) (*Header, *Payload, error) {
	return nil, nil, nil
}
