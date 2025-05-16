package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			comparisons += 1
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		time.Sleep(delay * 10)
	}
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Array original:", arr)

	selectionSort(arr)
	fmt.Println("Array ordenado:", arr)
}
