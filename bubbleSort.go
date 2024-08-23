package main

import "fmt"

func bubbleSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 5, 3, 8, 2, 6, 9, 1}
	fmt.Println(bubbleSort(arr))
}
