package main

func MergeSort(arr []int) {
	buf := make([]int, len(arr))

	var sort func(start, end int)
	sort = func(start, end int) {
		if end <= start {
			return
		}

		p := start + (end-start)/2
		sort(start, p)
		sort(p+1, end)
		merge(arr, buf, start, p+1, end)
	}

	sort(0, len(arr)-1)
}

func merge(arr, buf []int, p1, p2, end int) {
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
