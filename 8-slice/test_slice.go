package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slices = %v\n", len(slice1), slice1)

	var slice2 []int
	if slice2 == nil {
		fmt.Println("slice2 == nil")
	} else {
		fmt.Println("slice2 != nil")
	}
	slice2 = make([]int, 2)
	slice2[0] = 100
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	var slice3 []int = make([]int, 2)
	fmt.Printf("len = %d, slice3 = %v\n", len(slice3), slice3)

	slice4 := make([]int, 2)
	fmt.Printf("len = %d, slice4 = %v\n", len(slice4), slice4)

}
