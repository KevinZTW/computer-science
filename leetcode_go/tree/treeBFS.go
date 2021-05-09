package main

import "fmt"

func main() {
	root := &treeNode{
		value: 12,
	}
	root.left = NewTreeNode(7)
	root.right = NewTreeNode(1)
	root.left.left = NewTreeNode(9)
	root.right.left = NewTreeNode(10)
	root.right.right = NewTreeNode(5)
	root.right.left.left = NewTreeNode(10)
	root.right.left.right = NewTreeNode(12)
	root.right.right.left = NewTreeNode(7)
	root.right.right.right = NewTreeNode(66)
	fmt.Println(traversal(root))
	fmt.Println(traversalZigZag(root))
}

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func NewTreeNode(n int) *treeNode {
	return &treeNode{
		value: n,
	}
}

func traversal(head *treeNode) [][]int {

	result := [][]int{}
	queue := []*treeNode{head}

	for len(queue) > 0 {
		l := len(queue)
		temp := make([]int, 0, l)
		for i := 0; i < l; i++ {

			if queue[i].left != nil {
				queue = append(queue, queue[i].left)
			}
			if queue[i].right != nil {
				queue = append(queue, queue[i].right)
			}
			temp = append(temp, queue[i].value)
		}
		queue = queue[l:]
		result = append(result, temp)

	}
	return result
}

func traversalZigZag(root *treeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*treeNode{root}
	result := [][]int{}
	leftToRight := true

	for len(queue) > 0 {
		l := len(queue)
		temp := []int{}
		for i := 0; i < l; i++ {

			if queue[i].left != nil {
				queue = append(queue, queue[i].left)
			}
			if queue[i].right != nil {
				queue = append(queue, queue[i].right)
			}
			if leftToRight {
				temp = append(temp, queue[i].value)
			} else {
				temp = append(temp, 0)
				copy(temp[1:], temp)
				temp[0] = queue[i].value
			}

		}
		if leftToRight {
			leftToRight = false
		} else {
			leftToRight = true
		}
		result = append(result, temp)
		queue = queue[l:]

	}

	return result

}
