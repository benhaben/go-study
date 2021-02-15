package main

import "fmt"

func main() {
	var myArray1 [10]int
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray1 {
		fmt.Println("index=", index, "value=", value)

	}

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	arrSlice := []int{1, 2, 3, 4}
	fmt.Printf("arrSlice types = %T\n", arrSlice)
	printArraySlice(arrSlice)
	for _, value := range arrSlice {
		fmt.Println("value=", value)
	}

	strArrSlice := []string{"1", "a"}
	for _, value := range strArrSlice {
		fmt.Println("value=", value)
	}
}

// 值拷贝，类型只能是4个长度的数组
func printArray(arr [4]int) {
	for index, value := range arr {
		fmt.Println("index=", index, "value=", value)
	}
}

//  引用传递
func printArraySlice(arr []int) {
	arr[0] = 9999
	for index, value := range arr {
		fmt.Println("index=", index, "value=", value)
	}
}
