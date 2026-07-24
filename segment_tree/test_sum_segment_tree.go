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
