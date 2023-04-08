package client

import (
	"fmt"
	"log"
	"os"
)

func GetTop(network string, address string, args []string) {
	conn := conn(network, address)
	var reply string
	data := ""

	err := conn.Call("Clip.GetTop", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}
	fmt.Fprintln(os.Stdout, reply)

}

func GetBottom(network string, address string, args []string) {
	conn := conn(network, address)
	var reply string
	data := ""

	err := conn.Call("Clip.GetBottom", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}
	fmt.Fprintln(os.Stdout, reply)
}
