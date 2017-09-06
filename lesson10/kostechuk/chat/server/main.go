package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
)

type channels struct {
	toClient chan string
	toLinker chan string
}

type client struct {
	ID    int
	cmd   string
	chans channels
}

func main() {
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal("Cannot listen:", err)
	}
	defer l.Close()

	cmdChan := make(chan client)
	go linker(cmdChan)
	nextID := 1

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
		}

		newuser := client{nextID, "add", channels{make(chan string), make(chan string)}}
		go userConnected(conn, newuser, cmdChan)
		cmdChan <- newuser
		nextID++
	}
}

func linker(cmdChan chan client) {
	users := map[int]client{}
	usereventmsg := ""
	newmessage := ""

	for {
		select {
		case userevent := <-cmdChan:
			if userevent.cmd == "add" {
				users[userevent.ID] = userevent
				usereventmsg = "System message: Guest" + strconv.Itoa(userevent.ID) + " has joined the chat\n"
			} else if userevent.cmd == "del" {
				delete(users, userevent.ID)
				usereventmsg = "System message: Guest" + strconv.Itoa(userevent.ID) + " has disconnected\n"
			}
		default:
		}

		if usereventmsg != "" {
			for _, u := range users {
				u.chans.toClient <- usereventmsg
			}
			usereventmsg = ""
		}

		for _, u := range users {
			select {
			case newmessage = <-u.chans.toLinker:
				newmessage = "Guest" + strconv.Itoa(u.ID) + ": " + newmessage
				break
			default:
			}
		}

		if newmessage != "" {
			for _, u := range users {
				u.chans.toClient <- newmessage
			}
			newmessage = ""
		}
	}
}

func userConnected(conn net.Conn, user client, cmdChan chan client) {
	conn.Write([]byte("Hello Guest" + strconv.Itoa(user.ID) + "!\n"))
	defer conn.Close()
	defer close(user.chans.toClient)
	defer close(user.chans.toLinker)

	go func() {
		for message := range user.chans.toClient {
			conn.Write([]byte(message))
		}
	}()

	rd := bufio.NewReader(conn)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			cmdChan <- client{user.ID, "del", channels{nil, nil}}
			log.Printf("Guest%v disconnected: %v", strconv.Itoa(user.ID), err)
			return
		}
		log.Println([]byte(str))
		if str == "exit\n" || str == "exit\r\n" {
			cmdChan <- client{user.ID, "del", channels{nil, nil}}
			log.Printf(`Guest%v typed "exit"`, strconv.Itoa(user.ID))
			return
		}
		user.chans.toLinker <- str
	}
}
