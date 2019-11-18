package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"socketlab/server/broadcast"
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
		startServerMode()
	} else if t == "client" {
		startClientMode()
	}
}

func startServerMode() {
	log.Println("Starting socker server ...")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Println(err)
	}
	cl := broadcast.ClientManager{
		Clients:    make(map[*broadcast.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *broadcast.Client),
		Unregister: make(chan *broadcast.Client),
	}
	go cl.Start()

	for {
		connection, _ := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		cl.StartServerMode(connection)
	}
}

func startClientMode() {
	fmt.Println("Starting client...")
	connection, error := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	if error != nil {
		fmt.Println(error)
	}

	broadcast.StartClientMode(connection)
}
