package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type client struct {
	ID           int
	toClientChan chan string
}

type msg struct {
	userID int
	body   string
}

var users map[int]client
var usersMutex *sync.Mutex

// SYSTEM_ID - ID for System messages
const SYSTEM_ID int = 0

func main() {
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal("Cannot listen:", err)
	}
	defer l.Close()

	users = make(map[int]client)
	usersMutex = &sync.Mutex{}
	toLinkerChan := make(chan msg, 10)
	go linker(toLinkerChan)
	nextID := 1

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
		}

		newuser := client{nextID, make(chan string, 5)}

		usersMutex.Lock()
		users[nextID] = newuser
		usersMutex.Unlock()

		go userConnected(conn, newuser, toLinkerChan)
		toLinkerChan <- msg{SYSTEM_ID, "Guest" + strconv.Itoa(newuser.ID) + " has joined the chat\n"}
		nextID++
	}
}

func linker(toLinkerChan chan msg) {
	var newmessage string

	for message := range toLinkerChan {
		if message.userID == 0 {
			newmessage = "System message: " + message.body
		} else {
			newmessage = "Guest" + strconv.Itoa(message.userID) + ": " + message.body
		}

		usersMutex.Lock()
		for _, u := range users {
			select {
			case u.toClientChan <- newmessage:
			case <-time.After(time.Millisecond * 10):
			}
		}
		usersMutex.Unlock()
	}
}

func userConnected(conn net.Conn, user client, toLinkerChan chan msg) {
	defer func() {
		usersMutex.Lock()
		delete(users, user.ID)
		usersMutex.Unlock()
		conn.Close()
		close(user.toClientChan)
	}()

	conn.Write([]byte("Hello Guest" + strconv.Itoa(user.ID) + "!\n"))

	go func() {
		for message := range user.toClientChan {
			conn.Write([]byte(message))
		}
	}()

	rd := bufio.NewReader(conn)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			log.Printf("Guest%v disconnected: %v", strconv.Itoa(user.ID), err)
			toLinkerChan <- msg{SYSTEM_ID, "Guest" + strconv.Itoa(user.ID) + " has been disconnected\n"}
			return
		}
		if str == "exit\n" || str == "exit\r\n" {
			log.Printf(`Guest%v typed "exit"`, strconv.Itoa(user.ID))
			toLinkerChan <- msg{SYSTEM_ID, "Guest" + strconv.Itoa(user.ID) + " has been disconnected\n"}
			return
		}
		toLinkerChan <- msg{user.ID, str}
	}
}
