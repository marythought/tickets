package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func CheckTicket(winning_ticket string) (bool, error) {
	tickets, err := getTickets()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	for _, ticket := range tickets {
		if ticket == winning_ticket {
			return true, nil
		}
	}
	return false, nil
}

func getTickets() ([]string, error) {
	content, err := ioutil.ReadFile("tickets.txt")
	if err != nil {
		return []string{}, errors.New("no tickets.txt file found")
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
