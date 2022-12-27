package main

func SelectionSort(items []int) []int {
	length := len(items)
	for i := 0; i < length-1; i++ {
		position := i
		//7, 3, 1, 9, 4, 9, 3, -1, -5, 3, 6, 7, 10
		for j := i + 1; j < length; j++ {
			if items[j] < items[position] {
				position = j
			}
		}
		Swap(&items[i], &items[position])
		if i == 1 {
			break
		}
	}
	return items
}
