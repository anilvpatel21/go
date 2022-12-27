package main

import "fmt"

type Heap struct {
	items []int
}
//[20,30,56,78]
func (h *Heap) insertHeap(item int) {
	h.items = append(h.items, item);
	i := len(h.items) - 1;
	for (i > 0) {
		iParent := int((i-1)/2);
		if item > h.items[iParent] {
			swap(&h.items[iParent], &h.items[i]);
			i = iParent;
		} else {
			break;
		}
	}
}

func (h *Heap) deleteHeap() int{
	root := h.items[0];
	length  := len(h.items);
	swap(&h.items[0],&h.items[length - 1])
    h.items = h.items[:length-1];
	i := 0
	length = len(h.items)
	for (i < int(length/2)) {
		left := 2*i + 1
		right := 2*i + 2
		largestIndex := i;
		if left < length {
			if right < length {
				if h.items[left] > h.items[right] && h.items[left] > h.items[i] {
					largestIndex = left
				} else if h.items[right] > h.items[i] {
					largestIndex = right
				}
			} else if h.items[left] > h.items[i] {
				largestIndex = left
			}
		}

		if largestIndex != i {
			swap(&h.items[largestIndex], &h.items[i])
			i = largestIndex
		} else {
			break;
		}
	}
	return root;
}

func swap(a, b *int) {
	temp := *a;
	*a = *b;
	*b = temp;
}

func Mainheap() {
   myHeap := Heap{};
 //  myHeap.items = append(myHeap.items, 0)
   myHeap.insertHeap(5)
   myHeap.insertHeap(8)
   myHeap.insertHeap(4)
   myHeap.insertHeap(7)
   myHeap.insertHeap(-1)
   myHeap.insertHeap(6)
   myHeap.insertHeap(16)
   myHeap.insertHeap(9)
   myHeap.insertHeap(14)
   myHeap.insertHeap(10)
   fmt.Println(myHeap)
   fmt.Println(myHeap.deleteHeap())
   fmt.Println(myHeap.deleteHeap())
   fmt.Println(myHeap)
}