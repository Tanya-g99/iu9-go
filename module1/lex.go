package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type AVLNode struct {
	key    string
	Value  int
	height int
	left   *AVLNode
	right  *AVLNode
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *AVLNode) rebalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *AVLNode) add(key string, value int) *AVLNode {
	if n == nil {
		return &AVLNode{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		n.Value = value
	}
	return n.rebalanceTree()
}

func (n *AVLNode) search(key string) (int, bool) {
	if n == nil {
		return 0, false
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n.Value, true
	}
}

/*
func (t *AVLTree) DisplayInOrder() {
	t.root.displayNodesInOrder()
}

func (n *AVLNode) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.key, " ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
} */

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Assign(s string, x int) {
	t.root = t.root.add(s, x)
}

func (t *AVLTree) Lookup(s string) (x int, exists bool) {
	return t.root.search(s)

}

type AssocArray interface {
	//DisplayInOrder()
	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

func lex(sentence string, array AssocArray) []int {
	var res []int
	words := strings.Fields(sentence)
	i := 0
	for _, word := range words {
		node, ok := array.Lookup(word)
		if ok {
			res = append(res, node)
		} else {
			i++
			array.Assign(word, i)
			res = append(res, i)
			//array.DisplayInOrder()
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	//fmt.Println(str)
	tree := &AVLTree{}
	t := AssocArray(tree)
	array := lex(str, t)
	for _, v := range array {
		fmt.Printf("%d ", v)
	}

}
