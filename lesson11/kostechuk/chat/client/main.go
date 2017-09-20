package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
				log.Fatal("Connection lost. ", err)
			}
			fmt.Print(str)
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text() + "\n"
			conn.Write([]byte(input))
		}

		if err := scanner.Err(); err != nil {
			log.Fatal("Scanner broke. ", err)
		}
	}
}
