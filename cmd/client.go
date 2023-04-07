package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"path"
	"strings"
)

// Command line flags

var actions map[string]func()
var client *rpc.Client

var (
	network string = "unix" // TODO : add multiple networks
	address string = "/tmp/rpc.sock"
	action  string = "add-top"
	data    string = ""
)

func init() {

	actions = map[string]func(){
		"add":           addTop,
		"add-top":       addTop,
		"add-bottom":    addBottom,
		"get":           getTop,
		"get-top":       getTop,
		"get-bottom":    getBottom,
		"remove":        func() {},
		"remove-top":    removeTop,
		"remove-bottom": removeBottom,
	}

	flag.StringVar(&network, "network", network, "network type : default \"unix\"")
	flag.StringVar(&address, "address", address, "path to network socket : default \"/tmp/rpc.sock\"")
	flag.StringVar(&action, "action", action, "action to do : default \"add-top\"")

	flag.Usage = usage
	flag.Parse()

	data = flag.Arg(0)

	c, err := rpc.DialHTTP(network, address)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	client = c

}

func main() {

	handler, ok := actions[action]

	if !ok {
		fmt.Fprintf(os.Stderr, "action %s is not valid", action)
		return
	}

	handler()

}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage : %s [flags] [data]\n", path.Base(os.Args[0]))

	flag.VisitAll(func(fg *flag.Flag) {
		// Don't let users know about flags they shouldn't use.
		if fg.Name == "network" {
			return
		}
		fmt.Printf("--%s\n\t%s\n", fg.Name,
			strings.Replace(fg.Usage, "\n", "\n\t", -1))
	})

	fmt.Println(
		"Actions :",
		"add", "add-top", "add-bottom",
		"get", "get-top", "get-bottom",
		"remove", "remove-top", "remove-bottom",
		"from-clipboard", "to-clipboard",
	)

	os.Exit(1)
}

func addTop() {
	var reply string
	err := client.Call("Clip.AddTop", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}
}

func addBottom() {
	var reply string
	err := client.Call("Clip.AddBottom", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}
}

func getTop() {
	var reply string
	err := client.Call("Clip.GetTop", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}
	fmt.Fprintln(os.Stdout, reply)

}

func getBottom() {
	var reply string
	err := client.Call("Clip.GetBottom", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}
	fmt.Fprintln(os.Stdout, reply)
}

func removeTop() {
	var reply string
	err := client.Call("Clip.RemoveTop", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}

}

func removeBottom() {
	var reply string
	err := client.Call("Clip.RemoveBottom", &data, &reply)
	if err != nil {
		log.Fatal("clip error:", err)
	}

}
