package main

import (
	"booking-app/helper"
	"fmt"
)

// Define constants and variables
var conferenceName = "Go conference"

const totalTickets int = 50

var remainingTickets int = 50
var bookings = make([]map[string]string, 0) //slice of maps to store user data

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets int
}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < totalTickets {
		// Get user input
		firstName, lastName, email, userTickets := getUserInput(remainingTickets)
		// Validate user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//city validation
		//var isValidCity := city == "Singapore" || city == "London" || city == "Berlin"
		//var isInValidCity := city != "Singapore" && city != "London" && city != "Berlin"

		// slices: abstraction of an array, ha nem tudjuk elore hogy hany elemu lesz

		//bookings with first and last names

		// Logic to calculate and display results
		if isValidName && isValidEmail && isValidTicketNumber {
			bookings, remainingTickets = bookTicket(userTickets, firstName, lastName, email, bookings, remainingTickets)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings re: %v\n", firstNames)

			// Check if all tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("\nAll tickets are sold out! Thank you for participating.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain '@' sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}
			fmt.Println("Please try again.")
		}
	}

}

// Function to greet users
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets available for this conference.\n\n", totalTickets)
}

// Function to get first names from bookings
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ means we dont want to use that variable
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput(remainingTickets int) (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// Ask user for their information
	fmt.Printf("We have %v tickets remaining.\n", remainingTickets)
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets would you like to buy? ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string, bookings []map[string]string, remainingTickets int) ([]map[string]string, int) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = fmt.Sprint(userTickets)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("\nSuccess! %v %v, you have booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("A confirmation email has been sent to %v.\n", email)
	fmt.Printf("There are %v tickets still available.\n", remainingTickets)

	return bookings, remainingTickets
}
