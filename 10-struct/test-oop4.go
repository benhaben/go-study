package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	v, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not  string...")
	} else {
		fmt.Println("arg is string... = ", v)
		fmt.Printf("value type is %T\n", v)

	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{auth: "yin"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
