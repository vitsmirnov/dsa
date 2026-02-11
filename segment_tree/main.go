package main

import (
	"fmt"
	"math/rand/v2"
)

type SegmentTree struct{ sums []int }

func MakeSegmentTree(arr []int) *SegmentTree {
	sums := make([]int, len(arr)*4)
	buildST(arr, sums, 0, 0, len(arr)-1)
	return &SegmentTree{sums: sums}
}

func buildST(arr, sums []int, pos int, segLeft, segRight int) {
	if segLeft == segRight {
		sums[pos] = arr[segLeft]
	} else {
		mid := segLeft + (segRight-segLeft)/2
		leftChild, rightChild := pos*2+1, pos*2+2
		buildST(arr, sums, leftChild, segLeft, mid)
		buildST(arr, sums, rightChild, mid+1, segRight)
		sums[pos] = sums[leftChild] + sums[rightChild]
	}
}

func sumST(sums []int, pos int, segLeft, segRight, qLeft, qRight int) int {
	if qLeft > qRight {
		return 0
	}
	if segLeft == qLeft && segRight == qRight {
		return sums[pos]
	}
	mid := segLeft + (segRight-segLeft)/2
	return sumST(sums, pos*2+1, segLeft, mid, qLeft, min(qRight, mid)) +
		sumST(sums, pos*2+2, mid+1, segRight, max(qLeft, mid+1), qRight)
}

func updateST(sums []int, pos int, segLeft, segRight int, index int, value int) {
	if segLeft == segRight {
		sums[pos] = value
	} else {
		mid := segLeft + (segRight-segLeft)/2
		if index <= mid {
			updateST(sums, pos*2+1, segLeft, mid, index, value)
		} else {
			updateST(sums, pos*2+2, mid+1, segRight, index, value)
		}
		sums[pos] = sums[pos*2+1] + sums[pos*2+2]
	}
}

func (st *SegmentTree) Sum(left, right int) int {
	return sumST(st.sums, 0, 0, len(st.sums)/4-1, left, right)
}

func (st *SegmentTree) Update(index int, value int) {
	updateST(st.sums, 0, 0, len(st.sums)/4-1, index, value)
}

// temp

func sum(arr []int, left, right int) int {
	sum := 0
	for ; left <= right; left++ {
		sum += arr[left]
	}
	return sum
}

func GenerateNums(size int, minNum, maxNum int) []int {
	nums := make([]int, size)
	numsRange := maxNum - minNum + 1
	for i := range nums {
		nums[i] = rand.IntN(numsRange) + minNum
	}
	return nums
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

func (t *STTest) Update2() { //index int, value int) {
	index := rand.IntN(len(t.nums))
	value := rand.IntN(t.maxNum-t.minNum+1) + t.minNum
	t.nums[index] = value
	t.st.Update(index, value)
}

// todo: cleanup

func testSegmentTree() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// n := len(nums)
	// st := MakeSegmentTree(nums)
	// fmt.Println(nums)
	// fmt.Println(st.Sum(0, n-1))
	// fmt.Println(sum(nums, 0, n-1))
	// st.

	// nums2 := Nums{}
	// nums2.Generate(10, 0, 9)
	// fmt.Println(nums2.nums)

	const numsLen = 1_000_000
	const minNum = -1_000_000
	const maxNum = 1_000_000

	test := MakeSTTest(numsLen, minNum, maxNum)
	for range 1000 {
		left := rand.IntN(numsLen)
		right := left + rand.IntN(numsLen-left)
		// fmt.Println(left, right)
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
}

func main() {
	testSegmentTree()
}
