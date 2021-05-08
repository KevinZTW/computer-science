package main

import "fmt"

func main() {
	linkList := NewLinkList(1)
	linkList.push(2)
	linkList.push(3)
	linkList.push(4)
	linkList.push(5)
	linkList.push(6)
	linkList.print()
	linkList.head = reverseBetweenMe(linkList.head, 2, 4)

	linkList.print()
}

type LinkList struct {
	head *ListNode
	tail *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkList(val int) *LinkList {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}

	return &LinkList{
		head: node,
		tail: node,
	}
}

func (c *LinkList) push(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}

	c.tail.Next = node
	c.tail = node
}

func (c *LinkList) print() {
	node := c.head
	for node.Next != nil {
		fmt.Print(node.Val, " ->")
		node = node.Next
	}
	fmt.Print(node.Val)
}
func reverseBetweenMe(head *ListNode, m int, n int) *ListNode {
	dummy := ListNode{Val: 0}
	dummy.Next = head

	prvnode := &dummy

	for prvnode.Next.Val != m {
		prvnode = prvnode.Next
	}

	pre := prvnode.Next
	fmt.Println(pre.Val)
	curr, tail := pre.Next, pre
	for pre.Val != n {
		nxt := curr.Next
		curr.Next = pre
		pre = curr
		curr = nxt
	}

	tail.Next = curr

	prvnode.Next = pre

	return dummy.Next
}
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	dummy := ListNode{Val: 0}
	dummy.Next = head

	node := &dummy

	for i := 1; i < m; i++ {
		node = node.Next
	}

	pre := node
	curr := node.Next
	tail := curr
	for i := m; i <= n; i++ {
		nxt := curr.Next
		curr.Next = pre
		pre = curr
		curr = nxt
	}

	tail.Next = curr

	node.Next = pre

	return dummy.Next
}

func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}
