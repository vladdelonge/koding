package moh

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

// Publisher is the counterpart for Subscriber.
// It is a HTTP server accepting websocket connections.
type Publisher struct {
	// Registered filters, holds pointers to open connections.
	// All clients are registered to the "all" key by default for allowing broadcasting.
	// Modifier operations on this type is made by registrar() function.
	filters *Filters

	auth Authenticator

	websocket.Server // implements http.Handler interface
}

// This is the magic subscription key for broadcast events.
// Hoping that it is unique enough to not collide with another key.
const all = "4658f005d49885355f4e771ed9dace10cca9563e"

type Authenticator interface {
	Authenticate(userName, kodingKey string) bool
}

// NewPublisher creates a new Publisher and returns a pointer to it.  The
// publisher will listen on addr and accept websocket connections from
// Subscribers.
func NewPublisher(auth Authenticator) *Publisher {
	p := &Publisher{
		auth:    auth,
		filters: NewFilters(),
		Server:  websocket.Server{},
	}
	p.Server.Handler = p.handleWebsocketConn
	return p
}

// Publish sends a message to registered Subscribers with the key.
func (p *Publisher) Publish(key string, message []byte) {
	// log.Println("Sending message to send channel", string(message))
	for c := range p.filters.Get(key) {
		select {
		case c.send <- message:
			// log.Println("Message sent to send channel")
		default:
			// Buffer is full, writer() is not fast enough to send all published messages .
			// Drop the websocket client and let it synchronize by re-connecting.
			log.Println("Websocket buffer is full. Dropping socket")
			go c.ws.Close()
		}
	}
}

// Broadcast sends a message to all of the connected Subscribers.
func (p *Publisher) Broadcast(message []byte) {
	p.Publish(all, message)
}

func (p *Publisher) handleWebsocketConn(ws *websocket.Conn) {
	c := &connection{
		ws:        ws,
		publisher: p,
		send:      make(chan []byte, 256),
		keys:      make(map[string]bool),
	}

	p.filters.Add(c, all)
	defer func() {
		p.filters.RemoveAll(c)
		close(c.send)
	}()

	go c.writer()
	c.reader()
}

// connection represents a connected Subscriber in Publisher.
type connection struct {
	ws        *websocket.Conn
	publisher *Publisher

	// Buffered channel of outbount messages
	send chan []byte

	// Subscription keys
	keys map[string]bool
}

// reader reads the subscription requests from websocket and saves it in a map
// for accessing later.
func (c *connection) reader() {
	for {
		var cmd subscriberCommand
		err := websocket.JSON.Receive(c.ws, &cmd)
		if err != nil {
			log.Println("reader: Cannot receive message from websocket")
			break
		}

		if c.publisher.auth != nil && !c.publisher.auth.Authenticate(cmd.Auth.Username, cmd.Auth.Password) {
			break
		}

		// log.Printf("reader: Received a command from websocket: %+v\n", cmd)
		if cmd.Name == "subscribe" {
			key := cmd.Args["key"].(string)
			c.publisher.filters.Add(c, key)
		} else if cmd.Name == "unsubscribe" {
			key := cmd.Args["key"].(string)
			c.publisher.filters.Remove(c, key)
		} else {
			log.Println("Unknown command, dropping client")
			break
		}
	}
	c.ws.Close()
}

// writer writes the messages to the websocket from the send channel.
func (c *connection) writer() {
	for message := range c.send {
		err := websocket.Message.Send(c.ws, message)
		if err != nil {
			log.Println("writer: Cannot send message to websocket")
			break
		}
	}
	c.ws.Close()
}
