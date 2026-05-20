package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func MakeMinSparseTable(nums []int) *SparseTable {
	return MakeSparseTable(nums, func(a, b int) int { return min(a, b) })
}
func MakeMaxSparseTable(nums []int) *SparseTable {
	return MakeSparseTable(nums, func(a, b int) int { return max(a, b) })
}

type SparseTable struct {
	mins [][]int
	logs []int
	f    func(int, int) int
}

func MakeSparseTable(nums []int, f func(int, int) int) *SparseTable {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}

	logs := make([]int, numsLen+1)
	logs[1] = 0
	for i := 2; i <= numsLen; i++ {
		logs[i] = logs[i>>1] + 1
	}

	// p := 0
	// for (1 << p) <= numsLen {
	// 	p++
	// }
	p := logs[numsLen] + 1
	mins := make([][]int, p)
	mins[0] = make([]int, numsLen)
	copy(mins[0], nums)
	for i := 1; i < p; i++ {
		mins[i] = make([]int, numsLen)
		prevLevel := i - 1
		prevLevelLen := 1 << prevLevel
		for j := numsLen - (1 << i); j >= 0; j-- {
			mins[i][j] = f(mins[prevLevel][j], mins[prevLevel][j+prevLevelLen])
		}
	}

	return &SparseTable{
		mins: mins,
		logs: logs,
		f:    f}
}

// min/max
func (st *SparseTable) Query(left, right int) int {
	l := st.logs[right-left+1]
	return st.f(st.mins[l][left], st.mins[l][right-(1<<l)+1])
}

// sum would be like that
func (st *SparseTable) Sum(left, right int) int {
	sum := 0
	length := right - left + 1
	for i := 0; length != 0; i++ {
		if length&1 == 1 {
			sum += st.mins[i][left]
			left += 1 << i
		}
		length >>= 1
	}
	return sum
}

//   l       r
// 1 2 3 4 5 6 7
// length: 5
// 101

func testSparseTable() {
	const minNum int = -1e5
	const maxNum int = 1e5
	const numsRange = maxNum - minNum + 1
	const numsLen int = 1e5
	const testCount int = 100
	const queryCount int = 100
	// funcs := [2]func(int, int) int{min, max}
	funcs := [2]func(int, int) int{
		func(a, b int) int { return min(a, b) },
		func(a, b int) int { return max(a, b) }}
	nums := make([]int, numsLen)
	t := time.Now()
	for range testCount {
		for i := range nums {
			nums[i] = rand.IntN(numsRange) + minNum
		}
		f := funcs[rand.IntN(len(funcs))]
		st := MakeSparseTable(nums, f)
		for range queryCount {
			length := rand.IntN(numsLen-1) + 1
			left := rand.IntN(numsLen)
			right := min(left+length-1, numsLen-1)
			res1 := st.Query(left, right)
			res2 := agr(nums[left:right+1], f)
			if res1 != res2 {
				fmt.Printf("%v != %v: [%v, %v]\n", res1, res2, left, right)
			}
		}
	}
	fmt.Printf("testSparseTable time: %v\n", time.Since(t))
}

func testSparseTableSum() {
	const minNum int = -1e2
	const maxNum int = 1e2
	const numsRange = maxNum - minNum + 1
	const numsLen int = 1e5
	const testCount int = 100
	const queryCount int = 500
	nums := make([]int, numsLen)
	sum := func(a, b int) int { return a + b }
	t := time.Now()
	for range testCount {
		for i := range nums {
			nums[i] = rand.IntN(numsRange) + minNum
		}
		st := MakeSparseTable(nums, sum)
		for range queryCount {
			length := rand.IntN(numsLen-1) + 1
			left := rand.IntN(numsLen)
			right := min(left+length-1, numsLen-1)
			res1 := st.Sum(left, right)
			res2 := agr(nums[left:right+1], sum)
			if res1 != res2 {
				fmt.Printf("%v != %v: [%v, %v]\n", res1, res2, left, right)
			}
		}
	}
	fmt.Printf("testSparseTableSum time: %v\n", time.Since(t))
}

func agr(nums []int, f func(int, int) int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		res = f(res, num)
	}
	return res
}

func main() {
	// logs := make([]int, 22)
	// logs[1] = 0
	// for i := 2; i < len(logs); i++ {
	// 	logs[i] = logs[i>>1] + 1
	// }
	// for i, log := range logs {
	// 	fmt.Printf("%v: %v\n", i, log)
	// }

	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// numsLen := len(nums)
	// st := MakeSparseTable(nums, func(a, b int) int { return min(a, b) })
	// fmt.Println(nums)
	// fmt.Println(st.Query(0, numsLen-1))
	// return

	testSparseTable()
	testSparseTableSum()
}

//            l             r
//         0  1  2  3  4  5 6
// 0 (1):  1  2  3  4  5  6 7
// 1 (2):  3  5  7  9 11 13
// 2 (4): 10 14 18 22
// 3 (8): X

// ln
// 1: 0
// 2: 1
// 3: 1
// 4: 2
// 5: 2
// 6: 2
// 7: 2
// 8: 3
// 9: 3
// 10: 3
// 11: 3
// 12: 3
// 13: 3
// 14: 3
// 15: 3
// 16: 4

// 0..2^3-1
// 0..7
// 0..2^(3-1)-1 + 2^(3-1)..2^3-1
// 0..2^2-1 + 2^2..2^3-1
// 0..3 + 4..7

// 0..2^4-1
// 0..2^3-1 + 2^3..2^4-1

// 0..7
// 0..1 + 2..3 + 4..7

// 0..15
// 0..1 + 2..3 + 4..7 + 8..15

// 22 = 2 + 2^2 + 2^4 = 2 + 4 + 16 = 22
// 10110

// 2^i
// 2^(i-1)

// 011
// 101
// 110
