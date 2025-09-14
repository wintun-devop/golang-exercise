package utils

// If you omit the size when initializing an array, Go will infer it from the number of elements you provide:
import (
	"fmt"
)

func MyArrayTest2() {
	nums_array := [...]int{1, 2, 3}
	fmt.Println(nums_array)
}
