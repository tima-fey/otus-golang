// The package provide double linked list data type
package main

import (
	"errors"
	"fmt"
)

// Item define the element from List
type Item struct {
	prev       *Item
	next       *Item
	value      interface{}
	parentList *List
}

func (i Item) Prev() Item {
	return *i.prev
}
func (i Item) Next() Item {
	return *i.next
}
func (i Item) Value() interface{} {
	return i.value
}

// List define double linked list realization
type List struct {
	len   int
	first *Item
	last  *Item
}

func (l List) Len() int {
	return l.len
}
func (l List) First() Item {
	return *l.first
}
func (l List) Last() Item {
	return *l.last
}

// PushFront method push an element to the front of the lust
func (l *List) PushFront(v interface{}) {
	l.len++
	l.first = &Item{next: l.first, value: v, parentList: l}
	if l.first.next != nil {
		l.first.next.prev = l.first
	}
	if l.last == nil {
		l.last = l.first
	}
}

// PushBack method push an element to the back of the lust
func (l *List) PushBack(v interface{}) {
	l.len++
	l.last = &Item{value: v, prev: l.last, parentList: l}
	if l.last.prev != nil {
		l.last.prev.next = l.last
	}
	if l.first == nil {
		l.first = l.last
	}
}

// RemoveItem method remove an element from list
func (l *List) RemoveItem(i Item) error {
	if l != i.parentList {
		return errors.New("Can't remove itemv from another list")
	}
	next := i.next
	prev := i.prev
	if next != nil {
		next.prev = prev
	} else {
		l.last = prev
	}
	if prev != nil {
		prev.next = next
	} else {
		l.first = next
	}
	l.len--
	return nil
}
func main() {
	var myList List
	myList.PushFront(11)
	myList.PushFront(12)
	myList.PushBack(13)
	fmt.Println(myList)
	fmt.Println(myList.first.value)
	fmt.Println(myList.last.value)
	myList.RemoveItem(*myList.first)
	myList.RemoveItem(*myList.first)
	fmt.Println(myList)
	fmt.Println(myList.first.value)
}
