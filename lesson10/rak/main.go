package main

import (
	"bufio"
	"log"
	"net"
	"sync"
	"time"
)

type User struct {
	Name string
	Conn net.Conn
}
type Message struct {
	From User
	Text string
}
type Users struct {
	array map[int]User
	mutex sync.Mutex
}

var users Users

func (m Message) Send(u User) {
	u.Conn.Write([]byte(m.From.Name + m.Text))
}
func sender(c chan Message) {
	for {
		message := <-c
		for _, v := range users.array {
			message.Send(v)
		}
		time.Sleep(time.Second)
	}
}
func (users Users) delete_user(user_id int) {
	users.mutex.Lock()
	defer users.mutex.Unlock()
	delete(users.array, user_id)
}
func (users Users) add_user(user_id int, user User) {
	users.mutex.Lock()
	defer users.mutex.Unlock()
	users.array[user_id] = user
}
func main() {
	user_id := 0
	c1 := make(chan Message, 10)
	users.array = make(map[int]User)
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal("Cannot listen", err)
	}
	defer l.Close()
	go sender(c1)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
		}
		user_id++
		go handleConnection(conn, c1, user_id)
	}
}
func handleConnection(conn net.Conn, c chan Message, user_id int) {
	defer conn.Close()
	conn.Write([]byte("Plese, input your name: "))
	rd := bufio.NewReader(conn)
	name, err := rd.ReadString('\n')
	if err != nil {
		log.Printf("Error while reading user name: %v", err)
	}
	conn.Write([]byte("Welcome to secret chat, " + name))
	new_user := User{name, conn}
	users.add_user(user_id, new_user)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			log.Printf("Error while reading: %v", err)
		}
		if str == "exit\n" || str == "exit\r\n" {
			log.Printf("Client typed 'exit'")
			users.delete_user(user_id)
			return
		}
		message := Message{new_user, str}
		c <- message
	}
}
