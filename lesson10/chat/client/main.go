package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2222")
	if err != nil {
		log.Fatal("Cannot connect", err)
	}
	defer conn.Close()
	conn.Write([]byte("Hello!"))
	time.Sleep(time.Minute * 60)
}
