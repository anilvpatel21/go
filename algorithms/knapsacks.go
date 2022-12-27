package main

import "math"

//where there is  a choice and sub array problem
func Knapsacks(weight, value []int, sum int) float64 {
	length := len(weight)
	return recursiveKnapsacks(weight, value, length-1, sum)
}

func recursiveKnapsacks(weight, value []int, i, sum int) float64 {
	if i < 0 || sum == 0 {
		return 0
	}

	if weight[i] <= sum {
		return math.Max(float64(value[i])+recursiveKnapsacks(weight, value, i-1, sum-weight[i]), recursiveKnapsacks(weight, value, i-1, sum))
	} else {
		return recursiveKnapsacks(weight, value, i-1, sum)
	}

}
