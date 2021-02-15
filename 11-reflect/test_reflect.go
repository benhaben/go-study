package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("call")
	fmt.Printf("%v", this)
}

func (this *User) Call1() {
	fmt.Println("call")
	fmt.Printf("%v", this)
}

func (this User) Call2() {
	fmt.Println("call")
	fmt.Printf("%v", this)
}

func reflectNum(arg interface{}) {
	fmt.Println("type=", reflect.TypeOf(arg))
	fmt.Println("value=", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)

	user := User{
		1, "yin", 33,
	}
	user.Call()
	DoFieldAndMethod(user)
	DoFieldAndMethod1(&user)
}

func DoFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("type=", inputType)

	v := reflect.ValueOf(input)
	fmt.Println("value=", v)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("%s:%v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}
}

func DoFieldAndMethod1(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("type=", inputType)

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}
}
