package main

import (
	"bufio"
	"compress/gzip"
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

	log.Println("Client started & connected.")

	gzWriter := gzip.NewWriter(conn)
	log.Println("gzWriter created.")

	gzReader, err := gzip.NewReader(conn)
	log.Println("gzReader created.")

	if err != nil {
		log.Printf("Something wrong with gzReader. %v", err)
		return
	}

	go func() {
		rd := bufio.NewReader(gzReader)
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
			gzWriter.Write([]byte(input))
			err := gzWriter.Flush()

			if err != nil {
				log.Printf("Can't flush. %v", err)
				return
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal("Scanner broke. ", err)
		}
	}
}
