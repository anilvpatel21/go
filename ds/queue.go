package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items,i)
}

func (q *Queue) Dequeue() int {
	l := len(q.items);
	toRemove := q.items[0]
	q.items = q.items[1:l]
	return toRemove
}

func Mainqueue() {
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(30)
	myQueue.Enqueue(70)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}