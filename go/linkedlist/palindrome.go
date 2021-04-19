package linkedlist

// 链表回文数

/*
	思路1：开一个栈存放链表前半段
	时间复杂度：O(N)
	空间复杂度：O(N)
*/
func isPalindrome1(l *LinkedList) bool {
	length := l.length
	if 0 == length {
		return false
	}
	if 1 == length {
		return true
	}
	s := make([]string, 0, length/2)
	cur := l.head
	for i := uint(1); i <= length; i++ {
		cur = cur.next
		// 如果链表有奇数个节点，中间的直接忽略
		if length%2 != 0 && i == (length/2+1) {
			continue
		}
		if i <= length/2 { //前一半入栈
			s = append(s, cur.GetValue().(string))
		} else { //后一半与前一半进行对比
			if s[length-i] != cur.GetValue().(string) {
				return false
			}
		}
	}
	return true
}

/*
	思路2
	找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
	时间复杂度：O(N)
*/
func isPalindrome2(l *LinkedList) bool {
	length := l.length
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	var isPalindrome = true

	// 前一半链表反转
	step := length / 2
	var pre *ListNode = nil
	cur := l.head.next
	next := l.head.next.next
	for i := uint(1); i <= step; i++ {
		tmp := cur.GetNext()
		cur.next = pre
		pre = cur
		cur = tmp
		next = cur.GetNext()
	}
	mid := cur

	// 从中间遍历左右两边链表 判断是否为回文
	var left, right *ListNode = pre, nil
	if length%2 != 0 { //
		right = mid.GetNext()
	} else {
		right = mid
	}
	for nil != left && nil != right {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	//复原链表
	cur = pre
	pre = mid
	for nil != cur {
		next = cur.GetNext()
		cur.next = pre
		pre = cur
		cur = next
	}
	l.head.next = pre

	return isPalindrome
}
