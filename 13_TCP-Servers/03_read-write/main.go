package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main(){
	li, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handle(conn)

	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s \n",ln)
	}

	/* Code never gets here since we always listen the connection*/
	fmt.Printf("CODE GET HERE")
}
/*
 * Run the program
 * in another terminal $ telnet localhost 8080
 * after opening telnet connection, send messages and you'll see the output.
 */

/*
 * Run this file then open localhost:8080 on browser
 */