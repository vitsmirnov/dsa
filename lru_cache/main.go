package main

import "container/list"

type LRUCache struct {
	nodes    map[int]*list.Element
	queue    *list.List
	capacity int
}

type KeyValue struct {
	key int
	val int
}

func MakeLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		nodes:    make(map[int]*list.Element, capacity),
		queue:    list.New(),
		capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	node, has := this.nodes[key]
	if !has {
		return -1
	}
	this.queue.MoveToBack(node)
	return node.Value.(KeyValue).val
}

func (this *LRUCache) Put(key int, value int) {
	node, has := this.nodes[key]
	if has {
		node.Value = KeyValue{key: key, val: value}
		this.queue.MoveToBack(node)
	} else {
		if this.queue.Len() >= this.capacity {
			kv := this.queue.Remove(this.queue.Front())
			delete(this.nodes, kv.(KeyValue).key)
		}
		this.nodes[key] = this.queue.PushBack(KeyValue{key: key, val: value})
	}
}
