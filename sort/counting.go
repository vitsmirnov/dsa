package main

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func CountingSort[T Integer](nums []T) {
	// time: O(n+r), space: O(r)
	// n - len(nums)
	// r - nums range (max(nums) - min(nums))

	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums {
		minNum = min(minNum, num)
		maxNum = max(maxNum, num)
	}
	numFreq := make([]int, maxNum-minNum+1)
	for _, num := range nums {
		numFreq[num-minNum]++
	}
	i := 0
	for num, count := range numFreq {
		_num := T(num) + minNum
		for ; count != 0; count-- {
			nums[i] = _num
			i++
		}
	}
}
