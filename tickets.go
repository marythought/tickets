package main

func CheckTicket(winning_ticket string) bool {
	tickets := []string{"0020172", "0225210", "0924192", "0223067", "0125163", "0222110", "0911115", "0119122", "0119123", "0035133", "0027099"}
	for _, ticket := range tickets {
		if ticket == winning_ticket {
			return true
		}
	}
	return false
}
