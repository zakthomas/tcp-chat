package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Start")
	s := newServer("Test1")
	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("Server error : ", err.Error())
	}
	log.Println("Server is up")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error : ", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}
