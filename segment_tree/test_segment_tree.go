package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

type STTest struct {
	nums           []int
	st             *SegmentTree[MinMax, int]
	maxNum, minNum int
}

func MakeSTTest(numsSize int, minNum, maxNum int) *STTest {
	nums := GenerateNums(numsSize, minNum, maxNum)
	initNode := func(val int) MinMax {
		return MinMax{min: val, max: val}
	}
	buildNode := func(leftChild, rightChild MinMax) MinMax {
		return MinMax{
			min: min(leftChild.min, rightChild.min),
			max: max(leftChild.max, rightChild.max)}
	}
	return &STTest{
		nums:   nums,
		st:     MakeSegmentTree(nums, initNode, buildNode),
		maxNum: maxNum,
		minNum: minNum}
}

func (t *STTest) TestQuery(left, right int) bool {
	node := t.st.Query(left, right)
	return MinRange(t.nums, left, right) == node.min &&
		MaxRange(t.nums, left, right) == node.max
}

func (t *STTest) Update(index int, value int) {
	t.nums[index] = value
	t.st.Update(index, value)
}

// todo: cleanup

func testSegmentTree() {
	const numsLen = 1_000_000
	const minNum = -1_000_000
	const maxNum = 1_000_000
	const loopCount = 1_000

	t := time.Now()
	test := MakeSTTest(numsLen, minNum, maxNum)
	if !slices.Equal(test.nums, test.st.Items(func(node MinMax) int { return node.min })) {
		fmt.Println("nums and ST are not equal")
	}
	for range loopCount {
		left := rand.IntN(numsLen)
		right := left + rand.IntN(numsLen-left)
		if !test.TestQuery(left, right) {
			fmt.Println("test query failed")
		}

		index := rand.IntN(numsLen)
		value := rand.IntN(maxNum-minNum+1) + minNum
		test.Update(index, value)
		if !test.TestQuery(index, index) {
			fmt.Println("test query failed after update")
		}
	}
	fmt.Printf("testSegmentTree() time: %v\n", time.Since(t))
}
