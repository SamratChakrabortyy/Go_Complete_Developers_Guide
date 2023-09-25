package main

func main() {
	head := &node{val: 0}

	tail := head
	tail = addToTail(tail, 1)
	tail = addToTail(tail, 2)
	tail = addToTail(tail, 3)
	head = addToHead(head, -1)

	head.print()
}
