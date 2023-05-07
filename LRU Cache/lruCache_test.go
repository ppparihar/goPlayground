package main

import "testing"

func TestLruCache(t *testing.T) {

	lru := Constructor(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	got := lru.Get(1)
	want := 0
	validate(got, want, t)

	lru.Put(3, 3)
	got = lru.Get(2)
	want = -1
	validate(got, want, t)

	lru.Put(4, 4)
	got = lru.Get(1)
	want = -1
	validate(got, want, t)
}

func TestLruCache2(t *testing.T) {

	lru := Constructor(2)
	lru.Put(1, 1)

	got := lru.Get(1)
	want := 1
	validate(got, want, t)

	lru.Put(1, 2)
	got = lru.Get(1)
	want = 2
	validate(got, want, t)
}

func validate(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
