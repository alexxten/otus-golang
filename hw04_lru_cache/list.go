package hw04lrucache

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
		l.tail.Next = newItem
	}
	l.tail = newItem
	l.len++
	return newItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	switch {
	case l.head == nil:
		l.head = newItem
	default:
		newItem.Next = l.head
		l.head = newItem
	}
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.head == i:
		l.head = l.head.Next
	case l.tail == i:
		l.tail = l.tail.Prev
	default:
		i.Prev.Next = i.Next
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.len--
}