package jedi4

import (
	"fmt"
)

// Ex4 tbd
func Ex4() {
	fmt.Println("\nEx4()")

	q := make(chan int)
	c := gen(q)

	receive(c, q)
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
	}()

	return c
}
