package main

import (
	"fmt"
	"myapp/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	test := os.Getenv("APP_SECRET_KEY")
	fmt.Printf("API Key form Env%v\n", test)
	// utils.MyArrayTest2()
	utils.LoopTest1()
}

/*
// local function can not import

func executed() {
	fmt.Println("I just got executed!")
}

*/
