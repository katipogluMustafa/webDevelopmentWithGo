package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	/*
	 * net.Listen() takes 2 arguments
	 * 1st one is the kind of server you want to listen
	 * 2dn what port do you wanna listen on
	 * net.Listen() will give us a Listener and a Error
	 * Listener has Accept(), Close() and Addr() methods.
	 */
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		// connection is both a Reader and Writer
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I Hope!")

		conn.Close()
	}

}
