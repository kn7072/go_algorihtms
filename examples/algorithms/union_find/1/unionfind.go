package main

// go get github.com/korzhnev/unionfind
// https://pkg.go.dev/github.com/korzhnev/unionfind#section-readme

// пример взят
// https://github.com/theodesp/unionfind/blob/master/unionfind.go

type UnionFind struct {
	root []int
	size []int
}

// New returns an initialized list of size
func NewUnionFind(size int) *UnionFind {
	return new(UnionFind).init(size)
}

// Constructor initializes root and size arrays
func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)
	uf.root = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

// Union connects p and q by finding their roots and comparing their respective
// size arrays to keep the tree flat
func (uf *UnionFind) Union(p, q int) {
	qRoot := uf.Root(q)
	pRoot := uf.Root(p)

	if uf.size[qRoot] < uf.size[pRoot] {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.root[pRoot] = uf.root[qRoot]
		uf.size[qRoot] += uf.size[pRoot]
	}
}

// Root or Find traverses each parent element while compressing the
// levels to find the root element of p
// If we attempt to access an element outside the array it returns -1
func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

// Root or Find
func (uf *UnionFind) Find(p int) int {
	return uf.Root(p)
}

// Check if items p,q are connected
func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Root(p) == uf.Root(q)
}
