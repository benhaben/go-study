package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("a=", a, "b=", b)
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)

}

func swap(a, b *int) {
	*a, *b = *b, *a
}
