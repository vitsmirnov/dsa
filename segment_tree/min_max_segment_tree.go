package main

import "math"

type MinMax struct{ min, max int }

type MinMaxSegmentTree struct{ items []MinMax }

func MakeMinMaxSegmentTree(arr []int) *MinMaxSegmentTree {
	st := &MinMaxSegmentTree{items: make([]MinMax, len(arr)*4)}
	st.build(arr, 0, 0, len(arr)-1)
	return st
}

func (st *MinMaxSegmentTree) build(arr []int, pos int, segLeft, segRight int) {
	if segLeft == segRight {
		st.items[pos].min = arr[segLeft]
		st.items[pos].max = arr[segLeft]
	} else {
		mid := segLeft + (segRight-segLeft)/2
		leftChild, rightChild := pos*2+1, pos*2+2
		st.build(arr, leftChild, segLeft, mid)
		st.build(arr, rightChild, mid+1, segRight)
		st.items[pos].min = min(st.items[leftChild].min, st.items[rightChild].min)
		st.items[pos].max = max(st.items[leftChild].max, st.items[rightChild].max)
	}
}

func (st *MinMaxSegmentTree) Min(left, right int) int {
	return st.MinMax(left, right).min
}

func (st *MinMaxSegmentTree) Max(left, right int) int {
	return st.MinMax(left, right).max
}

func (st *MinMaxSegmentTree) MinMax(left, right int) MinMax {
	return st.item(0, 0, len(st.items)/4-1, left, right)
}

func (st *MinMaxSegmentTree) item(pos int, segLeft, segRight, qLeft, qRight int) MinMax {
	if qLeft > qRight {
		return MinMax{min: math.MaxInt, max: math.MinInt}
	}
	if segLeft == qLeft && segRight == qRight {
		return st.items[pos]
	}
	mid := segLeft + (segRight-segLeft)/2
	mmLeft := st.item(pos*2+1, segLeft, mid, qLeft, min(qRight, mid))
	mmRight := st.item(pos*2+2, mid+1, segRight, max(qLeft, mid+1), qRight)
	return MinMax{
		min: min(mmLeft.min, mmRight.min),
		max: max(mmLeft.max, mmRight.max)}
}

func (st *MinMaxSegmentTree) item2(pos int, segLeft, segRight, qLeft, qRight int) MinMax {
	if segLeft == qLeft && segRight == qRight {
		return st.items[pos]
	}
	mid := segLeft + (segRight-segLeft)/2
	if qRight <= mid {
		return st.item(pos*2+1, segLeft, mid, qLeft, qRight)
	} else if qLeft > mid {
		return st.item(pos*2+2, mid+1, segRight, qLeft, qRight)
	} else {
		mmLeft := st.item(pos*2+1, segLeft, mid, qLeft, mid)
		mmRight := st.item(pos*2+2, mid+1, segRight, mid+1, qRight)
		return MinMax{
			min: min(mmLeft.min, mmRight.min),
			max: max(mmLeft.max, mmRight.max)}
	}
}

func (st *MinMaxSegmentTree) Update(index int, value int) {
	st.update(0, 0, len(st.items)/4-1, index, value)
}

func (st *MinMaxSegmentTree) update(pos int, segLeft, segRight int, index int, value int) {
	if segLeft == segRight {
		st.items[pos].min = value
		st.items[pos].max = value
	} else {
		mid := segLeft + (segRight-segLeft)/2
		leftChild, rightChild := pos*2+1, pos*2+2
		if index <= mid {
			st.update(leftChild, segLeft, mid, index, value)
		} else {
			st.update(rightChild, mid+1, segRight, index, value)
		}
		st.items[pos].min = min(st.items[leftChild].min, st.items[rightChild].min)
		st.items[pos].max = max(st.items[leftChild].max, st.items[rightChild].max)
	}
}

func (st *MinMaxSegmentTree) Nums() []int {
	res := make([]int, len(st.items)/4)
	for i := range res {
		res[i] = st.Min(i, i)
	}
	return res
}
