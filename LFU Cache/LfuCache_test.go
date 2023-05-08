package main

import "testing"

//[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]

func TestLfuCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	Validate(cache.Get(1), 1, t)

	cache.Put(3, 3)
	Validate(cache.Get(2), -1, t)
	Validate(cache.Get(3), 3, t)
	cache.Put(4, 4)
	Validate(cache.Get(1), -1, t)
	Validate(cache.Get(3), 3, t)
	Validate(cache.Get(4), 4, t)
}
func TestLfuCache2(t *testing.T) {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	Validate(cache.Get(4), 4, t)
}
func Validate(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("got %v, wanted %v", actual, expected)
	}
}
