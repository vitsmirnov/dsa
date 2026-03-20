package main

import "container/list"

type OrderedMap[T any] struct {
	list  *list.List
	items map[int]*list.Element
}

func MakeOrderedMap[T any]() *OrderedMap[T] {
	return &OrderedMap[T]{
		list:  list.New(),
		items: map[int]*list.Element{}}
}
func (this *OrderedMap[T]) Len() int { return this.list.Len() }
func (this *OrderedMap[T]) Put(key int, value T) {
	item, has := this.items[key]
	if has {
		item.Value = value
	} else {
		this.items[key] = this.list.PushBack(value)
	}
}
func (this *OrderedMap[T]) Get(key int) T {
	return this.items[key].Value.(T)
}
func (this *OrderedMap[T]) Exist(key int) bool {
	_, has := this.items[key]
	return has
}
func (this *OrderedMap[T]) Remove(key int) {
	item, has := this.items[key]
	if !has {
		return
	}
	this.list.Remove(item)
	this.items[key] = nil
	delete(this.items, key)
}
func (this *OrderedMap[T]) Front() T {
	return this.list.Front().Value.(T)
}
func (this *OrderedMap[T]) Back() T {
	return this.list.Back().Value.(T)
}

// todo: iterator
