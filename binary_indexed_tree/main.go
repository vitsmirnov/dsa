/*

Binary Indexed Tree / Fenwick Tree

*/

package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type BIT struct{ sums []int }

func MakeBIT(size int) BIT { return BIT{sums: make([]int, size)} }
func MakeBITFromNums(nums []int) BIT {
	size := len(nums)
	sums := make([]int, size)
	for i, num := range nums {
		sums[i] += num
		if j := h(i); j < size {
			sums[j] += sums[i]
		}
	}
	return BIT{sums: sums}
}

func (bit *BIT) Item(index int) int { return bit.Sum(index, index) }
func (bit *BIT) Sum(left, right int) int {
	if left == 0 {
		return bit.sum(right)
	}
	return bit.sum(right) - bit.sum(left-1)
}
func (bit *BIT) sum(right int) int {
	sum := 0
	for right >= 0 {
		sum += bit.sums[right]
		right = g(right) - 1
	}
	return sum
}
func (bit *BIT) Update(index int, delta int) {
	for size := len(bit.sums); index < size; index = h(index) {
		bit.sums[index] += delta
	}
}
func (bit BIT) String() string {
	if len(bit.sums) == 0 {
		return "[]"
	}
	res := make([]byte, 0, len(bit.sums)*2+2)
	res = append(res, '[')
	for i := range bit.sums {
		res = fmt.Appendf(res, "%v ", bit.Item(i))
	}
	res[len(res)-1] = ']'
	return string(res)
}

// func (bit *BIT) Assign(index int, value int) {}

func g(n int) int { return n & (n + 1) } // off trailing ones
func h(n int) int { return n | (n + 1) } // flip the last unset bit

// temp

func sum(nums []int, left, right int) int {
	sum := 0
	for ; left <= right; left++ {
		sum += nums[left]
	}
	return sum
}

func testBIT() {
	const loopCount = 1000
	const numsLen int = 1e5
	const numRange int = 1e5
	const numDelta int = 100

	t := time.Now()
	nums := make([]int, numsLen)
	for i := range nums {
		nums[i] = rand.IntN(numRange*2+1) - numRange
	}
	bit := MakeBITFromNums(nums)
	for i, num := range nums {
		if bit.Item(i) != num {
			fmt.Printf("bit.Item(%v): %v != %v\n", i, bit.Item(i), num)
		}
	}
	// fmt.Println(nums)
	// fmt.Println(bit)

	for range loopCount {
		right := rand.IntN(numsLen)
		left := rand.IntN(right + 1)
		if bit.Sum(left, right) != sum(nums, left, right) {
			fmt.Printf("bit.Sum(%v, %v) is not correct\n", left, right)
		}
		index := rand.IntN(numsLen)
		delta := rand.IntN(numDelta*2+1) - numDelta
		bit.Update(index, delta)
		nums[index] += delta
		if bit.Item(index) != nums[index] {
			fmt.Printf("bit.Sum(%v, %v) is not correct after being updated\n", left, right)
		}
		left, right = max(index-1, 0), min(index+1, numsLen-1)
		if bit.Sum(left, right) != sum(nums, left, right) {
			fmt.Printf("bit.Sum(%v, %v) is not correct after being updated\n", left, right)
		}
	}
	fmt.Printf("testBIT() time: %v\n", time.Since(t))
}

func main() {
	testBIT()
}

// 1001011
// 1001100

// 1001011
// 1001100

// g:
// 1001011
// 1001000
// 1000000
// 0000000

// h:
// 1001011
// 1001111
// 1011111
// 1111111

// 5:
// g:
// 101 5
// 100 4
// 000 0

// h:
// 101 5
// 111 7

// h 4
// 100 4
// 101 5
// 111 7
