package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = NewTreeNode(4)
	root.Right = NewTreeNode(8)
	root.Left.Left = NewTreeNode(11)
	root.Left.Left.Left = NewTreeNode(7)
	root.Left.Left.Right = NewTreeNode(2)
	root.Right.Left = NewTreeNode(13)
	root.Right.Right = NewTreeNode(4)
	root.Right.Right.Left = NewTreeNode(5)
	root.Right.Right.Right = NewTreeNode(1)
	fmt.Println(HasAllPathSum(root, 22))
	fmt.Println(pathSum(root, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(n int) *TreeNode {
	return &TreeNode{
		Val: n,
	}
}

func HasAllPathSum(root *TreeNode, sum int) [][]int {
	temp := []int{}
	result := [][]int{}
	hasAllPathSum(root, sum, temp, &result)
	return result
}

func hasAllPathSum(root *TreeNode, sum int, temp []int, result *[][]int) {

	if root == nil {
		return
	}
	temp = append(temp, root.Val)

	if root.Val == sum && root.Left == nil && root.Right == nil {

		fmt.Printf("temp is %v and current val is %v\n", temp, root.Val)
		//this would cause false
		*result = append(*result, temp[:])

		//need to make copy
		tempCopy := make([]int, len(temp))
		copy(tempCopy, temp)
		*result = append(*result, tempCopy)

	}

	hasAllPathSum(root.Left, sum-root.Val, temp, result)
	hasAllPathSum(root.Right, sum-root.Val, temp, result)
	return
}

func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	slice = findPath(root, sum, slice, []int(nil))
	return slice
}

func findPath(n *TreeNode, sum int, slice [][]int, stack []int) [][]int {
	if n == nil {
		return slice
	}
	sum -= n.Val
	stack = append(stack, n.Val)
	if sum == 0 && n.Left == nil && n.Right == nil {
		slice = append(slice, append([]int{}, stack...))
		stack = stack[:len(stack)-1]
	}
	slice = findPath(n.Left, sum, slice, stack)
	slice = findPath(n.Right, sum, slice, stack)
	return slice
}
