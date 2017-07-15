package main

import (
	"testing"
	"github.com/slowtrailrunner/tickets/check"
)

func TestWinningTicket(t *testing.T) {
	actual, _ := check.Ticket("0020172")
	expected := true
	if expected != actual {
		t.Error(expected, actual, "expected not equal")
	}
}

func TestNotWinningTicket(t *testing.T) {
	actual, _ := check.Ticket("0020173")
	expected := false
	if expected != actual {
		t.Error(expected, actual, "expected not equal")
	}
}
