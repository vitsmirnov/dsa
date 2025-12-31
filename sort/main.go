package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func isSorted(arr []int) bool {
	for i, l := 1, len(arr); i < l; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func testSort(sort func([]int), sortName string, arrLen int, maxNum int, loopCount int) {
	arr := make([]int, arrLen)
	t := time.Now()
	for range loopCount {
		for i := range arr {
			arr[i] = rand.IntN(maxNum*2+1) - maxNum
		}
		if isSorted(arr) {
			fmt.Println("random array is sorted!")
		}
		sort(arr)
		if !isSorted(arr) {
			fmt.Println("sort does not work!")
		}
	}
	// funcName := runtime.FuncForPC(uintptr(runtime.FuncForPC))
	fmt.Printf("time (%v): %v\n", sortName, time.Since(t))
}

func testSortDefault(sort func([]int), sortName string) {
	const arrLen int = 100_000
	const maxNum int = 1_000_000

	testSort(sort, sortName, arrLen, maxNum, 100)
}

func main() {
	// arr := []int{3, 7, 9, 1, 2, 5, 4, 3, 2}
	// arr := []int{7, 5, 1, 9, 4, 6, 8, 3, 2}
	// arr := []int{5, 3, 7, 0, 1, 2, 9, 8, 4, 6}
	// fmt.Println(arr)
	// QuickSort(arr)
	// HeapSort(arr)
	// MergeSort(arr)
	// fmt.Println(arr)

	testSortDefault(HeapSort, "heap")
	testSortDefault(MergeSort, "merge")
	testSortDefault(QuickSort, "quick")
}
