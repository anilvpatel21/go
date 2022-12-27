package main

import "fmt"

func main() {
	var warr = []int{23, -7, -21, 8, 9}
	//var varr = []int{60, 100, 200}

	NegativeSliding(warr, 3)
	fmt.Println(Kadane(warr))
}
