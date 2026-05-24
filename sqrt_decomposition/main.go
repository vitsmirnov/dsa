package main

// todo: tests

type SqrtDec struct {
	nums      []int
	adds      []int
	numFreqs  []map[int]int
	blockSize int
}

func MakeSqrtDec(nums []int) *SqrtDec {
	// time: O(n), space: O(n)

	numsLen := len(nums)
	blockSize := sqrt(numsLen)
	blockCount := (numsLen + blockSize - 1) / blockSize // ceil
	numsCopy := make([]int, numsLen)
	adds := make([]int, blockCount)
	numFreqs := make([]map[int]int, blockCount)
	for i := range numFreqs {
		numFreqs[i] = map[int]int{}
	}
	for i, num := range nums {
		numsCopy[i] = num
		numFreqs[i/blockSize][num]++
	}
	return &SqrtDec{
		nums:      numsCopy,
		adds:      adds,
		numFreqs:  numFreqs,
		blockSize: blockSize}
}

// add val to range [left..right]
func (sd *SqrtDec) Add(left, right int, val int) {
	// time: O(sqrt(n)), space: O(1)
	// n - query range (right-left+1)

	blockSize := sd.blockSize
	rightIsLast := right == len(sd.nums)-1
	for i := left; i <= right; {
		if i%blockSize == 0 && (i+blockSize-1 <= right || rightIsLast) { // ~
			sd.adds[i/blockSize] += val
			i += blockSize
		} else {
			j := i / blockSize
			curVal := sd.nums[i]
			sd.numFreqs[j][curVal]--
			newVal := curVal + val
			sd.numFreqs[j][newVal]++
			sd.nums[i] = newVal
			i++
		}
	}
}

func (sd *SqrtDec) CountNums(num int) int {
	// time: O(sqrt(n)), sapce:(1)
	// n - len(nums)

	count := 0
	for i, numFreqs := range sd.numFreqs {
		add := sd.adds[i]
		count += numFreqs[num-add]
	}
	return count
}

func sqrt(n int) int {
	// time: O(log n), space: O(1)

	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {

}

// 3
// 1 2 3 4 5 6 7 8
