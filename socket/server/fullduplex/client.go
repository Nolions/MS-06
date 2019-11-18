package fullduplex

import (
	"log"
	"net"
)

const (
	Message       = "Ping"
	StopCharacter = "\r\n\r\n"
)

func ClientConnHandler(c net.Conn) {
	defer c.Close()

	c.Write([]byte(Message))
	c.Write([]byte(StopCharacter))
	log.Printf("Send: %s", Message)

	buff := make([]byte, 1024)
	n, _ := c.Read(buff)
	log.Printf("Receive: %s", buff[:n])
}
