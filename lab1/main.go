package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	c1 := NewCandidate("Candidate 1", 1)
	c2 := NewCandidate("Candidate 2", 2)
	c3 := NewCandidate("Candidate 3", 3)
	c4 := NewCandidate("Candidate 4", 4)

	candidates := []*Candidate{c1, c2, c3, c4}

	srv := Server{
		Address:    "localhost:8080",
		Candidates: candidates,
	}

	fmt.Printf("Server listening on Addr: %s...\n", srv.Address)

	go func() {
		err := srv.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	conn, err := net.Dial("tcp", srv.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Hello from client\n")

	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recieved from server: %v", resp)
}
