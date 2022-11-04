package main

import (
	"fmt"
)

var (
	pl      = fmt.Println
	fstring = fmt.Printf
)

func main() {
	anums := []int{1, 2, 3, 4}
	for idx, num := range anums {
		fstring("%d : %d\n", idx, num)
	}
}
