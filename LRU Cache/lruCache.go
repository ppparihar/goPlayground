package main

import "container/list"

type LRUCache struct {
	list     *list.List
	capacity int
	cache    map[int]*list.Element
}
type LRUNode struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.capacity = capacity
	lru.list = list.New()
	lru.cache = make(map[int]*list.Element)
	return lru
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.list.MoveToFront(v)
		return v.Value.(LRUNode).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		this.list.MoveToFront(v)
		node := (v.Value.(LRUNode))
		node.val = value
		v.Value = node
	} else {
		if this.list.Len() == this.capacity {
			element := this.list.Back()

			delete(this.cache, element.Value.(LRUNode).key)
			this.list.Remove(element)
		}
		listNode := LRUNode{key, value}
		this.cache[key] = this.list.PushFront(listNode)
	}
}
