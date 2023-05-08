package main

import "container/list"

type LFUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type LFUNode struct {
	key   int
	value int
	freq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		node.Value.(*LFUNode).freq++
		this.tryMoveLeft(node)

		return node.Value.(*LFUNode).value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if node, ok := this.cache[key]; ok {
		node.Value.(*LFUNode).value = value
		node.Value.(*LFUNode).freq++
		this.tryMoveLeft(node)
		return
	}

	if this.capacity == this.list.Len() {
		delete(this.cache, this.list.Back().Value.(*LFUNode).key)
		this.list.Remove(this.list.Back())
	}

	node := &LFUNode{
		key:   key,
		value: value,
		freq:  1,
	}
	element := this.list.PushBack(node)
	this.cache[key] = element
	this.tryMoveLeft(element)
}

func (this *LFUCache) tryMoveLeft(node *list.Element) {
	cur := node.Value.(*LFUNode)
	for e := node.Prev(); e != nil; {
		prev := e.Value.(*LFUNode)
		if prev.freq > cur.freq {
			break
		}

		this.list.MoveBefore(node, e)
		e = node.Prev()
	}
}
