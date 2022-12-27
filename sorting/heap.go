package main

func HeapSort(items []int) []int {
	length := len(items)
	//Heapify means ... arrange array into a tree structure
	//In a max heap, the value of a node is always greater than or equal to the value of each of its children. Conversely, in a min heap, the value of a parent is always <= the value of each of its children.
	//Build a max heap to sort in increasing order, build a min heap to sort in decreasing order.

	MaxHeapify(items, length)
	for i := length - 1; i >= 1; i-- {
		Swap(&items[i], &items[0])
		heapifyProcess(items, i, 0)
	}
	return items
}

func MaxHeapify(items []int, length int) {
	for i := int(length/2) - 1; i >= 0; i-- {
		heapifyProcess(items, length, i)
	}
}

func heapifyProcess(items []int, length int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && items[left] > items[largest] {
		largest = left
	}

	if right < length && items[right] > items[largest] {
		largest = right
	}

	if largest != i {
		Swap(&items[largest], &items[i])
		heapifyProcess(items, length, largest)
	}
}
