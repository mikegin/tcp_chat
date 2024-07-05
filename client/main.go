package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Connecting Bob")
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	result := make([]byte, 1024)
	n, _ := conn.Read(result)
	fmt.Println(string(result[:n]))

	// time.Sleep(1 * time.Second)

	// max := 10
	// min := 1
	// v := rand.Intn(max-min) + min
	// // Send the data
	conn.Write([]byte("bob" + "\n"))

	fmt.Println("Connecting Alice")
	conn2, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn2.Close()

	conn2.Write([]byte("alice" + "\n"))

	n, _ = conn.Read(result)
	fmt.Println(string(result[:n]))

	n, _ = conn2.Read(result)
	fmt.Println(string(result[:n]))

	// time.Sleep(1 * time.Second)

	// time.Sleep(1 * time.Second)

	// conn.Write([]byte("Hey there, I'm Bob"))
	// conn.Write([]byte("Hey there, I'm Alice"))
}
