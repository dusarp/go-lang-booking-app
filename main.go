package main

import "fmt"

func main() {
	// Define constants and variables
	const totalTickets int = 50
	var remainingTickets int = 50
	var firstName string
	var lastName string
	var email string
	var userTickets int

	var bookings = [50]string{} //go ban meg kell hatarozni hogy mennyi elemu lehet benne, meg a type homogen kell hogy legyen
	//vagy booking [50]string
	bookings[0] = "Nana"
	bookings[1] = "Ana"
	bookings[2] = "Arpad"
	bookings[3] = "Doe"

	// 1. Initial greeting
	fmt.Println("Welcome to the Conference Booking System!")
	fmt.Printf("There are %v tickets available in total.\n\n", remainingTickets)

	// 2. Ask user for their information
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets would you like to buy? ")
	fmt.Scan(&userTickets)

	// 3. Logic to calculate and display results
	if userTickets <= remainingTickets && userTickets > 0 {
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("\nSuccess! %v %v, you have booked %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("A confirmation email has been sent to %v.\n", email)
		fmt.Printf("There are %v tickets still available.\n", remainingTickets)
	} else {
		fmt.Printf("\nSorry, we cannot fulfill that request. We only have %v tickets left.\n", remainingTickets)
	}

}
