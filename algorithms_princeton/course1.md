# Course 1

## step to developing a useful algorithms

1. Model the problem ( think about what step need to take)
2. Find a algorithm to solve it
3. Fast enough? Fit memory?
4. If not, figure out why

## 1.5 Union Find

### Quick Find

So called eager algorithms.

#### Data structure :

- Integer array id [ ] of size N
- intepretation : p and q are connected iff (if and only if) they have the same id e.g. [1,3,1, 3, 3] means that 0=2 1=3=4

Find : - To check if they have same id

Union : - To connect p and q, change all entries whose id equals id[p] and id[q]

```go

func main(){
	fmt.Println("course 1")
	union := NewUnion(5)
	fmt.Println(union.IsConnected(2,4))
	union.Union(2,4)
	fmt.Println(union.IsConnected(2,4))
}


type Union []int


func NewUnion(n int) Union{
	union := make(Union, n)
	for i:=0;i<n;i++{
		union[i]=i
	}
	return union
}

func (c Union)IsConnected(p,q int) bool{
	return c[p]==c[q]
}

func (c Union)Union (p,q int) {
	for i, _ := range c{
		if c[i] == c[q]{
			c[i] = c[p]
		}
	}
}
```

But it's too slow

### Quick Union

so called lazy approach

#### Data structure :

- integer array id[ ] of size N
- intepretation : id[i] is parent of i ([0,1,2,1]) => 0 1=3 2

Find : find if p and q have same root
Union : set id of p's root to q's root

### Improvement1 - Weighted Quick Union

- Put the smaller tree under the larger tree

#### Data Structure

- Same as quick union, but maintain extra array sz[i] to count number of objects in the tree root of i

```go
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
```

#### Running time

##### Find :

take time propotional to depth of p and q
logN since the depth would be logN

##### Union :

take contant time, given roots
logN

##### Proposition :

Depth of any node x is at most lgN, why?
Increases by 1 when tree T1 containing x is merged into another tree T2.
・The size of the tree containing x at least doubles since | T 2 | ≥ | T 1 |
・Size of tree containing x can double at most lg N times.

### Improvement2 - Path Compression

when we compute the root of p, we are touching all the nodes on the path from that node to the root, while we are doing that we could make node just point to the root

```go
	func (c QuickUnion) root(i int) int {
	for c[i] != i {
		i = c[i]
		// add one line code below could make the tree almost flat
		c[i] = c[c[i]]
	}
	return i
}
```

#### Proposition

Hopcroft- Ulman, Tarjan prove that , any sequence of M union-find ops
on N objects makes ≤ c ( N + M lg\* N ) array accesses.

lg* N is the number of times we need to take the log of N to get one ( N 取對數到變成 1 的次數) it's called the iterated log function. In the real world we could think it would less than 5 ( lg*N = 5 when N = 2^65536)

### Union Find Application - Percolation

- It a modle of many physical system
- N-by-N grid
- each grid is open with posibility p
- Systme `percolates` iff top and bottom are connected by one sites

#### Monte Carlo simulation

- create an object for each site and name them 0 to N^2-1
- Site are in same component if connected by open sites
- Percolate iff any site on the top row is connected to the bottom row

- If we check whether every site on the top is connected to the site on the bottom, it would be too slow (N^2)

##### create virtual site on top and bottom

check once!
