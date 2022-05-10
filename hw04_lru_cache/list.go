package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	switch {
	case l.head == nil:
		l.head = newItem
	default:
		oldTail := l.tail
		fmt.Println(oldTail)
		l.tail.Next = newItem
		newItem.Prev = l.tail
	}
	l.tail = newItem
	l.len++
	return newItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	switch {
	case l.tail == nil:
		l.tail = newItem
	default:
		oldHead := l.head
		newItem.Next = oldHead
		oldHead.Prev = newItem
	}
	l.head = newItem
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.head == i:
		itemNext := i.Next
		itemNext.Prev = nil
		l.head = itemNext
	case l.tail == i:
		itemPrev := i.Prev
		itemPrev.Next = nil
		l.tail = itemPrev
	default:
		itemPrev := i.Prev
		itemNext := i.Next
		itemPrev.Next = i.Next
		itemNext.Prev = itemPrev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case l.head == i:
		fmt.Println("No action required - item is already at beginning")
	case l.tail == i:
		itemPrev := i.Prev
		itemPrev.Next = nil
		l.tail = itemPrev

		oldHead := l.head
		oldHead.Prev = i
		i.Prev = nil
		i.Next = oldHead
		l.head = i
	default:
		itemPrev := i.Prev
		itemNext := i.Next
		itemPrev.Next = i.Next
		itemNext.Prev = itemPrev

		oldHead := l.head
		oldHead.Prev = i
		i.Prev = nil
		i.Next = oldHead
		l.head = i

	}
}
