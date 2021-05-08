package main

func main() {
	root := &treeNode{
		value: 12,
	}
	root.left = NewTreeNode(7)
	root.right = NewTreeNode(1)
	root.left.left = NewTreeNode(9)
	root.right.left = NewTreeNode(10)
	root.right.right = NewTreeNode(5)

}

type binaryTree struct {
	root *treeNode
}

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func NewBinaryTree() binaryTree {
	return binaryTree{
		root: &treeNode{},
	}
}

func NewTreeNode(n int) *treeNode {
	return &treeNode{
		value: n,
	}
}

func (c binaryTree) Insert(n int) {
	if c.root == nil {
		c.root = &treeNode{
			value: n,
		}
		return
	}
}

func Traverse(head *treeNode) []int {
	var result []int
	var queue []*treeNode
	queue = append(queue, head)

	for len(queue) > 0 {
		currentLevel := len(result) + 1

	}

}
