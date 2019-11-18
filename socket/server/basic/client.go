package basic

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

var message string

func ClientConnHandler(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// writr, send request to server

		c.Write([]byte(input))

		addr := c.RemoteAddr()

		if input == "quit" {
			return
		}

		// read response form server
		// _, _ :=
		c.Read(buf)
		log.Printf("echo msg: %s, form to %s", string(buf), addr)
	}
}
