package basic

import (
	"fmt"
	"log"
	"net"
)

// ConnMap ...
var ConnMap = make(map[string]net.Conn)

func ServerConnHandler(c net.Conn) {
	ConnMap[c.RemoteAddr().String()] = c

	fmt.Println(len(ConnMap))
	defer c.Close()

	log.Println("Connect :", c.RemoteAddr())
	buf := make([]byte, 1024)
	var data string

ILOOP:
	for {
		// access requesr from clinet
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			break ILOOP
		}

		addr := c.RemoteAddr()
		if n == 0 {
			log.Printf("%s has disconnect\n", addr)
		}

		data = string(buf[:n])
		log.Printf("Receive '%s' from '%s'\n", data, addr)

		// time.Sleep(10 * time.Second)
		// c.Write([]byte(buf[:n]))
		// c.Write([]byte("2018-12-28 14:49:00"))
		for _, conn := range ConnMap {
			fmt.Println(conn)
			// conn.Write(b)
			_, err := conn.Write(buf[:n])

			fmt.Println(err)
		}
		// inStr := strings.TrimSpace(data)
		// inputs := strings.Split(inStr, " ")
		// if !action(c, inputs[0], inputs[1:]) {
		// 	break ILOOP
		// }
	}

	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
}

// func boradcastMessage(message string) {

// 	b := []byte(message)
// 	for _, conn := range ConnMap {
// 		fmt.Println(conn)
// 		// conn.Write(b)
// 		conn.Write(b)
// 	}
// }

func action(c net.Conn, a string, m []string) bool {
	// time.Sleep(20 * time.Second)
	// c.Write([]byte("ping"))

	// switch a {
	// case "ping":
	// 	time.Sleep(20 * time.Second)
	// 	log.Printf("sned data to %s", c.RemoteAddr())
	// 	c.Write([]byte("ping"))
	// case "echo":
	// 	c.Write([]byte(strings.Join(m, " ")))
	// case "quit":
	// 	fmt.Println("quit")
	// 	c.Close()
	// 	return false
	// }

	return true
}
