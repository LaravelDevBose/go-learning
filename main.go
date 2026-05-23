package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// init slice
	var bookings []string
	// var bookings  = []string{}
	// bookings := []string{}

	fmt.Printf("conferenceTickets is %T, conferenceName is %T, remainingTickets is %T\n", conferenceTickets, conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var userName string
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name

		fmt.Print("Enter Your First Name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			userName = firstName + " " + lastName
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName)

			fmt.Printf("Thank You %v for booking %v tickets. You will receive a confermation email at %v \n", userName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstnames := []string{}
			for _, booking := range bookings {
				var name = strings.Fields(booking)
				firstnames = append(firstnames, name[0])
			}
			fmt.Printf("The first names of bookings are : %v \n", firstnames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you enterd is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered dons't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("your number of ticket is invalid")
			}
		}

	}

}
