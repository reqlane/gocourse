package basics

import (
	"fmt"
	"slices"
)

func main() {

	// []elementType

	// var slice1 []int
	// var slice2 []int = []int{1, 2, 3}
	// slice3 := []int{9, 8, 7}
	// slice4 := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	slice := a[1:4] // [1,4)
	fmt.Println(slice)

	slice = append(slice, 6, 7)
	fmt.Println("Slice:", slice)

	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("Slice copy:", sliceCopy)

	var nilSlice []int
	fmt.Println(nilSlice)

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index 3 of slice:", slice[3])

	// slice[3] = 50
	// fmt.Println("Element at index 3 of slice:", slice[3])

	if slices.Equal(slice, sliceCopy) {
		fmt.Println("slice is equal to sliceCopy")
	}

	fmt.Println(&slice == &sliceCopy)

	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}
	fmt.Println(twoDSlice)

	// slice[low:high]
	slice2 := slice[2:4] // [2,4)
	fmt.Println(slice2)

	fmt.Println("The capacity of slice2 is:", cap(slice2))
	fmt.Println("The length of slice2 is:", len(slice2))

	fmt.Println(len(slice), cap(slice))
	slice = slices.Delete(slice, 0, 5)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
}
