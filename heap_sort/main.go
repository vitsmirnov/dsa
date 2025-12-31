package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func HeapSort(arr []int) {
	arrLen := len(arr)
	for i := (arrLen - 2) / 2; i >= 0; i-- {
		Sink2(arr, arrLen, i)
	}
	for i := arrLen - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		Sink2(arr, i, 0)
	}
}

func Sink2(arr []int, length int, index int) {
	cur := index
	child := cur*2 + 1 // left
	for done := false; !done && child < length; {
		if right := child + 1; right < length && arr[right] > arr[child] {
			child = right
		}
		done = arr[cur] >= arr[child]
		if !done {
			arr[cur], arr[child] = arr[child], arr[cur]
		}
		cur = child
		child = cur*2 + 1 // left
	}
}

func Sink(arr []int, length int, index int) {
	maxChild := func(p int) int {
		left := p*2 + 1
		right := left + 1
		if left >= length {
			return -1
		} else if right >= length || arr[left] > arr[right] {
			return left
		} else {
			return right
		}
	}

	cur := index
	mc := maxChild(cur)
	for mc != -1 && arr[mc] > arr[cur] {
		arr[cur], arr[mc] = arr[mc], arr[cur]
		cur = mc
		mc = maxChild(cur)
	}
}

func isSorted(arr []int) bool {
	for i, l := 1, len(arr); i < l; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func testSort(sort func([]int)) {
	arr := make([]int, 100_000)
	t := time.Now()
	for range 100 {
		for i := range arr {
			arr[i] = rand.IntN(2_000_001) - 1_000_000
		}
		if isSorted(arr) { // slices.IsSorted(arr) {
			// panic("random array is sorted!")
			fmt.Println("random array is sorted!")
		}
		sort(arr)
		if !isSorted(arr) { // !slices.IsSorted(arr) {
			// panic("quick sort does not work!")
			fmt.Println("sort does not work!")
		}
	}
	fmt.Printf("time: %v\n", time.Since(t))
}

func main() {
	// arr := []int{5, 3, 7, 0, 1, 2, 9, 8, 4, 6}
	// HeapSort(arr)
	// return

	testSort(HeapSort)
}

//         9
//     8         7
//  4     6   2     5
// 0 3  1
