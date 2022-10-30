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
	head   *ListItem
	tail   *ListItem
	length int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	switch {
	case l.head == nil:
		l.head = newItem
		l.tail = newItem
	case l.head != nil:
		head := l.head
		head.Prev = newItem
		newItem.Next = head
		l.head = newItem
	}
	l.length += 1
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	switch {
	case l.tail == nil:
		l.head = newItem
		l.tail = newItem
	case l.tail != nil:
		tail := l.tail
		tail.Next = newItem
		newItem.Prev = tail
		l.tail = newItem
	}
	l.length += 1
	return l.tail
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.head == l.tail && l.head == i: // one element in list
		l.head = nil
		l.tail = nil
	case l.head == i:
		l.head = i.Next
		l.head.Prev = nil
	case l.tail == i:
		l.tail = i.Prev
		l.tail.Next = nil
	default:
		prev := i.Prev
		next := i.Next
		prev.Next = next
		next.Prev = prev
	}
	l.length -= 1
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case l.head == i:
	default:
		// delete item from list
		l.Remove(i)
		// change pointers of moved item
		i.Next = l.head
		i.Prev = nil
		l.head = i
		// change prev pointer of 2nd item
		l.head.Next.Prev = i
		// add +1 to list as we just moved it
		l.length += 1
	}
}
