package main

import "fmt"

// main is the entry point for the application (ez egy komment)
func main() {

	var conferenceName string = "Go Conference"
	const conferencTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Println("We are excited to have you here today.")
	fmt.Println("Get your tickets here to attend. ")
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferencTickets, remainingTickets)

	var userName string
	// ask user for their name
	userName = "Alice"
	userTicket := 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)
}
