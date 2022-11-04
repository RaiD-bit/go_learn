package main

import (
	"fmt"
)

func main() {
	/*
	   slices are arrays but they can grow in size
	   you can make a slice like this
	     -> var name []datatype
	   but we generally use the make function to create slices.
	     -> sl1 := make([]string, 6)
	*/
	sl1 := make([]string, 6)
	sl1[0] = "devashish"
	sl1[1] = "normandy"
	sl1[2] = "panipat"
	sl1[3] = "versaille"
	sl1[4] = "istanbul"
	sl1[5] = "ceylon"
	fmt.Println("slice size : ", len(sl1))
	for i := 0; i < len(sl1); i++ {
		fmt.Println(sl1[i])
	}
	for _, val := range sl1 {
		fmt.Println(val)
	}
	// slice is basically points at an array so you can create an array of an array
	// let's take an example
	normalArray := [5]int{
		1, 2, 3, 4, 5,
	}
	sl2 := normalArray[2:5]
	fmt.Println(sl2)

	// append a value like so
	sl3 := normalArray[0:2]

	sl3 = append(sl3, 12)

	fmt.Println(sl3)
	fmt.Println(normalArray)
}
