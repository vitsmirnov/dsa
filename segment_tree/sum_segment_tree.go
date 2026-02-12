package main

type SumSegmentTree struct{ sums []int }

func MakeSumSegmentTree(arr []int) *SumSegmentTree {
	st := &SumSegmentTree{sums: make([]int, len(arr)*4)}
	st.build(arr, 0, 0, len(arr)-1)
	return st
}

func (st *SumSegmentTree) build(arr []int, pos int, segLeft, segRight int) {
	if segLeft == segRight {
		st.sums[pos] = arr[segLeft]
		return
	}
	mid := segLeft + (segRight-segLeft)/2
	leftChild, rightChild := pos*2+1, pos*2+2
	st.build(arr, leftChild, segLeft, mid)
	st.build(arr, rightChild, mid+1, segRight)
	st.sums[pos] = st.sums[leftChild] + st.sums[rightChild]
}

func (st *SumSegmentTree) Sum(left, right int) int {
	return st.sum(0, 0, len(st.sums)/4-1, left, right)
}

func (st *SumSegmentTree) sum(pos int, segLeft, segRight, qLeft, qRight int) int {
	if qLeft > qRight {
		return 0
	}
	if segLeft == qLeft && segRight == qRight {
		return st.sums[pos]
	}
	mid := segLeft + (segRight-segLeft)/2
	return st.sum(pos*2+1, segLeft, mid, qLeft, min(qRight, mid)) +
		st.sum(pos*2+2, mid+1, segRight, max(qLeft, mid+1), qRight)
}

func (st *SumSegmentTree) Update(index int, value int) {
	st.update(0, 0, len(st.sums)/4-1, index, value)
}

func (st *SumSegmentTree) update(pos int, segLeft, segRight int, index int, value int) {
	if segLeft == segRight {
		st.sums[pos] = value
	} else {
		mid := segLeft + (segRight-segLeft)/2
		if index <= mid {
			st.update(pos*2+1, segLeft, mid, index, value)
		} else {
			st.update(pos*2+2, mid+1, segRight, index, value)
		}
		st.sums[pos] = st.sums[pos*2+1] + st.sums[pos*2+2]
	}
}

func (st *SumSegmentTree) NumsLen() int { return len(st.sums) / 4 }

func (st *SumSegmentTree) Nums() []int {
	res := make([]int, len(st.sums)/4)
	for i := range res {
		res[i] = st.Sum(i, i)
	}
	return res
}
