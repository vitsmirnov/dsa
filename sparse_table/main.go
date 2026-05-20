package main

type SparseTable struct {
	mins [][]int
	logs []int
}

func MakeSparseTable(nums []int) *SparseTable {
	numsLen := len(nums)
	// p := 0
	// for (1 << p) <= numsLen {
	// 	p++
	// }

	logs := make([]int, numsLen+1)
	logs[1] = 0
	for i := 2; i < len(logs); i++ {
		logs[i] = logs[i>>1] + 1
	}

	p := logs[numsLen] + 1
	mins := make([][]int, p)
	mins[0] = make([]int, numsLen)
	copy(mins[0], nums)
	for i := 1; i < p; i++ {
		mins[i] = make([]int, numsLen)
		for j := numsLen - (1 << i); j >= 0; j-- {
			mins[i][j] = min(mins[i-1][j], mins[i-1][j+(1<<(i-1))])
		}
	}

	return &SparseTable{
		mins: mins,
		logs: logs}
}
func (st *SparseTable) Min(left, right int) int {
	l := st.logs[right-left+1]
	return min(st.mins[l][left], st.mins[l][right-left+1])
}

// sum would be like that
// func (st *SparseTable) Sum(left, right int) int {
// 	return st.mins[][left]
// }

func testSparseTable() {

}

func getMin(nums []int) int {
	minNum := nums[0]
	for _, num := range nums {
		minNum = min(minNum, num)
	}
	return minNum
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

	testSparseTable()
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

// int lg[MAXN+1];
// lg[1] = 0;
// for (int i = 2; i <= MAXN; i++)
//     lg[i] = lg[i/2] + 1;

// 111

// 1: 1
// 2: 2
// 3: 2 + 1
// 4: 4
// 5: 4 + 1
// 6: 4 + 2
// 7: 4 + 4
// 8: 8
// 9: 8 + 1
// 10: 8 + 2
// 11: 8 + 4
// 12: 8 + 4
// 13: 8 + 8
// 14: 8 + 8
// 15: 8 + 8
// 16: 16
// 17: 16 + 1
// 18: 16 + 2
// 19: 16 + 4
// 20: 16 + 4
