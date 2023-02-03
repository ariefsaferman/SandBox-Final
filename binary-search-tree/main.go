package main

import (
	"errors"
	"fmt"
)

type node struct {
	value int
	left *node
	right *node
}

func newNode (data int) *node {
	return &node {value: data}
}

func (n *node) Insert(value int) error {
	if n == nil {
		return errors.New("tree is nil")
	}

	if n.value == value {
		return errors.New("this node value already exists")
	}

	if n.value > value {
		if n.left == nil {
			n.left = &node{value: value}
			return nil
		}
		return n.left.Insert(value)
	}

	if n.value < value {
		if n.right == nil {
			n.right = &node{value: value}
			return nil
		}
		return n.right.Insert(value)
	}
	return nil
}

func (n *node) Delete(value int) {
	n.remove(value); 
}

func (n *node) remove(value int) *node {
	if n == nil {
		return nil 
	}

	if value < n.value {
		n.left = n.left.remove(value); 
		return n;
	}

	if value > n.value {
		n.right = n.right.remove(value); 
		return n; 
	}

	if n.left == nil && n.right == nil {
		n = nil; 
		return n; 
	}

	if n.left == nil {
		n = n.right; 
		return n; 
	}

	if n.right == nil {
		n = n.left;
		return n;
	}

	smallestValOnRight := n.right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	n.value = smallestValOnRight.value
	n.right = n.right.remove(n.value)
	return n
}

func (n *node) FindMin() int {
	if n.left == nil {
		return n.value
	}
	return n.left.FindMin()
}

func (n *node) FindMax() int {
	if n.right == nil {
		return n.value
	}
	return n.right.FindMax()
}

func (n *node) PrintInorder() {
	if n == nil {
		return
	}

	n.left.PrintInorder()
	fmt.Println(n.value)
	n.right.PrintInorder()
}


func main()  {
	node := newNode(10); 
	node.Insert(5); 
	node.Insert(15); 
	node.PrintInorder();

	fmt.Println("After delete 5");
	node.Delete(5);
	node.PrintInorder();
}

