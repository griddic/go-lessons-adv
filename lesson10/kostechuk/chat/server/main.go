package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
)

// type channels struct {
// 	toClient chan string
// 	toLinker chan string
// }

type client struct {
	ID           int
	cmd          string
	toClientChan chan string
}

type msg struct {
	userID int
	body   string
}

func main() {
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal("Cannot listen:", err)
	}
	defer l.Close()

	cmdChan := make(chan client)
	toLinkerChan := make(chan msg, 10)
	go linker(cmdChan, toLinkerChan)
	nextID := 1

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
		}

		newuser := client{nextID, "add", make(chan string)}
		go userConnected(conn, newuser, cmdChan, toLinkerChan)
		cmdChan <- newuser
		nextID++
	}
}

func linker(cmdChan chan client, toLinkerChan chan msg) {
	users := map[int]client{}
	usereventmsg := ""

	for {
		select {
		case userevent := <-cmdChan:
			if userevent.cmd == "add" {
				users[userevent.ID] = userevent
				usereventmsg = "System message: Guest" + strconv.Itoa(userevent.ID) + " has joined the chat\n"
			} else if userevent.cmd == "del" {
				delete(users, userevent.ID)
				usereventmsg = "System message: Guest" + strconv.Itoa(userevent.ID) + " has been disconnected\n"
			}

			if usereventmsg != "" {
				for _, u := range users {
					u.toClientChan <- usereventmsg
				}
				usereventmsg = ""
			}
		case message := <-toLinkerChan:
			newmessage := "Guest" + strconv.Itoa(message.userID) + ": " + message.body
			for _, u := range users {
				u.toClientChan <- newmessage
			}
		}

		// if usereventmsg != "" {
		// 	for _, u := range users {
		// 		u.chans.toClient <- usereventmsg
		// 	}
		// 	usereventmsg = ""
		// }
		//
		// for _, u := range users {
		// 	select {
		// 	case newmessage = <-u.chans.toLinker:
		// 		newmessage = "Guest" + strconv.Itoa(u.ID) + ": " + newmessage
		// 		break
		// 	default:
		// 	}
		// }
		//
		// if newmessage != "" {
		// 	for _, u := range users {
		// 		u.chans.toClient <- newmessage
		// 	}
		// 	newmessage = ""
		// }
	}
}

func userConnected(conn net.Conn, user client, cmdChan chan client, toLinkerChan chan msg) {
	conn.Write([]byte("Hello Guest" + strconv.Itoa(user.ID) + "!\n"))
	defer conn.Close()
	defer close(user.toClientChan)

	go func() {
		for message := range user.toClientChan {
			conn.Write([]byte(message))
		}
	}()

	rd := bufio.NewReader(conn)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			cmdChan <- client{user.ID, "del", nil}
			log.Printf("Guest%v disconnected: %v", strconv.Itoa(user.ID), err)
			return
		}
		// log.Println([]byte(str))
		if str == "exit\n" || str == "exit\r\n" {
			cmdChan <- client{user.ID, "del", nil}
			log.Printf(`Guest%v typed "exit"`, strconv.Itoa(user.ID))
			return
		}
		toLinkerChan <- msg{user.ID, str}
	}
}
