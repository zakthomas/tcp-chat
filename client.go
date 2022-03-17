package main

import (
	"bufio"
	"net"
)

type message struct {
	from *client
	msg  string
}
type client struct {
	id         net.Addr
	name       string
	connection net.Conn
	chat       chan<- message
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.connection).ReadString('\n')
		if err != nil {
			return
		}

		ping := message{from: c, msg: msg}
		c.chat <- ping
	}
}

func (c *client) hear(msg message) {
	if c.id != msg.from.id {
		c.connection.Write([]byte(msg.from.name + " > " + msg.msg + "\n"))
	}
}
