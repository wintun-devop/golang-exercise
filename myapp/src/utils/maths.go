package utils

import (
	"fmt"
)

// Add takes two integers and returns their sum.
func Add(a int, b int) int {
	fmt.Println("print in direct function a", a, "b", b)
	return a + b
}

// func VariableTest() string {
// 	const successMessage = "success"
// 	type Person struct {
// 		FirstName string
// 		LastName  string
// 		Age       int
// 	}
// 	// func (p Person) Name() string {
// 	// return p.FirstName + " " + p.LastName
// 	// }
// 	// func (p Person) Describe() string {
// 	// return fmt.Sprintf("%s is %d years old", p.Name(), p.Age)
// 	// }
// 	return successMessage
// }
