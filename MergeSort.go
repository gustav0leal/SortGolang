package main

import (
	"fmt"
)

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)
	merge(arr, l, m, r)
}

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	left := make([]int, n1)
	right := make([]int, n2)
	copy(left, arr[l:m+1])
	copy(right, arr[m+1:r+1])

	i, j, k := 0, 0, l
	for i < n1 || j < n2 {
		if j == n2 || (i < n1 && left[i] <= right[j]) {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
}


func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Array original:", arr)

	mergeSort(arr, 0, len(arr)-1)
	fmt.Println("Array ordenado:", arr)
}
