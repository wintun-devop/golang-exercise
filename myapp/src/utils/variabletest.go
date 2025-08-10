package utils

import (
	"fmt"
)

var MyPackageName string = "utils"

func MyVariableTest() {
	type Product struct {
		ProductName string `json:"product_name"`
	}
	fmt.Println("This is my variable test function")
	fmt.Println("package name:", MyPackageName)

	var x int16 = 170
	var y int16 = 83
	//Addition
	fmt.Printf(" addition :  %d + %d = %d\n ", x, y, x+y)
	//Subtraction
	fmt.Printf("subtraction : %d - %d = %d\n", x, y, x-y)
	//Multiplication
	fmt.Printf(" multiplication : %d * %d = %d\n", x, y, x*y)
	//Division
	fmt.Printf(" division : %d / %d = %d\n", x, y, x/y)
	//Modulus
	fmt.Printf(" remainder : %d %% %d = %d\n", x, y, x%y)
	var userName string
	var userEmaail string
	var bookings [50]string

	// print out
	fmt.Printf("%v\n", userName)
	fmt.Printf("%v\n", userEmaail)
	fmt.Printf("%v\n", bookings)
	fmt.Printf("%T\n", bookings)
	// constatn
	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
