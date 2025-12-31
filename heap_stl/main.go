package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func main() {
	const arrLen int = 1_000_000
	const maxNum int = 1_000_000
	arr := make([]int, arrLen)
	for i := range arr {
		arr[i] = rand.IntN(maxNum*2+1) - maxNum
	}
	// fmt.Println(arr)
	if slices.IsSorted(arr) {
		fmt.Println("Random array is sorted!")
	}

	pq := NewPriorityQueue[int](Less)
	for _, num := range arr {
		pq.PushVal(num)
	}

	arr2 := make([]int, arrLen)
	for i := range arr2 {
		arr2[i] = pq.PopVal()
	}
	// fmt.Println(arr2)
	if !slices.IsSorted(arr2) {
		fmt.Println("Sort by heap failed, array is not sorted!")
	}
}
