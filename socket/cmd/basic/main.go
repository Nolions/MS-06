package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"socketlab/server/basic"
)

var (
	port string
)

func main() {
	flag.StringVar(&port, "p", "5489", "Socket Sevicer's port")
	flag.Parse()

	if len(flag.Args()) < 1 {
		return
	}

	run(flag.Args()[0])
}

func run(t string) {
	if t == "server" {
		server()
	} else if t == "client" {
		client()
	}
}

func server() {
	// build and listen
	listener, _ := net.Listen("tcp", fmt.Sprintf(":%s", port))
	log.Printf("socker server run on %s", port)

	for {
		// accept
		conn, _ := listener.Accept()
		basic.ServerConnHandler(conn)
	}
}

func client() {
	// connection
	conn, _ := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%s", port))

	fmt.Println("Connect success ....")
	basic.ClientConnHandler(conn)
}
