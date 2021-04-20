package linkedlist

import (
	"fmt"
	"testing"
)

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestIReverse(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	l.Reverse()
	l.Print()
}

func TestIReverse1(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	newHead := Reverse(l.head.next)
	format := ""
	for nil != newHead {
		format += fmt.Sprintf("%+v", newHead.GetValue())
		newHead = newHead.next
		if nil != newHead {
			format += "->"
		}
	}
	fmt.Println(format)
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(10))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}

func TestHasCycle(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 6; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	t.Log(l.HasCycle())
	l.head.next.next.next.next.next.next = l.head.next.next.next
	t.Log(l.HasCycle())
}

func TestDeleteBottomN(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	l.DeleteBottomN(3)
	l.Print()
}

func TestFindMiddleNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	t.Log(l.FindMiddleNode())
	l.DeleteBottomN(7)
	l.Print()
	t.Log(l.FindMiddleNode())
}

func TestMergeSortedList(t *testing.T) {
	n5 := &ListNode{value: 9}
	n4 := &ListNode{value: 7, next: n5}
	n3 := &ListNode{value: 5, next: n4}
	n2 := &ListNode{value: 3, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l1 := &LinkedList{head: &ListNode{next: n1}}

	n10 := &ListNode{value: 10}
	n9 := &ListNode{value: 8, next: n10}
	n8 := &ListNode{value: 6, next: n9}
	n7 := &ListNode{value: 4, next: n8}
	n6 := &ListNode{value: 2, next: n7}
	l2 := &LinkedList{head: &ListNode{next: n6}}

	MergeSortedList(l1, l2).Print()
}
