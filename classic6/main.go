package main

import (
	"fmt"
)

func main() {
	arr1 := [7]int{10, 11, 12, 13, 14, 15, 16}
	myslice := arr1[0:]
	myslice[0] = 2

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("arr =%v\n", arr1)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	arr2 := [6]int{10, 11, 12, 13, 14, 15}
	myslice2 := make([]int, len(arr1))
	copy(myslice2, arr2[:])

}
