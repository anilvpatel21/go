package main

import "fmt"

func BubbleSort(items []int) []int{
	length := len(items)
	fmt.Println(items)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if items[j] < items[j+1] {
				Swap(&items[j], &items[j+1])
			}
		}
		fmt.Println(items)
	}
	return items
}