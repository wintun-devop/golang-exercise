package utils

import (
	"fmt"
)

// Add takes two integers and returns their sum.
func Add(a int, b int) int {
	fmt.Println("print in direct function a", a, "b", b)
	return a + b
}

func VariableTest() string {
	const successMessage = "success"
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	// func (p Person) Name() string {
	// return p.FirstName + " " + p.LastName
	// }
	// func (p Person) Describe() string {
	// return fmt.Sprintf("%s is %d years old", p.Name(), p.Age)
	// }
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
	return successMessage
}
