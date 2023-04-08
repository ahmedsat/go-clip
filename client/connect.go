package client

import (
	"fmt"
	"log"
	"net/rpc"
)

func conn(network string, address string) *rpc.Client {
	c, err := rpc.DialHTTP(network, address)
	if err != nil {
		fmt.Printf("Cant not connect to { %s %s } check up is server if server is running\n", network, address)
		log.Fatal("dialing:", err)
	}

	return c
}
