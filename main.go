package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferencTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We are excited to have you here today.")
	fmt.Println("Get your tickets here to attend. ")
	fmt.Println("We have a total of", conferencTickets, "tickets and", remainingTickets, "are still available.")

}
