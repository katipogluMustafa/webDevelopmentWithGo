package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
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

func handle(conn net.Conn){
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	defer conn.Close()

	if err != nil {
		log.Fatalln("CONN Timeout")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn,"I heard you say: %s \n",ln)
	}
	/* Now Code gets here since connections time out*/
	fmt.Printf("CODE GET HERE")
}