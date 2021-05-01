package main

import "fmt"

func main() {
	fmt.Println("week 2")

	// Stack implemented in link list
	linkListStack := NewLinkListStack()
	linkListStack.Push("2")
	linkListStack.Push("5")
	linkListStack.Pop()
	fmt.Println(linkListStack.PrintStack())

	//Stack implemented in slice

	a := ArrayStack{}
	arrayStack := &a
	arrayStack.Push("2")
	arrayStack.Push("5")
	arrayStack.Pop()
	fmt.Println(arrayStack.Print())

}

type LinkListStack struct {
	head *LinkListNode
}

type LinkListNode struct {
	next  *LinkListNode
	value string
}

func NewLinkListStack() *LinkListStack {
	return &LinkListStack{}
}

func (c *LinkListStack) Push(n string) {
	node := &LinkListNode{
		value: n,
	}
	if c.head == nil {
		c.head = node
	} else {
		node.next, c.head = c.head, node
	}
}

func (c *LinkListStack) Pop() string {
	r := c.head.value
	c.head = c.head.next
	return r
}

func (c *LinkListStack) PrintStack() []string {
	n := c.head
	result := make([]string, 0)

	if n != nil {
		result = append(result, n.value)
		for n.next != nil {
			n = n.next
			result = append(result, n.value)
		}
	}

	return result
}

type ArrayStack []string

func (c *ArrayStack) Push(n string) {
	*c = append(*c, n)

}

func (c *ArrayStack) Pop() {
	*c = (*c)[:len((*c))-1]
}

func (c ArrayStack) Print() []string {
	return c
}
