package main

//Maximum sum sub array problem
func Kadane(arr []int) int {
	length := len(arr)
	maxSum := 0
	if length > 0 {
		maxSum = arr[0]      //it will always be the result for 1 length array
		currentSum := arr[0] //0 for O(n) and arr[0] for O(n-1)
		for i := 1; i < length; i++ {
			if currentSum < 0 { // currentSum > 0 then make it 0 for minimum sub array
				currentSum = 0
			}
			currentSum += arr[i]
			if currentSum > maxSum { //currentSum < maxSum for minimum sub array
				maxSum = currentSum
			}
		}
	}
	return maxSum
}
