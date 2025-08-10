package utils

import (
	"fmt"
)

func MyLoopTest() {
	fmt.Println("Testing Loop")
	for x := 1; x <= 5; x++ {
		fmt.Println("Number:", x)
		// For loop breaks when the value of x = 4
		if x == 4 {
			break
		}
	}
	fmt.Println("Testing while loop")
}
