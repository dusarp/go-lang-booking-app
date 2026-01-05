package main

import "fmt"

func main() {
	// Define constants and variables
	const totalTickets int = 50
	var remainingTickets int = 50
	var bookings []string //this is a dynamic array (slice)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

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

		// slices: abstraction of an array, ha nem tudjuk elore hogy hany elemu lesz

		//igy jobb, mivel van first name es last name is

		bookings = append(bookings, firstName+" "+lastName)

		// print some array elements
		fmt.Printf("The bookings are: %v\n", bookings)
		fmt.Printf("The first booking is: %v\n", bookings[0])
		fmt.Printf("The slice length is: %v\n", len(bookings))

		// 3. Logic to calculate and display results
		if userTickets <= remainingTickets && userTickets > 0 {
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("\nSuccess! %v %v, you have booked %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("A confirmation email has been sent to %v.\n", email)
			fmt.Printf("There are %v tickets still available.\n", remainingTickets)

			// Check if all tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("\nAll tickets are sold out! Thank you for participating.")
				break
			}
		} else {
			fmt.Printf("\nSorry, we cannot fulfill that request. We only have %v tickets left.\n", remainingTickets)
		}
	}

}
