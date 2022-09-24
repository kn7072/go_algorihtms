package main

import (
	"sync"
)

// https://github.com/theodesp/unionfind/blob/master/unionfind.go

// Thread safe version of Union-Find using plain locks.


type ThreadSafeUnionFind struct {
	sync.RWMutex
	uf *UnionFind
}

func NewThreadSafeUnionFind(size int) ThreadSafeUnionFind {
	safeUnionFind := ThreadSafeUnionFind{}
	safeUnionFind.uf = NewUnionFind(size)

	return safeUnionFind
}

func (suf *ThreadSafeUnionFind) Union(p int, q int) {
	suf.Lock()
	defer suf.Unlock()

	suf.uf.Union(p, q)
}


func (suf *ThreadSafeUnionFind) Root(p int) int {
	suf.Lock()
	defer suf.Unlock()

	return suf.uf.Root(p)
}

// Root or Find
func (suf *ThreadSafeUnionFind) Find(p int) int {
	return suf.uf.Root(p)
}

// Unfortunately all the calls are coerced to writes thats why we use a Writer lock
func (suf *ThreadSafeUnionFind) Connected(p int, q int) bool {
	suf.Lock()
	defer suf.Lock()

	return suf.uf.Root(p) == suf.uf.Root(p)
}