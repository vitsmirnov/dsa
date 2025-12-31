package main

import (
	"cmp"
	"container/heap"
)

// type LessFunc[T any] func(a, b T) bool

type PriorityQueue[T any] struct {
	items []T
	less  func(a, b T) bool
}

func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{items: []T{}, less: less}
}
func (pq *PriorityQueue[T]) Len() int           { return len(pq.items) }
func (pq *PriorityQueue[T]) Less(i, j int) bool { return pq.less(pq.items[i], pq.items[j]) }
func (pq *PriorityQueue[T]) Swap(i, j int)      { pq.items[i], pq.items[j] = pq.items[j], pq.items[i] }
func (pq *PriorityQueue[T]) Push(x any)         { pq.items = append(pq.items, x.(T)) }
func (pq *PriorityQueue[T]) Pop() any {
	l := len(pq.items) - 1
	res := pq.items[l]
	pq.items = pq.items[:l]
	return res
}
func (pq *PriorityQueue[T]) PushVal(val T) { heap.Push(pq, val) }
func (pq *PriorityQueue[T]) PopVal() T     { return heap.Pop(pq).(T) }
func (pq *PriorityQueue[T]) Top() T        { return pq.items[0] }

func Less[T cmp.Ordered](a, b T) bool    { return a < b }
func Greater[T cmp.Ordered](a, b T) bool { return b < a }

func NewMinPQ[T cmp.Ordered](dummy T) *PriorityQueue[T] { return NewPriorityQueue[T](Less) }
func NewMaxPQ[T cmp.Ordered](dummy T) *PriorityQueue[T] { return NewPriorityQueue[T](Greater) }

// temp (methods with checks)

// func (pq *PriorityQueue[T]) Less(i, j int) bool {
// 	if pq.less == nil {
// 		panic("Trying to call 'nil' less function from PriorityQueue!")
// 	}
// 	return pq.less(pq.items[i], pq.items[j])
// }
// func (pq *PriorityQueue[T]) Pop() any {
// 	if pq.Len() == 0 {
// 		panic("Trying to call Pop() from an empty PriorityQueue!")
// 	}

// 	l := len(pq.items) - 1
// 	res := pq.items[l]
// 	pq.items = pq.items[:l]
// 	return res
// }
// func (pq *PriorityQueue[T]) Top() T {
// 	if pq.Len() == 0 {
// 		panic("Trying to call Top() from an empty PriorityQueue!")
// 	}
// 	return pq.items[0]
// }
