package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal("Cannot listen:", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
		}
		go clientConnected(conn)
	}
}

func clientConnected(conn net.Conn) {
	conn.Write([]byte("Hello, stranger!\n"))
	defer conn.Close()
	rd := bufio.NewReader(conn)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			log.Printf("Client disconnected: %v", err)
			return
		}
		// log.Println([]byte(str))
		if str == "exit\n" || str == "exit\r\n" {
			log.Printf(`Client typed "exit"`)
			return
		}
		time.Sleep(time.Second)
		conn.Write([]byte(str))
	}
}
