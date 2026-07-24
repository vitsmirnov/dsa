package main

// type SegTreeItem struct{ min, max int }

type SegmentTree[T any] struct {
	nodes     []T
	initNode  func(val int) T
	buildNode func(leftChild, rightChild T) T
	size      int
}

func MakeSegmentTree[T any](arr []int,
	initNode func(val int) T,
	buildNode func(leftChild, rightChild T) T) *SegmentTree[T] {

	size := len(arr)
	st := &SegmentTree[T]{
		nodes:     make([]T, size*4),
		initNode:  initNode,
		buildNode: buildNode,
		size:      size}
	st.build(arr, 0, 0, size-1)
	return st
}

func (st *SegmentTree[T]) build(arr []int, pos int, left, right int) {

	if left == right {
		// st.items[pos].min = arr[segLeft]
		// st.items[pos].max = arr[segLeft]
		st.nodes[pos] = st.initNode(arr[left])
		return
	}

	mid := left + (right-left)/2
	leftChild, rightChild := pos*2+1, pos*2+2
	st.build(arr, leftChild, left, mid)
	st.build(arr, rightChild, mid+1, right)
	// st.items[pos].min = min(st.items[leftChild].min, st.items[rightChild].min)
	// st.items[pos].max = max(st.items[leftChild].max, st.items[rightChild].max)
	st.nodes[pos] = st.buildNode(st.nodes[leftChild], st.nodes[rightChild])
}

func (st *SegmentTree[T]) Query(left, right int) T {
	return st.item(0, 0, st.size-1, left, right)
}

func (st *SegmentTree[T]) item(pos int, segLeft, segRight, qLeft, qRight int) T {
	if segLeft == qLeft && segRight == qRight {
		return st.nodes[pos]
	}

	mid := segLeft + (segRight-segLeft)/2
	if qRight <= mid {
		return st.item(pos*2+1, segLeft, mid, qLeft, qRight)
	} else if qLeft > mid {
		return st.item(pos*2+2, mid+1, segRight, qLeft, qRight)
	} else {
		resLeft := st.item(pos*2+1, segLeft, mid, qLeft, mid)
		resRight := st.item(pos*2+2, mid+1, segRight, mid+1, qRight)
		// return SegTreeItem{
		// 	min: min(mmLeft.min, mmRight.min),
		// 	max: max(mmLeft.max, mmRight.max)}
		return st.buildNode(resLeft, resRight)
	}
}

func (st *SegmentTree[T]) Update(index int, value int) {
	st.update(0, 0, st.size-1, index, value)
}

func (st *SegmentTree[T]) update(pos int, left, right int, index int, value int) {
	if left == right {
		// st.items[pos].min = value
		// st.items[pos].max = value
		st.nodes[pos] = st.initNode(value)
		return
	}

	mid := left + (right-left)/2
	leftChild, rightChild := pos*2+1, pos*2+2
	if index <= mid {
		st.update(leftChild, left, mid, index, value)
	} else {
		st.update(rightChild, mid+1, right, index, value)
	}
	// st.items[pos].min = min(st.items[leftChild].min, st.items[rightChild].min)
	// st.items[pos].max = max(st.items[leftChild].max, st.items[rightChild].max)
	st.nodes[pos] = st.buildNode(st.nodes[leftChild], st.nodes[rightChild])
}

// func (st *SegmentTree[T]) Nums() []int {
// 	res := make([]int, st.size)
// 	for i := range res {
// 		res[i] = st.Query(i, i)
// 	}
// 	return res
// }
