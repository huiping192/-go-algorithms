package list

import (
	"testing"
)

func TestAddToHead(t *testing.T) {
	list := LinkList{}

	list.AddToHead(5)

	if list.head.value != 5{
		t.Fatalf("head value must be 5")
	}

	list.AddToHead(10)
	if list.head.value != 10 {
		t.Fatalf("head value must be 10")
	}

	if list.head.next.value != 5 {
		t.Fatalf("second value must be 5")
	}
}

func TestIterateList(t *testing.T) {
	list := LinkList{}

	list.AddToHead(5)
	list.AddToHead(2)
	list.AddToHead(6)

	desc := list.IterateList()

	if desc != "6,2,5," {
		t.Fatalf("wrong value")
	}
}

func TestLastNode(t *testing.T) {
	list := LinkList{}

	list.AddToHead(5)
	list.AddToHead(2)
	list.AddToHead(6)

	if list.LastNode().value != 5 {
		t.Fatalf("last value must be 5")
	}
}

func TestAddToEnd(t *testing.T) {
	list := LinkList{}

	list.AddToEnd(5)
	list.AddToEnd(2)
	list.AddToEnd(6)

	if list.LastNode().value != 6 {
		t.Fatalf("last value must be 5")
	}
}

func TestNodeWithValue(t *testing.T) {
	list := LinkList{}

	list.AddToHead(5)
	list.AddToHead(2)
	list.AddToHead(6)

	node := list.NodeWithValue(2)

	if node.value != 2 {
		t.Fatalf("wrong node")
	}

	if node.next.value != 5 {
		t.Fatalf("wrong node")
	}
}

func TestAddAfter(t *testing.T) {
	list := LinkList{}

	list.AddToHead(5)

	list.AddAfter(5,2)

	if list.head.next.value != 2 {
		t.Fatalf("second value must be 2")
	}


	list.AddAfter(2,6)

	if list.head.next.next.value != 6 {
		t.Fatalf("second value must be 6")
	}
}