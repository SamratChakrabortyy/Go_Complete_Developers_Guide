package main

import "fmt"

type node struct {
	val  int
	next *node
}

func addToTail(tail *node, val int) *node {
	n := node{val: val}
	tail.next = &n
	return tail.next
}

func addToHead(head *node, val int) *node {
	n := node{val: val, next: head}
	return &n
}

func (head *node) print() {

	if head == nil {
		return
	}

	curr := head
	for curr.next != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Printf("%d", curr.val)
}
