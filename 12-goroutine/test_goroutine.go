package main

import (
	"fmt"
	"time"
)

//子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new routine: i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

//主goroutine，退出，子goroutine也退出
func main() {
	//create go routine
	go newTask()

	i := 0

	for {
		i++
		fmt.Printf("main routine: i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}
