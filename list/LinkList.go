package list

import "fmt"

type LinkList struct {
	head *node
}


type node struct {
	value int
	next *node
}

func (list *LinkList) AddToHead(value int) {
	var newNode node
	newNode.value = value

	if list.head != nil {
		newNode.next = list.head
	}

	list.head = &newNode
}


func (list *LinkList) IterateList() {
	var item *node
	for item = list.head ; item != nil ; item = item.next {
		fmt.Println(item.value)
	}
}


// O(n)
func (list *LinkList) LastNode() *node {
	var item *node
	var lastNode *node

	for item = list.head ; item != nil ; item = item.next {
		if item.next == nil {
			lastNode = item
			return lastNode
		}
	}

	return lastNode
}


// O(n)
func (list *LinkList) AddToEnd(value int) {
	var newNode node
	newNode.value = value

	lastNode := list.LastNode()

	lastNode.next = &newNode
}


// O(1..n)
func (list *LinkList) NodeWithValue(value int) *node {
	var targetNode *node

	var item *node
	for item = list.head ; item != nil ; item = item.next {
		if item.value == value {
			targetNode = item
			return targetNode
		}
	}

	return targetNode
}


func (list *LinkList) AddAfter(nodeValue int,value int) {
	item := list.NodeWithValue(nodeValue)

	var newNode node
	newNode.value = value
	newNode.next = item.next

	item.next = &newNode
}