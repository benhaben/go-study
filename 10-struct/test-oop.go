package main

import "fmt"

//声明一种新类型
type myInt int

//定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	book.auth = "666"
}

func changeBook1(book *Book) {
	book.auth = "666"
}

func main() {
	var book Book
	book.auth = "yin"
	book.title = "Mastering blockchian"
	fmt.Printf("Book = %v\n", book)

	changeBook(book)
	fmt.Printf("Book = %v\n", book)
	changeBook1(&book)
	fmt.Printf("Book = %v\n", book)

}
