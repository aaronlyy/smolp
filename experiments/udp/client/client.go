package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

type Header struct {
	MagicBytes           uint32 // 4 bytes - identifies this as a smolp packet
	Version              uint8  // 1 byte  - protocol version
	SeqNr                uint16 // 2 bytes - sequence number for ordering
	Timestamp            uint64 // 8 bytes - unix timestamp in milliseconds
	PayloadLengthInBytes uint16 // 2 bytes - length of the payload that follows
}

func encodeHeader(header *Header) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, header); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	addr := "localhost:9000"

	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving address: %v\n", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	var seqNr uint16 = 0

	for {
		header := Header{
			MagicBytes:           0x534D4F4C,                     // "SMOL" in ASCII hex
			Version:              1,                              // protocol version 1
			SeqNr:                seqNr,                          // current sequence number
			Timestamp:            uint64(time.Now().UnixMilli()), // current time in ms
			PayloadLengthInBytes: 0,                              // no payload
		}

		headerBytes, err := encodeHeader(&header)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[udp-client] encode error: %v\n", err)
			continue
		}

		_, err = conn.Write(headerBytes)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[udp-client] send error: %v\n", err)
			continue
		}

		fmt.Printf("[udp-client] sent seq=%d\n", seqNr)

		seqNr++

		time.Sleep(time.Second)
	}
}
