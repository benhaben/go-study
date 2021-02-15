package main

import "fmt"

func main() {

	var myMap map[string]string
	if myMap == nil {
		fmt.Println("my map is nil")
	}

	myMap1 := make(map[string]string, 10)
	if myMap1 != nil {
		fmt.Println("my map1 is not nil")
	}

	myMap1["yin"] = "xxx"
	myMap1["one"] = "go"
	fmt.Println("myMap1=", myMap1)

	myMap2 := map[int]string{
		1: "php",
		2: "xxx",
	}
	myMap2[3] = "www"
	fmt.Println("myMap2=", myMap2)

	delete(myMap2, 3)
	for key, v := range myMap2 {
		delete(myMap2, key)
		fmt.Println("key=", key, "value=", v)
	}
	fmt.Println("myMap2 delete=", myMap2)
}
