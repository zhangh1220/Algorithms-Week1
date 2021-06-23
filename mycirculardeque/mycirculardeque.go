package mycirculardeque

import "fmt"

type MyCircularDeque struct {
	head *Node
	tail *Node
	nSize int // 队列大小
	nLen int // 当前队列长度
}

type Node struct {
	Val  int
	Pre  *Node
	Next *Node
}

func Constructor(k int) MyCircularDeque {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head
	mcd := MyCircularDeque{
		head: head,
		tail: tail,
		nSize: k,
		nLen: 0,
	}
	return mcd
}

func (m *MyCircularDeque) InsertFront(value int) bool {
	// 队列是否已满
	if m.IsFull(){
		return false
	}
	// 创建新的节点
	newNode := &Node{
		Val: value,
		Next:m.head.Next,
		Pre: m.head,
	}
	// 插入节点
	m.head.Next.Pre = newNode
	m.head.Next = newNode
	m.nLen++
	return true
}

func (m *MyCircularDeque) InsertLast(value int) bool {
	// 队列是否已满
	if m.IsFull() {
		return false
	}
	// 创建新的节点
	newNode := &Node{
		Val: value,
		Next:m.tail,
		Pre: m.tail.Pre,
	}
	// 插入节点
	m.tail.Pre.Next = newNode
	m.tail.Pre = newNode
	m.nLen++
	return true
}

func (m *MyCircularDeque) DeleteFront() bool {
	if m.nLen == 0 {
		return false
	}
	m.head.Next.Next.Pre = m.head
	m.head.Next = m.head.Next.Next
	m.nLen--
	return true
}

func (m *MyCircularDeque) DeleteLast() bool {
	if m.nLen == 0 {
		return false
	}
	m.tail.Pre.Pre.Next = m.tail
	m.tail.Pre = m.tail.Pre.Pre
	m.nLen--
	return true
}

func (m *MyCircularDeque) GetFront() int {
	if m.nLen == 0 {
		return -1
	}
	return m.head.Next.Val
}

func (m *MyCircularDeque) GetRear() int {
	if m.nLen == 0 {
		return -1
	}
	return m.tail.Pre.Val
}

func (m *MyCircularDeque) IsEmpty() bool {
	if m.nLen == 0 {
		return true
	}
	return false
}

func (m *MyCircularDeque) IsFull() bool {
	if m.nLen == m.nSize {
		return true
	}
	return false
}

func Test() {
	m:= Constructor(3)
	a:=m.InsertLast(1)
	fmt.Println(a)
	b:=m.InsertLast(2)
	fmt.Println(b)
	c:= m.InsertFront(3)
	fmt.Println(c)
	d:=m.InsertFront(4)
	fmt.Println(d)
	e:=m.GetRear()
	fmt.Println(e)
	f:= m.IsFull()
	fmt.Println(f)
	g:= m.DeleteLast()
	fmt.Println(g)
	h:=m.InsertFront(4)
	fmt.Println(h)
	i := m.GetFront()
	fmt.Println(i)
}