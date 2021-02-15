package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) show() {
	fmt.Println("Name=", this.name)
	fmt.Println("sex=", this.sex)
}

func (this *Human) Eat() {
	fmt.Println("Human eat")

}

func (this *Human) Walk() {
	fmt.Println("Human walk")
}

func (this *Human) getName() string {
	return this.name
}

func (this *Human) setName(name string) {
	this.name = name
}

type SuperMan struct {
	Human // SuperMan继承Human
	level int
}

//重定义
func (*SuperMan) Eat() {
	fmt.Println("SuperMan eat")
}
func (*SuperMan) Walk() {
	fmt.Println("SuperMan Walk")
}

//定义新方法
func (*SuperMan) Fly() {
	fmt.Println("SuperMan Fly")
}
func main() {

	hero := Human{
		name: "yin",
		sex:  "male",
	}
	hero.Eat()
	super := SuperMan{
		Human{
			"yin1",
			"male1",
		},
		1,
	}
	super.Eat()
	super.Fly()

	var s1 SuperMan
	s1.sex = "s1"
	s1.name = "s1"
	s1.level = 1
	s1.show()
}
