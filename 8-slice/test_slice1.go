package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d,cap=%d,slices=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len=%d,cap=%d,slices=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len=%d,cap=%d,slices=%v\n", len(numbers), cap(numbers), numbers)

	//已经满了，会自动增加cap容量
	numbers = append(numbers, 3)
	fmt.Printf("len=%d,cap=%d,slices=%v\n", len(numbers), cap(numbers), numbers)

	// 引用是相同的元素
	s := numbers[:]
	s[1] = 999
	fmt.Printf("len=%d,cap=%d,slices=%v\n", len(s), cap(s), s)
	fmt.Printf("len=%d,cap=%d,slices=%v\n", len(numbers), cap(numbers), numbers)

	s1 := make([]int, len(numbers))
	copy(s1, numbers)
	s1[1] = 111
	fmt.Printf("len=%d,cap=%d,slices=%v\n", len(s1), cap(s1), s1)
	fmt.Printf("len=%d,cap=%d,slices=%v\n", len(numbers), cap(numbers), numbers)

}
