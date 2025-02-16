package loops

import "fmt"

func ForLoops() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is", x)
	// 	x++
	// }
	x := []int{1, 3, 4, 5, 6}
	for _, value := range x {
		fmt.Printf("the value is %v \n", value)
	}
}
