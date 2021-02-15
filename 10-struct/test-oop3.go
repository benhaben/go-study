package main

import "fmt"

//本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

//dog

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(p AnimalIF) {
	p.Sleep()
	fmt.Println("color=", p.GetColor())
	fmt.Println("type=", p.GetType())

}
func main() {
	var animal AnimalIF
	animal = &Cat{
		"green",
	}
	showAnimal(animal)

	animal = &Dog{
		"yellow",
	}
	showAnimal(animal)
}
