# Raffle Ticket Checker!

A program to check if you are holding a winning raffle ticket at #gophercon.

## Install

`go get github.com/marythought/tickets`

## Customize

Replace tickets.go line 4 with your personal ticket numbers.
`tickets := []string{"0020172", "0225210", "0924192", ... }`

## Run

 `$ go build`
  
 Type in "tickets" followed by the winning ticket number:
 
 `$ tickets 0220173`
 
 If you are holding a winning ticket number, you will see:
 
`winner!!, 0220173`

or

`nope`

## License
MIT
