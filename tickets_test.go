package main

import "testing"

func TestWinningTicket(t *testing.T) {
	actual := CheckTicket("0020172")
	expected := true
	if expected != actual {
		t.Error(expected, actual, "expected not equal")
	}
}

func TestNotWinningTicket(t *testing.T) {
	actual := CheckTicket("0020173")
	expected := false
	if expected != actual {
		t.Error(expected, actual, "expected not equal")
	}
}
