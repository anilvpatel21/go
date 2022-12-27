package main

import "fmt"

func QuickSort(items []*int, low, high int) {
	if low < high {
		pivot := partition(items, low, high)
		QuickSort(items, low, pivot-1)
		QuickSort(items, pivot+1, high)
	}
}

func partition(items []*int, low, high int) int {
	pivot := *items[low]
	i := low
	j := high
	for i < j {
		//3, 3, 1, 7, 4, 6, 3, -1, -5, 7, 9, 9, 10
		for i < j && *items[i] <= pivot {
			fmt.Println("i", i, *items[i])
			i++

		}
		//i = 3
		for j > 0 && *items[j] > pivot {
			fmt.Println("j", j, *items[j])
			j--

		}

		if i < j {
			fmt.Println("swap i & j", i, j, *items[i], *items[j])
			Swap(items[i], items[j])
		}
	}
	//It is to place pivot (which is at lower) to its perfect/sorted position (j)
	fmt.Println("swap pivot & j", *items[low], *items[j])
	Swap(items[low], items[j])
	return j
}

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
