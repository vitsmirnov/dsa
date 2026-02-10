package main

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"time"
)

func isSorted[T cmp.Ordered](arr []T) bool {
	for i := range len(arr) - 1 {
		if arr[i] > arr[i+1] {
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
		// if isSorted(arr) {
		// 	fmt.Println("random array is sorted!")
		// }
		sort(arr)
		if !isSorted(arr) {
			fmt.Printf("%q sort does not work!\n", sortName)
		}
	}
	// funcName := runtime.FuncForPC(uintptr(runtime.FuncForPC))
	// fmt.Printf("time (%v): %v\n", sortName, time.Since(t))
	fmt.Printf("time (%v): %v (len: %v, range: [%v : %v], loops: %v)\n",
		sortName, time.Since(t), arrLen, -maxNum, maxNum, loopCount)
}

func testSortDefault(sort func([]int), sortName string) {
	const arrLen int = 100_000
	const maxNum int = 1_000_000
	const loopCount = 100

	testSort(sort, sortName, arrLen, maxNum, loopCount)
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
	testSortDefault(CountingSort, "counting")
}
