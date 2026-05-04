package main

import (
	"fmt"
	"log"

	"github.com/aaronlyy/smolp"
)

func main() {

	// create field definition froms string
	fields, err := smolp.FieldDefsFromString("id:uint16,value:uint32")

	// create a new protocol definition
	protocol := smolp.NewProtocol("testp", 1, smolp.BaseUDP, fields)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", protocol)
}
