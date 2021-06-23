package mergetwolists

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}//创建起始链表头节点
	result := preHead//保存起始链表头 用于返回
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1 //建立联系
			l1 = l1.Next//目标链表向后移动
		} else {
			preHead.Next = l2//建立联系
			l2 = l2.Next//目标链表向后移动
		}
		preHead = preHead.Next//向后移动
	}
	//处理剩余节点
	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}
	return result.Next
}

//Test
func NewLists() *ListNode {
	l := &ListNode{}
	lHead := l
	for i:=1;i<4;i++ {
		lHead.Next= &ListNode{
			Val: i,
			Next :nil,
		}
		lHead = lHead.Next
	}
	return l.Next
}

func Test() {
	l1 := NewLists()
	l2:= NewLists()
	l3 := mergeTwoLists(l1,l2)
	for l3!= nil {
		fmt.Print( l3.Val)
		l3= l3.Next
	}
	fmt.Println()
}