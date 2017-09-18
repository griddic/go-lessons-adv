package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2222")
	if err != nil {
		log.Fatal("Cannot connect ", err)
	}
	defer conn.Close()

	go func() {
		rd := bufio.NewReader(conn)
		for {
			str, err := rd.ReadString('\n')
			if err != nil {
				log.Fatal("Connection lost ", err)
			}
			fmt.Println(str)
		}
	}()

	for {
		input := ""
		_, err = fmt.Scanln(&input)
		input = input + "\n"
		conn.Write([]byte(input))
	}
	// time.Sleep(time.Second * 30)
}
