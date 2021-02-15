package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutime game over")
		fmt.Println("going...")
		c <- 666
	}()

	// channel 保证了main不会先退出
	num := <-c
	fmt.Println("num=", num)
	fmt.Println("main game over")
}
