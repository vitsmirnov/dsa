package main

func HeapSort(arr []int) {
	// time: O(n log n), space: O(1)

	arrLen := len(arr)
	for i := arrLen/2 - 1; i >= 0; i-- { // (arrLen - 2) / 2 ?
		sink(arr, arrLen, i)
	}
	for i := arrLen - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		sink(arr, i, 0)
	}
}

func sink(arr []int, length int, index int) {
	// time: O(log n), space: O(1)

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

// todo: cleanup

func down(nums []int, numsLen, index int) {
	// time: O(log n), space: O(1)

	child := index*2 + 1 // left
	for child < numsLen {
		if child+1 < numsLen && nums[child+1] > nums[child] {
			child++
		}
		if nums[child] <= nums[index] {
			return
		}
		nums[index], nums[child] = nums[child], nums[index]
		index = child
		child = index*2 + 1 // left
	}
}

func up(nums []int, index int) {
	// time: O(log n), space: O(1)
	// ?

	for parent := (index - 1) / 2; nums[parent] > nums[index]; {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) / 2
	}
}

func sink0(arr []int, length int, index int) {
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
