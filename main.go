// Package main implements a conference ticket booking application.
package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// Conference configuration
const conferenceName = "Go conference"
const totalTickets int = 50

// Global state
var remainingTickets int = 50
var bookings = make([]UserData, 0)

// UserData represents a user's booking information.
type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets int
}

func main() {
	greetUsers()

	for remainingTickets > 0 && len(bookings) < totalTickets {
		firstName, lastName, email, userTickets := getUserInput(remainingTickets)
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookings, remainingTickets = bookTicket(userTickets, firstName, lastName, email, bookings, remainingTickets)

			// Send ticket asynchronously using a goroutine
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

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

// greetUsers displays a welcome message with conference details.
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets available for this conference.\n\n", totalTickets)
}

// getFirstNames extracts and returns all first names from the bookings slice.
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// getUserInput prompts the user for their booking details and returns the collected information.
func getUserInput(remainingTickets int) (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

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

// bookTicket processes a booking by adding user data and updating ticket count.
// Returns the updated bookings slice and remaining tickets.
func bookTicket(userTickets int, firstName string, lastName string, email string, bookings []UserData, remainingTickets int) ([]UserData, int) {
	remainingTickets = remainingTickets - userTickets

	userData := UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("\nSuccess! %v %v, you have booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("A confirmation email has been sent to %v.\n", email)
	fmt.Printf("There are %v tickets still available.\n", remainingTickets)

	return bookings, remainingTickets
}

// sendTicket simulates sending a ticket confirmation email to the user.
// This function is designed to be run as a goroutine.
func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
}
