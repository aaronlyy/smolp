package smolp

type Header struct {
	MacicBytes    uint32 // 4 bytes
	Version       uint8  // 1 byte
	SeqNr         uint16 // 2 bytes
	Timestamp     uint64 // 8 bytes
	PayloadLength uint16 // 2 bytes
	Checksum      uint32 // 4 bytes
}

func EncodeHeader(header *Header) ([]byte, error) {
	return nil, nil
}

func DecodeHeader(b []byte) (*Header, error) {
	return nil, nil
}
