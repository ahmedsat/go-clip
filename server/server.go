package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/ahmedsat/go-clip/handlers"
)

// const SockAddr = "/tmp/rpc.sock"

func Serve(SockAddr string) {

	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	clip := new(handlers.Clip)
	rpc.Register(clip)
	rpc.HandleHTTP()
	l, e := net.Listen("unix", SockAddr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("Serving...")
	http.Serve(l, nil)
}
