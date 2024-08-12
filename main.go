package main

import (
	"fmt"

	"github.com/Diegovalen47/data-structures/array"
)

func main() {
	myArray1 := []int{1, 2}
	myArray2 := []int{3, 4}

	fmt.Println(array.FindMedianSortedArrays(myArray1, myArray2))
}
