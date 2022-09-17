package main

import "log"

type tree struct {
	delimitator string
	path        string
	pack        string
	folders     []string
	files       []string
	nodeUp      *tree
	nodeDown    *tree
}

type DoubleList struct {
	head *tree
	tail *tree
	len  int
}

func printTree(t *tree) {
	if t != nil {
		log.Print(t)
	}
}

func (dl *DoubleList) push(t *tree) {
	if dl.head == nil {
		dl.head = t
		dl.tail = t
	} else {
		t.nodeUp = dl.head
		dl.head.nodeDown = t
		dl.head = t
	}
	dl.len++
}

func (dl *DoubleList) append(t *tree) {
	if dl.head == nil {
		dl.head = t
		dl.tail = t
	} else {
		dl.tail.nodeUp = t
		t.nodeDown = dl.tail
		dl.tail = t
	}
	dl.len++
}

func (dl *DoubleList) displayHead() {
	list := dl.head
	if list == dl.tail {
		log.Panic("Tail == Head")
	}
	for list != nil {
		printTree(list)
		list = list.nodeUp
	}
}

func (dl *DoubleList) displayTail() {
	list := dl.tail
	log.Print("OK")
	for list != nil {
		printTree(list)
		list = list.nodeDown
	}
}
