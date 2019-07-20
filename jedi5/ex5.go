package jedi5

import (
	"fmt"
)

// Ex5 tbd
func Ex5() {
	fmt.Println("\nEx5()")

	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c

	fmt.Println(v, ok)

	close(c)

	v, ok = <-c

	fmt.Println(v, ok)
}
