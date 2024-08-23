package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := arr[right]
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}

func main() {
	arr := []int{5, 8, 3, 2, 9, 6}
	fmt.Println(quickSort(arr))
}
