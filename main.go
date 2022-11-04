package main

import (
	"fmt"
)

func main() {
	// declaring our first array
	// type 1 just declare a fixed size array without
	// initialisation
	var arr1 [5]int
	//  |    |   |
	//  name size  data type

	// assign values to an array
	arr1[0] = 4
	arr1[1] = 2
	arr1[2] = 0
	arr1[3] = 6
	arr1[4] = 9

	// type 2 : declare and initialise at the same time
	arr2 := [5]int{7, 3, 3, 7, 1}
	// accessing elements by index
	fmt.Println("idx 0 : ", arr2[0])
	fmt.Println("idx 1 : ", arr2[1])
	// get lenght of the array
	fmt.Println("length of the array : ", len(arr2))

	for i := 0; i < len(arr2); i++ {
		fmt.Printf("%d ", arr2[i])
	}
	fmt.Println()
	// range based iteration
	for _, val := range arr1 {
		fmt.Printf("%d ", val)
	}

	fmt.Println()
	// creating multi-dim arrays
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	// print mult dim array
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[0]); j++ {
			fmt.Printf("%d ", arr3[i][j])
		}
		fmt.Println()
	}

	fmt.Println()
}
