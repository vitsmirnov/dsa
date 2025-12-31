package main

import (
	"cmp"
	"math/rand/v2"
)

func QuickSort[T cmp.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}

	i := partition(arr)
	QuickSort(arr[:i])
	QuickSort(arr[i+1:])
}

func partition[T cmp.Ordered](arr []T) int {
	arrLen := len(arr)
	pivotIndex := rand.IntN(arrLen)
	pivot := arr[pivotIndex]
	arr[pivotIndex], arr[arrLen-1] = arr[arrLen-1], arr[pivotIndex]
	i := 0
	for j := range arrLen - 1 {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[arrLen-1] = arr[arrLen-1], arr[i]
	return i
}

// 5
// 3 7 9 1 2 5 4 3 2
// 3 1 2 2 4 3 5 7 9
//             i   j
