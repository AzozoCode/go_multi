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

	// Define an array
	arr := [5]int{1, 2, 3, 4, 5}

	// Create a slice referencing arr
	slice1 := arr[1:5] // This points to elements [2, 3, 4]

	// Modify the slice
	slice1[1] = 100 // Changing slice1[1] (which is arr[2])

	fmt.Println("Array:", arr) // [1 2 100 4 5]  (arr is modified)
	fmt.Println("Slice1:", slice1)
}
