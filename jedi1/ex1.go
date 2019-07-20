package jedi1

import (
	"fmt"
)

// Ex1 tbd
func Ex1() {
	fmt.Println("\nEx1()")

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
