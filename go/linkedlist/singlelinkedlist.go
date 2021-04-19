package linkedlist

import "fmt"

//单链表基本操作
//head节点的value无意义
type ListNode struct {
	value interface{}
	next  *ListNode
}

type LinkedList struct {
	length uint
	head   *ListNode
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{value: v, next: nil}
}

func (l *ListNode) GetNext() *ListNode {
	return l.next
}

func (l *ListNode) GetValue() interface{} {
	return l.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{0, NewListNode(0)}
}

func (l *LinkedList) GetLength() uint {
	return l.length
}

func (l *LinkedList) GetHead() *ListNode {
	return l.head
}

// 在某个节点后插入节点 O(1)
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	// 1
	//oldNode := p.next
	//p.next = newNode
	//newNode.next = oldNode
	// 2
	newNode.next = p.next
	p.next = newNode
	l.length++
	return true
}

// 在某个节点前插入节点 O(n)
func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || l.head == p {
		return false
	}
	// 寻找p的前一个节点指针
	cur := l.head.next
	pre := l.head
	for nil != cur {
		if p == cur {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	l.length++
	return true
}

// 在链表头插入 O(1)
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// 在链表尾插入 O(n)
func (l *LinkedList) InsertToTail(v interface{}) bool {
	last := l.head
	for nil != last.next {
		last = last.next
	}
	return l.InsertAfter(last, v)
}

// 查找index的节点 O(n)
func (l *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= l.length {
		return nil
	}
	cur := l.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除节点 O(n)
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	l.length--
	return true
}

//打印链表
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
