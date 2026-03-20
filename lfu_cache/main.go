package main

import "container/list"

type LFUCache struct {
	freqs    map[int]int                // key: frequency
	items    map[int]*OrderedMap[*Item] // frequensy: items
	minFreq  int
	capacity int
	size     int
}

type Item struct {
	key   int
	value int
	freq  int
}

func MakeLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		freqs:    map[int]int{},
		items:    map[int]*OrderedMap[*Item]{},
		minFreq:  0,
		capacity: capacity,
		size:     0}
}

func (this *LFUCache) Get(key int) int {
	// time: O(1)

	freq, has := this.freqs[key]
	if !has {
		return -1
	}
	item := this.items[freq].Get(key)
	this.incFreq(item)
	return item.value
}

func (this *LFUCache) Put(key int, value int) {
	// time: O(1)

	freq, has := this.freqs[key]
	if has {
		item := this.items[freq].Get(key)
		item.value = value
		this.incFreq(item)
	} else {
		if this.size == this.capacity {
			item := this.items[this.minFreq].Front()
			this.items[this.minFreq].Remove(item.key)
			delete(this.freqs, item.key)
			this.size--
		}
		this.freqs[key] = 1
		if _, has := this.items[1]; !has {
			this.items[1] = MakeOrderedMap[*Item]()
		}
		this.items[1].Put(key, &Item{key: key, value: value, freq: 1})
		this.minFreq = 1
		this.size++
	}
}

func (this *LFUCache) incFreq(item *Item) {
	// time: O(1)

	this.items[item.freq].Remove(item.key)
	if this.items[item.freq].Len() == 0 && this.minFreq == item.freq {
		this.minFreq = item.freq + 1
	}
	item.freq++
	this.freqs[item.key]++
	if _, has := this.items[item.freq]; !has {
		this.items[item.freq] = MakeOrderedMap[*Item]()
	}
	this.items[item.freq].Put(item.key, item)
}

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
	// time: O(1)

	item, has := this.items[key]
	if has {
		item.Value = value
	} else {
		this.items[key] = this.list.PushBack(value)
	}
}
func (this *OrderedMap[T]) Get(key int) T {
	// time: O(1)

	return this.items[key].Value.(T)
}
func (this *OrderedMap[T]) Exist(key int) bool {
	// time: O(1)

	_, has := this.items[key]
	return has
}
func (this *OrderedMap[T]) Remove(key int) {
	// time: O(1)

	item, has := this.items[key]
	if !has {
		return
	}
	this.list.Remove(item)
	this.items[key] = nil
	delete(this.items, key)
}
func (this *OrderedMap[T]) Front() T {
	// time: O(1)

	return this.list.Front().Value.(T)
}
