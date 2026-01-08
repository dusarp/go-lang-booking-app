package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define constants and variables
	const totalTickets int = 50
	var remainingTickets int = 50
	var bookings []string //this is a dynamic array (slice)

	for remainingTickets > 0 && len(bookings) < totalTickets {
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

		//validation
		var isValidName:= len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail := strings.Contains(email, "@")
		var isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		//var isValidCity := city == "Singapore" || city == "London" || city == "Berlin"
		//var isInValidCity := city != "Singapore" && city != "London" && city != "Berlin"

		// slices: abstraction of an array, ha nem tudjuk elore hogy hany elemu lesz

		//igy jobb, mivel van first name es last name is

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The slice length is: %v\n", len(bookings))

		// 3. Logic to calculate and display results
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("\nSuccess! %v %v, you have booked %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("A confirmation email has been sent to %v.\n", email)
			fmt.Printf("There are %v tickets still available.\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings { // _ means we dont want to use that variable
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings so far: %v\n", firstNames)

			// Check if all tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("\nAll tickets are sold out! Thank you for participating.")
				break
			}
		} else {
			fmt.Println("Your input is invalid. Please try again.")
		}
	}

}
