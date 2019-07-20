package jedi3

import (
	"fmt"
)

// Ex3 tbd
func Ex3() {
	fmt.Println("\nEx3()")

	c := gen()
	receive(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
