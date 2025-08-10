package main

import (
	"fmt"
	"os"

	"myapp/helpers"
	"myapp/utils"

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

	fmt.Printf("API Key form Env%v\n", test)

	// calling outside function
	result := utils.Add(3, 5)
	fmt.Println("Sum: test inside moodule", result)
	id := helpers.GenerateCUID()
	fmt.Println("Generated CUID:", id)

	// Testing  user input
	// fmt.Printf("userName:\n")
	// fmt.Scan(&userName)
	// fmt.Printf("userEmail:\n")
	// fmt.Scan(&userEmaail)
	utils.MyVariableTest()
	executed()
	utils.MyLoopTest()
	utils.MyConditionalTest()
	utils.MyArrayTest()
}

func executed() {
	fmt.Println("I just got executed!")
}
