package main

import (
	"log"

	"github.com/aaronlyy/smolp"
)

func handler(payload *smolp.Payload) {
	log.Println(payload.Get("id"))
}

func main() {

	// create field definition froms string
	fields, err := smolp.FieldDefsFromString("id:uint16,value:uint32")

	// create a new protocol definition
	protocol := smolp.NewProtocol("testp", 1, smolp.BaseUDP, fields)
	if err != nil {
		log.Fatal(err)
	}

	// create new server with handler
	server, err := smolp.NewServer(protocol, handler)
	if err != nil {
		log.Fatal(err)
	}

	// start server
	server.Listen("localhost", 4444)
}
