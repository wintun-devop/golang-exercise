package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainTickests = 50

	fmt.Printf("Welcome to Our %v App.\n", conferenceName)
	fmt.Printf("We have total of %v and %v remaining.\n", conferenceTickets, remainTickests)
	fmt.Println("Get your ticket and go confrence.")
}

func myMessage() {
	fmt.Println("I just got executed!")
}
