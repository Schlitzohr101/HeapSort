package main

import (
	"fmt"
	"math"
)

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	data   int64
}

type Tree struct {
	root  *Node
	count int64
}

func (t *Tree) buildHeap(ar []int64) {
	for index := 0; index < len(ar); index++ {
		t.insert(ar[index])
	}
}

func (T *Tree) insert(data int64) *Tree {
	if T.root == nil {
		T.root = &Node{data: data, left: nil, right: nil, parent: nil}
	} else {
		T.root.insert(data)
	}
	T.count++
	return T
}

func (n *Node) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil, parent: n}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil, parent: n}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *Tree) print() {
	ar := make([]int64, int(math.Pow(float64(t.count), 2.0)))
	t.root.createAr(ar, 0)
	fmt.Println(ar)
}

func (n *Node) createAr(ar []int64, place int) {
	if n == nil {
		return
	}
	ar[place] = n.data
	n.left.createAr(ar, (place*2)+1)
	n.right.createAr(ar, (place*2)+2)
}
