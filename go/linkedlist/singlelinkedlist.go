package linkedlist

import (
	"fmt"
)

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

// 链表反转 O(n)
// 参考 http://c.biancheng.net/view/8105.html
func (l *LinkedList) Reverse() {
	// 至少有两个node
	if nil == l.head || nil == l.head.next || nil == l.head.next.next {
		return
	}

	// 迭代反转法
	//var pre *ListNode = nil
	//cur := l.head.next
	//end := l.head.next.next
	//
	//for {
	//	//反转 mid 所指节点的指向
	//	cur.next = pre
	//	//此时判断 end 是否为 NULL，如果成立则退出循环
	//	if nil == end {
	//		break
	//	}
	//	//整体向后移动 3 个指针
	//	pre = cur
	//	cur = end
	//	end = end.next
	//}
	////最后修改 head 头指针的指向
	//l.head.next = cur

	// 头插法反转链表
	//var newHead *ListNode = nil
	//var temp *ListNode = nil
	//head := l.head.next
	//for nil != head {
	//	// 老链表删除头
	//	temp = head
	//	head = head.next
	//
	//	// 头插法加入新链表
	//	temp.next = newHead
	//	newHead = temp
	//}
	////最后修改 l head 头指针的指向
	//l.head.next = newHead

	// 就地反转法
	var beg *ListNode = nil
	var end *ListNode = nil
	head := l.head.next
	beg = head
	end = head.next
	for nil != end {
		// 原始链表删除end节点
		beg.next = end.next
		// end 节点插入头节点
		end.next = head
		head = end
		// end向后移动
		end = beg.next
	}
	l.head.next = head

}

// 判断单链表是否有环 O(n)
func (l *LinkedList) HasCycle() bool {
	if nil != l.head {
		slow := l.head
		fast := l.head
		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// 删除倒数第N个节点 O(n)
func (l *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == l.head || nil == l.head.next {
		return
	}

	fast := l.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}

	if nil == fast {
		return
	}

	slow := l.head
	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
	l.length--
}

// 获取中间节点 O(n)
func (l *LinkedList) FindMiddleNode() *ListNode {
	if nil == l.head || nil == l.head.next {
		return nil
	}
	if nil == l.head.next.next {
		return l.head.next
	}

	slow, fast := l.head, l.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
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

// 递归反转链表
func Reverse(head *ListNode) *ListNode {
	if nil == head || nil == head.next {
		return head
	} else {
		newHead := Reverse(head.next)

		head.next.next = head
		head.next = nil
		return newHead
	}
}

// 两个有序单链表合并 O(n)
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}

	l := &LinkedList{head: &ListNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for nil != curl1 && nil != curl2 {
		if curl1.value.(int) > curl2.value.(int) {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next
	}

	if nil != curl1 {
		cur.next = curl1
	} else if nil != curl2 {
		cur.next = curl2
	}

	return l
}
