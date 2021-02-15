package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	fmt.Println("len(c)=", len(c), "cap(c)", cap(c))

	go func() {
		defer fmt.Println("sub goroutime game over")
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("sub goroutime going, len(c)=", len(c), "cap(c)", cap(c), "i=", i)
		}
		close(c)
	}()

	//for  {
	//	if data,ok := <-c; ok{
	//		fmt.Println("data=",data)
	//	}else{
	//		break
	//	}
	//}

	for data := range c {
		fmt.Println("data=", data)
	}
	fmt.Println("main game over")

}
