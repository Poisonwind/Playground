package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	//adress of client
	name := conn.RemoteAddr().String()

	log.Printf("%v connected\n", name)

	//write to client
	conn.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))

	scaner := bufio.NewScanner(conn)

	for scaner.Scan() {
		text := scaner.Text()
		if text == "quit" {
			conn.Write([]byte("Connection closing"))
			log.Printf("%s disconnected", name)
			break
		}

		log.Printf("%s enters %s", name, text)
		conn.Write([]byte(fmt.Sprintf("You enter %s\n", text)))
	}

}

func main() {

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}

	for {

		//connection to socket
		conn, err := listner.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleConnection(conn)
	}

}
