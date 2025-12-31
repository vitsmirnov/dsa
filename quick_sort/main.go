package main

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"time"
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

func isSorted(arr []int) bool {
	for i, l := 1, len(arr); i < l; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func testSort() {
	arr := make([]int, 100_000)
	t := time.Now()
	for range 100 {
		for i := range arr {
			arr[i] = rand.IntN(2_000_001) - 1_000_000
		}
		if isSorted(arr) { // slices.IsSorted(arr) {
			panic("random array is sorted!")
		}
		QuickSort(arr)
		if !isSorted(arr) { // !slices.IsSorted(arr) {
			panic("quick sort does not work!")
		}
	}
	fmt.Printf("time: %v\n", time.Since(t))
}

func main() {
	// arr := []int{3, 7, 9, 1, 2, 5, 4, 3, 2}
	// fmt.Println(arr)
	// quickSort(arr)
	// fmt.Println(arr)

	testSort()
}

// 5
// 3 7 9 1 2 5 4 3 2
// 3 1 2 2 4 3 5 7 9
//             i   j
