package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/netutil"
)

type Candidate struct {
	Name     string
	No       int
	NumVotes int
}

func NewCandidate(Name string, No int) *Candidate {
	c := &Candidate{Name: Name, No: No, NumVotes: 0}

	return c
}

type Server struct {
	Address    string
	Candidates []*Candidate
}

// Creating a new TCP Listener, accepting maximum of 100 concurrent connections.
// For each connection it follows one of the two defined protocol:
// Protocol 1:
// Client asks for list of candidates by sending request with body: "candidates"
// Server sends the list
// Cliend sends the number that he wants to vote for
// Server stores the vote, responds with "OK" and closes the connection
// Protocol2 :
// Client asks for the list of votes per candidate by sending request with body: "votes"
// Server sends the list of votes per candidate, and closes the connection
func (s *Server) Serve() error {

	// Opening the
	l, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}

	l = netutil.LimitListener(l, 100)

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println("error reading message", err)
			}
			fmt.Printf("Recieved from client: %v", msg)
			fmt.Fprint(c, "Hello from server\n")
			c.Close()
		}(conn)

	}
}
