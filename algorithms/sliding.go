package main

import "fmt"

//FInd out first negative occurance in the window of size k
func NegativeSliding(arr []int, k int) {
	length := len(arr)
	i, j := 0, 0
	narr := []int{}

	for j < length {
		if arr[j] < 0 {
			narr = append(narr, arr[j])
		}

		if j-i+1 == k {
			if len(narr) > 0 {
				fmt.Print(narr[0])
				fmt.Print(" ")
				if arr[i] < 0 {
					narr = narr[1:]
				}
			}

			i++
			j++
		}

		if j-i+1 < k {
			j++
		}
	}
	fmt.Println(" ")
}
