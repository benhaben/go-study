package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	fmt.Println("len(c)=", len(c), "cap(c)", cap(c))

	go func() {
		defer fmt.Println("sub goroutime game over")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("sub goroutime going, len(c)=", len(c), "cap(c)", cap(c), "i=", i)
		}
	}()

	// channel 保证了main不会先退出
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num=", num)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main game over")

}
