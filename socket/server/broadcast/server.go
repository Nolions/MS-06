package broadcast

import (
	"fmt"
	"net"
)

type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func (m ClientManager) StartServerMode(c net.Conn) {
	client := &Client{Socket: c, Data: make(chan []byte)}
	m.Register <- client
	go m.receive(client)
	go send(client)
}

func (m *ClientManager) receive(client *Client) {
	for {
		var message = make([]byte, 4096)
		length, err := client.Socket.Read(message)
		addr := client.Socket.RemoteAddr()
		if err != nil {
			m.Unregister <- client
			client.Socket.Close()
			break
		}
		if length > 0 {
			// data := fmt.Sprintf("%s, form %s", string(message), addr.String())
			data := fmt.Sprint(string(message) + addr.String())
			fmt.Print("RECEIVED: " + data)
			// log.Printf("echo msg: %s, form to %s", string(buf), addr)
			m.Broadcast <- []byte(data)
		}
	}
}

func send(client *Client) {
	defer client.Socket.Close()
	for {
		select {
		case message, ok := <-client.Data:
			if !ok {
				return
			}
			client.Socket.Write(message)
		}
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		case connection := <-manager.Register:
			manager.Clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.Unregister:
			if _, ok := manager.Clients[connection]; ok {
				close(connection.Data)
				delete(manager.Clients, connection)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.Broadcast:
			for connection := range manager.Clients {
				select {
				case connection.Data <- message:
				default:
					close(connection.Data)
					delete(manager.Clients, connection)
				}
			}
		}
	}
}
