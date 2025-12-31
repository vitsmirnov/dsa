package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func MergeSort2(arr []int) []int {
	buf := make([]int, len(arr))

	var sort func(start, end int)
	sort = func(start, end int) {
		if end <= start {
			return
		}

		p := start + (end-start)/2
		sort(start, p)
		sort(p+1, end)
		merge2(arr, buf, start, p+1, end)
	}

	sort(0, len(arr)-1)
	return arr
}

func merge2(arr, buf []int, p1, p2, end int) {
	resLen := end - p1 + 1
	for i, _p1, _p2 := 0, p1, p2; i < resLen; i++ {
		if _p2 > end || (_p1 < p2 && arr[_p1] < arr[_p2]) {
			buf[i] = arr[_p1]
			_p1++
		} else {
			buf[i] = arr[_p2]
			_p2++
		}
	}
	for i := range resLen {
		arr[p1] = buf[i]
		p1++
	}
}

func MergeSort1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	p := len(arr) / 2
	return merge1(MergeSort1(arr[:p]), MergeSort1(arr[p:]))
}

func merge1(arr1, arr2 []int) []int {
	l1, l2 := len(arr1), len(arr2)
	resLen := l1 + l2
	res := make([]int, resLen)
	for i, p1, p2 := 0, 0, 0; i < resLen; i++ {
		if p2 == l2 || (p1 < l1 && arr1[p1] < arr2[p2]) {
			res[i] = arr1[p1]
			p1++
		} else {
			res[i] = arr2[p2]
			p2++
		}
	}
	return res
}

func isSorted(arr []int) bool {
	for i, l := 1, len(arr); i < l; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func test1() {
	const r int = 1_000_000
	arr := make([]int, 100_000)
	t := time.Now()
	for range 100 {
		for i := range arr {
			arr[i] = rand.IntN(r*2+1) - r
		}
		if isSorted(arr) {
			fmt.Println("random array is sorted!")
		}
		arr = MergeSort2(arr)
		if !isSorted(arr) {
			fmt.Println("MergeSort failed!")
		}
	}
	fmt.Printf("time: %v\n", time.Since(t))
}

func main() {
	// arr := []int{7, 5, 1, 9, 4, 6, 8, 3, 2}
	// fmt.Println(arr)
	// fmt.Println(MergeSort2(arr))
	test1()
}
