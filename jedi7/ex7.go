package jedi7

import (
	"fmt"
)

// Ex7 tbd
func Ex7() {
	fmt.Println("\nEx7()")

	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
}
