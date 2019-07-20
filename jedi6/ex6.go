package jedi6

import (
	"fmt"
)

// Ex6 tbd
func Ex6() {
	fmt.Println("\nEx6()")

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
