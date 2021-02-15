package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) show() {
	fmt.Println("Name=", this.Name)
	fmt.Println("Ad=", this.Ad)
	fmt.Println("Level=", this.Level)
	fmt.Println("Hero=", this)

}

func (this *Hero) getName() string {
	return this.Name
}

func (this *Hero) setName(name string) {
	this.Name = name
}

func main() {

	hero := Hero{
		Name:  "yin",
		Ad:    0,
		Level: 1,
	}

	hero.show()
	hero.setName("xiong")
	hero.show()
}
