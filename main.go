package main

import (
	"fmt"
	"github.com/slowtrailrunner/tickets/check"
)

// usage: tickets 0220173
func main() {
	quitters := map[string]bool{"quit":true, "q":true}
	var ticket string
	for ticket != "quit" {
		fmt.Print("Please enter your ticket number ('quit or q' to quit): ")
		fmt.Scan(&ticket)
		if !quitters[ticket]  {
			fmt.Println(ticket, "is a ", checkTicket(ticket))
		} else {
			fmt.Println("Good bye.")
			break
		}
	}
}

func checkTicket(ticket string) string {
	winner, err := check.Ticket(ticket)
	if err != nil {
		fmt.Println(err)
	}
	if winner {
		return "Winner!!"
	} else {
		return "Loser!!"
	}
}
