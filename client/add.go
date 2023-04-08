package client

import (
	"log"
)

func AddTop(network string, address string, args []string) {

	conn := conn(network, address)
	var reply string

	data := ""
	for _, v := range args {
		data += v
	}

	err := conn.Call("Clip.AddTop", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}

}

func AddBottom(network string, address string, args []string) {

	conn := conn(network, address)
	var reply string

	data := ""
	for _, v := range args {
		data += v
	}

	err := conn.Call("Clip.AddBottom", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}

}
