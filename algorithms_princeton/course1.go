package main

import "fmt"


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