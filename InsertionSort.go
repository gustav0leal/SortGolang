package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}


func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Array original:", arr)

	insertionSort(arr)
	fmt.Println("Array ordenado:", arr)
}
