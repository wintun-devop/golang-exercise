package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	test := os.Getenv("APP_SECRET_KEY")
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainTickests = 50

	fmt.Printf("Welcome to Our %v App.\n", conferenceName)
	fmt.Printf("We have total of %v and %v remaining.\n", conferenceTickets, remainTickests)
	fmt.Println("Get your ticket and go confrence.")
	fmt.Printf("conferenceName is %T and conferenceTickets is %T.\n", conferenceName, conferenceTickets)

	var userName string
	var userEmaail string
	var bookings [50]string
	// user input
	fmt.Printf("userName:\n")
	fmt.Scan(&userName)
	fmt.Printf("userEmail:\n")
	fmt.Scan(&userEmaail)

	// print out
	fmt.Printf("%v\n", userName)
	fmt.Printf("%v\n", userEmaail)
	fmt.Printf("%v\n", bookings)
	fmt.Printf("%T\n", bookings)
	fmt.Printf("%v\n", test)

	// calling outside function
	executed()
}

func executed() {
	fmt.Println("I just got executed!")
}
