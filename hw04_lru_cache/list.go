package hw04lrucache

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type list struct {
	front, back *ListItem
	length      int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if l.length == 0 {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Next = l.front
		l.front.Prev = newItem
		l.front = newItem
	}
	l.length++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if l.length == 0 {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Prev = l.back
		l.back.Next = newItem
		l.back = newItem
	}
	l.length++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}

	l.Remove(i)
	i.Prev = nil
	i.Next = l.front
	if l.front != nil {
		l.front.Prev = i
	}
	l.front = i
	if l.back == nil {
		l.back = i
	}
	l.length++
}
