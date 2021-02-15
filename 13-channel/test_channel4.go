package main

import (
	"fmt"
)

func fib(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			sum := x + y
			x = y
			y = sum
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		defer fmt.Println("sub goroutime game over")
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fib(c, quit)
	fmt.Println("main game over")

}
