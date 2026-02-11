package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

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

type STTest struct {
	nums           []int
	st             *SegmentTree
	maxNum, minNum int
}

func MakeSTTest(numsSize int, minNum, maxNum int) *STTest {
	nums := GenerateNums(numsSize, minNum, maxNum)
	return &STTest{
		nums:   nums,
		st:     MakeSegmentTree(nums),
		maxNum: maxNum,
		minNum: minNum}
}

func (t *STTest) TestSum(left, right int) bool {
	return sum(t.nums, left, right) == t.st.Sum(left, right)
}

func (t *STTest) Update(index int, value int) {
	t.nums[index] = value
	t.st.Update(index, value)
}

func (t *STTest) Update2() {
	index := rand.IntN(len(t.nums))
	value := rand.IntN(t.maxNum-t.minNum+1) + t.minNum
	t.nums[index] = value
	t.st.Update(index, value)
}

// todo: cleanup

func testSegmentTree() {
	const numsLen = 1_000_000
	const minNum = -1_000_000
	const maxNum = 1_000_000

	t := time.Now()
	test := MakeSTTest(numsLen, minNum, maxNum)
	if !slices.Equal(test.nums, test.st.Nums()) {
		fmt.Println("nums and ST are not equal")
	}
	for range 10000 {
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
	fmt.Printf("time: %v\n", time.Since(t))
}

func main() {
	testSegmentTree()
}
