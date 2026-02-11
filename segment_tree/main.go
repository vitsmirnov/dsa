package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

type SSTTest struct {
	nums           []int
	st             *SumSegmentTree
	maxNum, minNum int
}

func MakeSSTTest(numsSize int, minNum, maxNum int) *SSTTest {
	nums := GenerateNums(numsSize, minNum, maxNum)
	return &SSTTest{
		nums:   nums,
		st:     MakeSumSegmentTree(nums),
		maxNum: maxNum,
		minNum: minNum}
}

func (t *SSTTest) TestSum(left, right int) bool {
	return sum(t.nums, left, right) == t.st.Sum(left, right)
}

func (t *SSTTest) Update(index int, value int) {
	t.nums[index] = value
	t.st.Update(index, value)
}

func (t *SSTTest) Update2() {
	index := rand.IntN(len(t.nums))
	value := rand.IntN(t.maxNum-t.minNum+1) + t.minNum
	t.nums[index] = value
	t.st.Update(index, value)
}

type MMSTTest struct {
	nums           []int
	st             *MinMaxSegmentTree
	maxNum, minNum int
}

func MakeMMSTTest(numsSize int, minNum, maxNum int) *MMSTTest {
	nums := GenerateNums(numsSize, minNum, maxNum)
	return &MMSTTest{
		nums:   nums,
		st:     MakeMinMaxSegmentTree(nums),
		maxNum: maxNum,
		minNum: minNum}
}

func (t *MMSTTest) TestMinMax(left, right int) bool {
	item := t.st.MinMax(left, right)
	return MinRange(t.nums, left, right) == item.min &&
		MaxRange(t.nums, left, right) == item.max
}

func (t *MMSTTest) Update(index int, value int) {
	t.nums[index] = value
	t.st.Update(index, value)
}

func GenerateNums(size int, minNum, maxNum int) []int {
	nums := make([]int, size)
	numsRange := maxNum - minNum + 1
	for i := range nums {
		nums[i] = rand.IntN(numsRange) + minNum
	}
	return nums
}

func sum(arr []int, left, right int) int {
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

func testSumSegmentTree() {
	const numsLen = 1_000_000
	const minNum = -1_000_000
	const maxNum = 1_000_000
	const loopCount = 1_000

	t := time.Now()
	test := MakeSSTTest(numsLen, minNum, maxNum)
	if !slices.Equal(test.nums, test.st.Nums()) {
		fmt.Println("nums and SST are not equal")
	}
	for range loopCount {
		left := rand.IntN(numsLen)
		right := left + rand.IntN(numsLen-left)
		if !test.TestSum(left, right) {
			fmt.Println("sums are not equal")
		}

		index := rand.IntN(numsLen)
		value := rand.IntN(maxNum-minNum+1) + minNum
		test.Update(index, value)
		if !test.TestSum(index, index) {
			fmt.Println("sums are not equal after update")
		}
	}
	fmt.Printf("testSumSegmentTree() time: %v\n", time.Since(t))
}

func testMinMaxSegmentTree() {
	const numsLen = 1_000_000
	const minNum = -1_000_000
	const maxNum = 1_000_000
	const loopCount = 1_000

	t := time.Now()
	test := MakeMMSTTest(numsLen, minNum, maxNum)
	if !slices.Equal(test.nums, test.st.Nums()) {
		fmt.Println("nums and MMST are not equal")
	}
	for range loopCount {
		left := rand.IntN(numsLen)
		right := left + rand.IntN(numsLen-left)
		if !test.TestMinMax(left, right) {
			fmt.Println("min/max are not equal")
		}

		index := rand.IntN(numsLen)
		value := rand.IntN(maxNum-minNum+1) + minNum
		test.Update(index, value)
		if !test.TestMinMax(index, index) {
			fmt.Println("min/max are not equal after update")
		}
	}
	fmt.Printf("testMinMaxSegmentTree() time: %v\n", time.Since(t))
}

func main() {
	testSumSegmentTree()
	testMinMaxSegmentTree()

	// nums := GenerateNums(10, -10000, 10000)
	// n := len(nums)
	// st := MakeMinMaxSegmentTree(nums)
	// fmt.Println(nums)
	// fmt.Println(MinRange(nums, 0, n-1))
	// fmt.Println(MaxRange(nums, 0, n-1))
	// fmt.Println(st.MinMax(0, n-1))
	// fmt.Println(st.items)
}
