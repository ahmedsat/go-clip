package client

import "log"

func RemoveTop(network string, address string, args []string) {
	conn := conn(network, address)
	var reply string
	data := ""

	err := conn.Call("Clip.RemoveTop", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}

}

func RemoveBottom(network string, address string, args []string) {
	conn := conn(network, address)
	var reply string
	data := ""

	err := conn.Call("Clip.RemoveBottom", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}

}
