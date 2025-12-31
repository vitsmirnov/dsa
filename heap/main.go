package main

import (
	"fmt"
	"math/rand"
	"slices"
)

// max heap
// type IndexedHeap struct {
// 	values  []int
// 	indexes map[int]int
// }

// func (h *IndexedHeap) Len() int {
// 	return len(h.values)
// }

// func (h *IndexedHeap) Top() int {
// 	return h.values[0]
// }

// func (h *IndexedHeap) Pop() int {
// 	res := h.Top()

// 	len := len(h.values) - 1
// 	h.values[0] = h.values[len]
// 	h.values = h.values[:len]

// 	maxChild := func(i int) int {
// 		if l, r := i*2+1, i*2+2; l >= len { // have no children
// 			return -1
// 		} else if r >= len || h.values[l] > h.values[r] { // have only left child ot left > right
// 			return l
// 		} else {
// 			return r
// 		}
// 	}

// 	for i, mc := 0, maxChild(0); mc != -1 && h.values[i] < h.values[mc]; {
// 		h.values[i], h.values[mc] = h.values[mc], h.values[i]
// 		i = mc
// 		mc = maxChild(i)
// 	}

// 	return res
// }

// func (h *IndexedHeap) push(val int) {
// 	h.values = append(h.values, val)
// 	i := h.Len() - 1
// 	p := (i - 1) / 2
// 	for i != 0 && h.values[p] < h.values[i] {
// 		h.values[p], h.values[i] = h.values[i], h.values[p]
// 		i = p
// 		p = (i - 1) / 2
// 	}
// }

// func (h *IndexedHeap) Push(values ...int) {
// 	for _, value := range values {
// 		h.push(value)
// 	}
// }

// func (h *IndexedHeap) Has(value int) bool {
// 	_, has := h.indexes[value]
// 	return has
// }

// func (h IndexedHeap) String() string {
// 	return fmt.Sprint(h.values)
// }

// max heap
type Heap struct {
	values []int
}

func (h *Heap) Len() int {
	return len(h.values)
}

func (h *Heap) Top() int {
	return h.values[0]
}

func (h *Heap) Pop() int {
	res := h.Top()

	len := len(h.values) - 1
	h.values[0] = h.values[len]
	h.values = h.values[:len]

	maxChild := func(i int) int {
		if l, r := i*2+1, i*2+2; l >= len { // have no children
			return -1
		} else if r >= len || h.values[l] > h.values[r] { // have only left child or left > right
			return l
		} else {
			return r
		}
	}

	for i, mc := 0, maxChild(0); mc != -1 && h.values[i] < h.values[mc]; {
		h.values[i], h.values[mc] = h.values[mc], h.values[i]
		i = mc
		mc = maxChild(i)
	}

	return res
}

// left child of i: i*2+1
// right child of i: i*2+2
// parent of i: (i-1) / 2

func (h *Heap) push(val int) {
	h.values = append(h.values, val)
	i := h.Len() - 1
	p := (i - 1) / 2
	for i != 0 && h.values[p] < h.values[i] {
		h.values[p], h.values[i] = h.values[i], h.values[p]
		i = p
		p = (i - 1) / 2
	}
}

func (h *Heap) Push(values ...int) {
	for _, value := range values {
		h.push(value)
	}
}

func (h Heap) String() string {
	return fmt.Sprint(h.values)
}

func testHeap() {
	arr := make([]int, 1_000_000)
	for i := range arr {
		arr[i] = rand.Intn(2_000_000) - 1_000_000
	}
	// for _, val := range arr {
	// 	fmt.Println(val)
	// }

	heap := Heap{}
	heap.Push(arr...)
	for i := len(arr) - 1; heap.Len() != 0; i-- {
		arr[i] = heap.Pop()
	}

	if slices.IsSorted(arr) {
		fmt.Println("Heap sort OK")
	} else {
		fmt.Println("Heap sort failed!")
	}
}

func main() {
	testHeap()
}

// 0
// 1        2
// 3   4    5     6
// 7 8 9 10 11 12 13 14
