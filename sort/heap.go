package main

func HeapSort(arr []int) {
	arrLen := len(arr)
	for i := (arrLen - 2) / 2; i >= 0; i-- {
		Sink(arr, arrLen, i)
	}
	for i := arrLen - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		Sink(arr, i, 0)
	}
}

func Sink(arr []int, length int, index int) {
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

func Sink0(arr []int, length int, index int) {
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

//         9
//     8         7
//  4     6   2     5
// 0 3  1
