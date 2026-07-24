package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

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

// todo: cleanup

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
