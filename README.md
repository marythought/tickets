# Raffle Ticket Checker!

A program to check if you are holding a winning raffle ticket at #gophercon.

This program does not tell you the winning raffle ticket numbers in advance, sorry. But it can help you get organized to immediately identify if one of your tickets is a winner, after they announce the winners.

## Install

`git clone github.com/marythought/tickets`

## Customize

Replace tickets.txt file with your personal ticket numbers.

## Run

 `$ go build`
  
 AT RAFFLE TICKET ANNOUNCEMENT TIME type in "tickets" followed by the winning ticket number they announce:
 
 `$ tickets 0220173`
 
 If you are holding a winning ticket number, you will see:
 
`winner!!, 0220173`

or

`nope`

If you add additional ticket numbers, they will be read at runtime. No need to rebuild. :)

## License
MIT
