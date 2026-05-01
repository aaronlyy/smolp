package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

type Header struct {
	MacicBytes           uint32 // 4 bytes
	Version              uint8  // 1 byte
	SeqNr                uint16 // 2 bytes
	Timestamp            uint64 // 8 bytes
	PayloadLengthInBytes uint16 // 2 bytes
}

func decodeHeader(b []byte) (*Header, error) {
	var header Header
	reader := bytes.NewReader(b)
	if err := binary.Read(reader, binary.BigEndian, &header); err != nil {
		return nil, err
	}
	return &header, nil
}

// exmaple handler, accept data and handle things
func handle(header *Header) {
	fmt.Printf("%v\n", header)
}

func main() {
	addr := "localhost:9000"

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving address: %v\n", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listening: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Printf("[udp-server] listening on %s\n", addr)

	buf := make([]byte, 17)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("[udp-server] read error: %v\n", err)
			continue
		}

		header, err := decodeHeader(buf[:n])
		if err != nil {
			fmt.Println("error while decoding header")
			continue
		}

		go handle(header)

	}
}
