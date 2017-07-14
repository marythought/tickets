package main

import (
	"fmt"
	"os"
)

// usage: tickets 0220173
func main() {
	winning_ticket := os.Args[1]
	winner, err := CheckTicket(winning_ticket)
	if err != nil {
		fmt.Println(err)
	}
	if winner {
		fmt.Println("winner!!", winning_ticket)
	} else {
		fmt.Println("nope")
	}
}
