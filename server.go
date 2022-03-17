package main

import (
	"log"
	"net"
	"strings"
)

type server struct {
	name    string
	members map[net.Addr]*client
	chat    chan message
}

func newServer(name string) *server {
	log.Println("New server setup")
	return &server{
		name:    name,
		members: make(map[net.Addr]*client),
		chat:    make(chan message),
	}
}

func (s *server) newClient(connection net.Conn) {
	log.Println("New client connected")
	c := &client{id: connection.RemoteAddr(), name: "newbie(" + (connection.RemoteAddr().String()) + ")", connection: connection, chat: s.chat}
	s.members[connection.RemoteAddr()] = c
	s.broadcast(message{from: c, msg: "A new player has joined the chat"})
	c.readInput()
}

func (s *server) run() {
	for ping := range s.chat {
		command := strings.Trim(ping.msg, "\r\n")
		switch {
		case strings.HasPrefix(command, "my name is"):
			newName := strings.TrimPrefix(command, "my name is")
			s.broadcast(ping)
			ping.from.name = newName
		case command == "quit":
			ping.msg = "Goodbye everyone..."
			ping.from.connection.Close()
			delete(s.members, ping.from.id)
			s.broadcast(ping)
		default:
			s.broadcast(ping)
		}
	}
}

func (s *server) broadcast(msg message) {
	for _, client := range s.members {
		client.hear(msg)
	}
}
