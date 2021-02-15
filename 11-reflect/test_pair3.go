package main

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {

}

func (this *Book) WriteBook() {

}
func main() {
	p := &Book{}

	var r Reader
	r = p
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()
}
