package arrays

import "fmt"

var A = []int{1000, 3000, 5, 6}

func ArraysDeclaration() {
	var nums = [3]int{20, 1000, 80}
	fmt.Printf("The array is: %v\n", nums)
}

func SliceDeclaration() {
	fmt.Println("the slice is", A)
}

func Append_slice(slice_name []int, element int) []int {
	slice_name = append(slice_name, element)
	fmt.Println(slice_name)
	return slice_name
}
