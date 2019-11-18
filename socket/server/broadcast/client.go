package broadcast

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	Socket net.Conn
	Data   chan []byte
}

func StartClientMode(c net.Conn) {
	client := &Client{Socket: c}
	go client.Receive()

	for {
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		c.Write([]byte(strings.TrimRight(data, "\n")))
	}
}

func (client *Client) Receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.Socket.Read(message)
		if err != nil {
			client.Socket.Close()
			break
		}
		// if length > 0 {
		fmt.Println("RECEIVED: " + string(message[:length]))
		// }
	}
}
