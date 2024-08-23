package main

import "fmt"

func quSoet(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := arr[right]
	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quSoet(arr[:left])
	quSoet(arr[left+1:])
	return arr
}

func main() {
	arr := []int{1, 5, 2, 3, 9, 6, 4, 2, 8}
	fmt.Println(quSoet(arr))
}
