package main

import (
	"fmt"
)

func main() {
	//start := time.Now()
	unsortedInput := []int{7, 3, 1, 9, 4, -90, -1, 5, 9, 0, 0, 10, 4, -9, -100, 60, 44}
	fmt.Println(HeapSort(unsortedInput))
	// var addInput []*int
	// for i := 0; i < len(unsortedInput); i++ {
	// 	addInput = append(addInput, &unsortedInput[i])
	// }

	// QuickSort(addInput, 0, len(unsortedInput)-1)

	// for i := 0; i < len(addInput); i++ {
	// 	fmt.Printf(" %v ", *addInput[i])

	// }
	fmt.Println("")
	//fmt.Println(BubbleSort(unsortedInput))
	// fmt.Println(HeapSort(unsortedInput))

	// t := time.Now()
	// elapsed := t.Sub(start)
	// fmt.Printf("\n operation that takes %v \n", elapsed)
	// sorted = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
}
