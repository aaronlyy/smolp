package main

import (
	"log"

	"github.com/aaronlyy/smolp"
)

func main() {

	// define how your protocol should look like
	header, err := smolp.NewHeader(3402, 1)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := smolp.NewPayload()
	if err != nil {
		log.Fatal(err)
	}

	protocol, err := smolp.NewProtocol(header, payload)
	if err != nil {
		log.Fatal(err)
	}

	server, err := smolp.NewServer(protocol)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Listen("localhost", 3000); err != nil {
		log.Fatal(err)
	}
}