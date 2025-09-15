package utils

import (
	"fmt"
)

func LoopTest1() {
	// Classic loop
	// fmt.Println("loop testing(classic loop):")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// // while styple
	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)
	// loop and concurrency
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

}
