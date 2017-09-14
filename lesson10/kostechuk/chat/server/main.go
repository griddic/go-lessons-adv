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

// Users - global access map listing all users connected
var Users map[int]client

// SYSTEM_ID - ID for System messages
const SYSTEM_ID int = 0

func main() {
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal("Cannot listen:", err)
	}
	defer l.Close()

	usersmutex := &sync.Mutex{}
	Users = make(map[int]client)
	toLinkerChan := make(chan msg, 10)
	go linker(toLinkerChan, usersmutex)
	nextID := 1

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
		}

		newuser := client{nextID, make(chan string, 5)}

		usersmutex.Lock()
		Users[nextID] = newuser
		usersmutex.Unlock()

		go userConnected(conn, newuser, toLinkerChan, usersmutex)
		toLinkerChan <- msg{SYSTEM_ID, "Guest" + strconv.Itoa(newuser.ID) + " has joined the chat\n"}
		nextID++
	}
}

func linker(toLinkerChan chan msg, usersmutex *sync.Mutex) {
	var newmessage string

	for {
		select {
		case message := <-toLinkerChan:
			if message.userID == 0 {
				newmessage = "System message: " + message.body
			} else {
				newmessage = "Guest" + strconv.Itoa(message.userID) + ": " + message.body
			}

			usersmutex.Lock()
			for _, u := range Users {
				select {
				case u.toClientChan <- newmessage:
				case <-time.After(time.Millisecond * 10):
				}
			}
			usersmutex.Unlock()
		}
	}
}

func userConnected(conn net.Conn, user client, toLinkerChan chan msg, usersmutex *sync.Mutex) {
	defer func() {
		usersmutex.Lock()
		delete(Users, user.ID)
		usersmutex.Unlock()
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
