package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// Listen for tcp connections
	listener, err := net.Listen("tcp", "0.0.0.0:9000")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	log.Println("Server is listening on port 9000")

	for {
		// Block until we recieve an incoming connection
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle the connection in a separate go routine
		go handleClient(conn)
	}

}

// handleClient handles a tcp connection
func handleClient(conn net.Conn) {
	// Close the connection after task has completed
	defer conn.Close()

	// read data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	log.Println("Recived data:", buf[:n])

	// write back the same data
	conn.Write(buf[:n])

}
