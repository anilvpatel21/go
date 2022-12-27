package main

func InsertionSort(items []int) []int {

	for i := 1; i < len(items); i++ {
		//7, 3, 1, 9, 4, 9, 3, -1, -5, 3, 6, 7, 10
		ele := items[i] //3
		k := 0
		for ele > items[k] { //3 > 7
			k++
		}

		for j := i; j > k; j-- { //j = 1; 1 > 2;
			items[j] = items[j-1]
		}

		items[k] = ele //
	}

	return items
}
