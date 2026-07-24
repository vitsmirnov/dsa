package main

import (
	"math/rand/v2"
)

func GenerateNums(size int, minNum, maxNum int) []int {
	nums := make([]int, size)
	numsRange := maxNum - minNum + 1
	for i := range nums {
		nums[i] = rand.IntN(numsRange) + minNum
	}
	return nums
}

func Sum(arr []int, left, right int) int {
	sum := 0
	for ; left <= right; left++ {
		sum += arr[left]
	}
	return sum
}

func MinRange(nums []int, left, right int) int {
	res := nums[left]
	for left++; left <= right; left++ {
		res = min(res, nums[left])
	}
	return res
}

func MaxRange(nums []int, left, right int) int {
	res := nums[left]
	for left++; left <= right; left++ {
		res = max(res, nums[left])
	}
	return res
}

// todo: cleanup

func main() {
	testSumSegmentTree()
	testMinMaxSegmentTree()
	testSegmentTree()

	// nums := GenerateNums(10, -10000, 10000)
	// n := len(nums)
	// st := MakeMinMaxSegmentTree(nums)
	// fmt.Println(nums)
	// fmt.Println(MinRange(nums, 0, n-1))
	// fmt.Println(MaxRange(nums, 0, n-1))
	// fmt.Println(st.MinMax(0, n-1))
	// fmt.Println(st.items)
}
