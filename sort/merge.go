package main

func MergeSort(arr []int) {
	// time: O(n log n), space: O(n)

	buf := make([]int, len(arr))

	var sort func(left, right int)
	sort = func(left, right int) {
		if left < right {
			mid := left + (right-left)/2
			sort(left, mid)
			sort(mid+1, right)
			merge(arr, buf, left, mid, right)
		}
	}

	sort(0, len(arr)-1)
}

func merge(arr, buf []int, left, mid, right int) {
	// time: O(n), space: O(1)

	resLen := right - left + 1
	for i, l, r := 0, left, mid+1; i < resLen; i++ {
		if r > right || (l <= mid && arr[l] < arr[r]) {
			buf[i] = arr[l]
			l++
		} else {
			buf[i] = arr[r]
			r++
		}
	}
	for i := range resLen {
		arr[left] = buf[i]
		left++
	}
}

// temp

func MergeSort0(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	p := len(arr) / 2
	return merge0(MergeSort0(arr[:p]), MergeSort0(arr[p:]))
}

func merge0(arr1, arr2 []int) []int {
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
