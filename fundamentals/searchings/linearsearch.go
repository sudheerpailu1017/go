package searchings

import "fmt"

func LinearSearch(x int) {
	arr := []int{20, 35, 45, 2, 456, 643, 52, 454356, 234}
	found := false
	for i, j := range arr {
		if x == j {
			fmt.Printf("Element %v is found at %v \n", x, i)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Element not found")
	}

}

/*
best case would be o(1) if the element is found at beginning
worst case: o(n) last
average case: o(n) general
*/
