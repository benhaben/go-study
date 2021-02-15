package main

import (
	"fmt"
	"time"
)

//主goroutine，退出，子goroutine也退出
func main() {
	//create go routine
	go func() {
		defer fmt.Println("A.Defer")
		func() {
			defer fmt.Println("B.Defer")
			//runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	go func(a int, b int) bool{
		fmt.Println("a=", a, "b=", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
