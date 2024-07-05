package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const (
	HOST = "0.0.0.0"
	PORT = "8080"
	TYPE = "tcp"
)

type Person struct {
	Name       string
	Connection net.Conn
}

func main() {
	fmt.Println("Starting TCP Chat Server...")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
	defer listen.Close()
	fmt.Println("Server listening on", HOST+":"+PORT)

	people := []Person{}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go handleRequest(conn, &people)
	}
}

func handleRequest(conn net.Conn, people *[]Person) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	name := ""
	gaveName := false
	newUser := true

	for {
		if !gaveName {
			conn.Write([]byte("Welcome to budgetchat! What shall I call you?\n"))
		}
		message, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %s", err)
			}
			if gaveName && !newUser {
				idx := 0
				for i, v := range *people {
					if v.Name == name {
						idx = i
					} else {
						v.Connection.Write([]byte("* " + name + " has left the room\n"))
					}
				}
				// todo: this needs to be atomic
				if len(*people) > 1 {
					*people = append((*people)[:idx], (*people)[idx+1:]...)
				} else {
					*people = (*people)[:idx]
				}
			}
			return
		}
		message = strings.Replace(message, "\n", "", -1)
		message = strings.Replace(message, "\r", "", -1)
		// fmt.Printf("input: %v", message)
		if !gaveName && newUser {
			if 0 < len(message) && len(message) <= 16 {
				name = message
				gaveName = true
			} else {
				conn.Write([]byte("Names must be between 1 and 16 characters inclusive\n"))
				conn.Close()
				return
			}

			names := []string{}
			for _, v := range *people {
				// todo: check if connection is still open?
				v.Connection.Write([]byte("* " + name + " has entered the room\n"))
				names = append(names, v.Name)
			}
			if len(names) > 0 {
				snames := strings.Join(names, ", ")
				conn.Write([]byte("* The room contains: " + snames + "\n"))
			}

			*people = append(*people, Person{Name: name, Connection: conn})
			newUser = false
		} else {
			for _, v := range *people {
				if v.Name != name {
					v.Connection.Write([]byte("[" + name + "]" + " " + message + "\n\r"))
				}
			}
		}

	}
}
