package utils

// If you omit the size when initializing an array, Go will infer it from the number of elements you provide:
import (
	"fmt"
)

func VariableTest() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainTickests = 50

	fmt.Printf("Welcome to Our %v App.\n", conferenceName)
	fmt.Printf("We have total of %v and %v remaining.\n", conferenceTickets, remainTickests)
	fmt.Println("Get your ticket and go confrence.")
	fmt.Printf("conferenceName is %T and conferenceTickets is %T.\n", conferenceName, conferenceTickets)
}
