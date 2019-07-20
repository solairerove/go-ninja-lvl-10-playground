package jedi2

import (
	"fmt"
)

// Ex2 tbd
func Ex2() {
	fmt.Println("\nEx2()")

	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)
}
