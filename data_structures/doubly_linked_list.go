package ds

import (
	"errors"
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type DLL struct {
	head   *Node
	length int
}

func CreateDLL() DLL {

	return DLL{
		length: 0,
		head:   nil,
	}
}

func (l *DLL) prepend(item int) {
	node := Node{
		Val: item,
	}
	l.length++
	if l.head == nil {
		l.head = &node
		return
	}

	node.Next = l.head
	l.head.Prev = &node
	l.head = &node
}

func (l *DLL) insertAt(item int, idx int) error {
	if idx > l.length {
		return errors.New("attempting to insert at idx larger than DLL size")
	}

	// index 0 is just a prepend
	if idx == 0 {
		l.prepend(item)
		return nil
	}

	cur := l.head

	for i := 1; i < idx; i++ {
		cur = cur.Next
	}

	node := Node{
		Val:  item,
		Next: cur.Next,
		Prev: cur,
	}
	if cur.Next != nil {
		cur.Next.Prev = &node
	}
	cur.Next = &node
	l.length++

	return nil
}

func (l *DLL) append(item int) error {
	node := Node{
		Val: item,
	}

	if l.head == nil {
		l.head = &node
		return nil
	}

	err := l.insertAt(item, l.length)

	if err != nil {
		return err
	}

	return nil
}

func (l *DLL) remove(item int) {
}
func (l *DLL) removeAt(item int, idx int) {

}
func (l *DLL) getAt(idx int) (int, error) {
	if idx >= l.length {
		return -1, errors.New("attempting to get at idx larger than DLL size")
	}

	cur := l.head
	for i := 0; i < idx; i++ {
		cur = cur.Next
	}

	return cur.Val, nil
}
