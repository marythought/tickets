# Raffle Ticket Checker!

A program to check if you are holding a winning raffle ticket at #gophercon.

## Install

`git clone github.com/marythought/tickets`

## Customize

Replace tickets.txt file with your personal ticket numbers.

## Run

 `$ go build`
  
 Type in "tickets" followed by the winning ticket number:
 
 `$ tickets 0220173`
 
 If you are holding a winning ticket number, you will see:
 
`winner!!, 0220173`

or

`nope`

If you add additional ticket numbers, they will be read at runtime. No need to rebuild. :)

## License
MIT
