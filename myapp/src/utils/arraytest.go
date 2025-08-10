package utils

import (
	"fmt"
)

var numberArray [5]int = [5]int{10, 20, 30, 40, 50}

func MyArrayTest() {
	fmt.Println("integer array test")
	fmt.Println("thef first room", numberArray[0])
	// Modify elements
	numberArray[2] = 99
	fmt.Println(numberArray) // Output: [10 20 99 40 50]
	// Length of array
	fmt.Println("The length of array", len(numberArray)) // Output: 5
	stringArray := [5]string{"a", "e", "i", "o", "u"}
	fmt.Println("string array", stringArray)
	//dynamic array using slice
	names := []string{"Win", "Tun", "Hlaing"}
	fmt.Println("dynamic array", names)
	names = append(names, "Dev")
	fmt.Println("dynamic after append", names)
	// dynamic arrary looping
	for i, name := range names {
		fmt.Println(i, name)
	}
	// number of elements
	fmt.Println(len(names))
	// underlying array capacity
	fmt.Println(cap(names))
	/* mixed array */
	var mixedArray []interface{}

	// Insert different types
	mixedArray = append(mixedArray, 42)      // int
	mixedArray = append(mixedArray, "hello") // string
	mixedArray = append(mixedArray, 'A')     // rune (character)
	mixedArray = append(mixedArray, 3.14)    // float64

	// Print all elements
	for i, v := range mixedArray {
		fmt.Printf("Index %d: %v (type: %T)\n", i, v, v)
	}
}
