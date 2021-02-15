package main

import (
	"fmt"
	"io"
	"os"
)

//对象的pair会一直传递
func main() {
	//tty: pair<type:*os.File,value:"/dev/tty">
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("Open file error", err)
		return
	}

	// r: pair<type:,value:>
	var r io.Reader
	//tty: pair<type:*os.File,value:"/dev/tty">
	r = tty

	// w: pair<type:,value:>
	var w io.Writer
	//tty: pair<type:*os.File,value:"/dev/tty">
	w = r.(io.Writer)

	w.Write([]byte("hello ...."))
}
