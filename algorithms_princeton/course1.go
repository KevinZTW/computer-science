package main

import "fmt"

func main() {
	fmt.Println("course 1")
	// QuickFind
	quickFind := NewQuickFind(5)
	fmt.Println(quickFind.IsConnected(2, 4))
	quickFind.Union(2, 4)
	fmt.Println(quickFind.IsConnected(2, 4))

	// QuickUnion
	quickUnion := NewQuickUnion(5)
	fmt.Println(quickUnion.IsConnected(2, 4))
	quickUnion.Union(2, 4)
	fmt.Println(quickUnion.IsConnected(2, 4))

	// QuickUnion improvement - weighted
	wightedQuickUnion := NewWeightedQuickUnioin(5)
	wightedQuickUnion.Union(2, 4)
	wightedQuickUnion.Union(1, 3)
	wightedQuickUnion.Union(1, 2)
	fmt.Println(wightedQuickUnion)
}

type QuickFind []int

func NewQuickFind(n int) QuickFind {
	union := make(QuickFind, n)
	for i := 0; i < n; i++ {
		union[i] = i
	}
	return union
}

func (c QuickFind) IsConnected(p, q int) bool {
	return c[p] == c[q]
}

func (c QuickFind) Union(p, q int) {
	for i, _ := range c {
		if c[i] == c[q] {
			c[i] = c[p]
		}
	}
}

type QuickUnion []int

func NewQuickUnion(n int) QuickUnion {
	quickUnion := make(QuickUnion, n)
	for i := 0; i < n; i++ {
		quickUnion[i] = i
	}
	return quickUnion
}

func (c QuickUnion) root(i int) int {
	for c[i] != i {
		i = c[i]
		c[i] = c[c[i]]
	}
	return i
}

func (c QuickUnion) IsConnected(p, q int) bool {

	return c.root(p) == c.root(q)
}

func (c QuickUnion) Union(p, q int) {

	pRoot := c.root(p)
	c[pRoot] = c.root(q)
}

type WeightedQuickUnion struct {
	QuickUnion
	sz []int
}

func NewWeightedQuickUnioin(n int) WeightedQuickUnion {
	return WeightedQuickUnion{
		QuickUnion: NewQuickUnion(n),
		sz:         make([]int, n),
	}
}

func (c WeightedQuickUnion) Union(p, q int) {
	i, j := c.root(p), c.root(q)
	//link smaller tree to bigger one
	if c.sz[i] > c.sz[j] {
		c.QuickUnion[j] = i
		c.sz[i]++
	} else {
		c.QuickUnion[i] = j
		//update the sz array
		c.sz[j]++
	}

}
