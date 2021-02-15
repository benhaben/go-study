package main

import "fmt"

//const enum

const (
	BEI_JING = iota
	SHANG_HAI
	SHENG_ZHENG
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func foo1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

func foo2(a string, b int) (r1 int, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	r1 = 100
	r2 = 200
	return r1, r2
}
func foo3(a string, b int) (r1, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	r1 = 1000
	r2 = 2000
	return r1, r2
}

func main() {
	fmt.Println("hello go!")
	var i, j = 1, "zxz"
	fmt.Println("i=", i, "j=", j)
	fmt.Printf("i= %T\n", i)
	fmt.Printf("j= %T\n", j)

	foo1("aaa", 999)
	r1, r2 := foo2("aaa", 999)
	fmt.Println("r1=", r1, "r2=", r2)
	{
		r1, r2 := foo3("aaa", 999)
		fmt.Println("r1=", r1, "r2=", r2)
	}

}
