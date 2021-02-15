package main

import "fmt"

func main() {
	var a string
	//pair<static type: string, value:"acb">
	a = "abc"

	var allType interface{}
	//pair<static type: string, value:"acb">
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
