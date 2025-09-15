package utils

// If you omit the size when initializing an array, Go will infer it from the number of elements you provide:
import (
	"fmt"
)

func MyArrayTest2() {
	// number array
	nums_array := [...]int{1, 2, 3}
	fmt.Println(nums_array)
	// straing array and  Range loop
	wordsArray := []string{"go", "is", "cool"}

	for index, w := range wordsArray {
		fmt.Println(index, w)
	}

}
