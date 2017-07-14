package main

import (
	"fmt"
	"os"
)

// usage: tickets 0220173
func main() {
	winning_ticket := os.Args[1]
	if CheckTicket(winning_ticket) {
		fmt.Println("winner!!", winning_ticket)
	} else {
		fmt.Println("nope")
	}
}
