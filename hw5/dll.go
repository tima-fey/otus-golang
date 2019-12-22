// The package provide double linked list data type
package main

import "fmt"

// Item define the element from List
type Item struct {
	Prev  *Item
	Next  *Item
	Value interface{}
}

// List define double linked list realization
type List struct {
	Len   int
	First *Item
	Last  *Item
}

// PushFront method push an element to the front of the lust
func (l *List) PushFront(v interface{}) {
	l.Len++
	oldFirst := l.First
	l.First = new(Item)
	l.First.Next = oldFirst
	if oldFirst != nil {
		oldFirst.Prev = l.First
	}
	if l.Last == nil {
		l.Last = l.First
	}
	l.First.Value = v
}

// PushBack method push an element to the back of the lust
func (l *List) PushBack(v interface{}) {
	l.Len++
	oldLast := l.Last
	l.Last = new(Item)
	l.Last.Prev = oldLast
	if oldLast != nil {
		oldLast.Next = l.Last
	}
	if l.First == nil {
		l.First = l.Last
	}
	l.Last.Value = v
}

// RemoveItem method remove an element from list
func (l *List) RemoveItem(i Item) {
	next := i.Next
	prev := i.Prev
	if next != nil {
		next.Prev = prev
	} else {
		l.Last = prev
	}
	if prev != nil {
		prev.Next = next
	} else {
		l.First = next
	}
	l.Len--
}
func main() {
	var myList List
	myList.PushFront(11)
	myList.PushFront(12)
	myList.PushBack(13)
	fmt.Println(myList)
	fmt.Println(myList.First.Value)
	fmt.Println(myList.Last.Value)
	myList.RemoveItem(*myList.First)
	myList.RemoveItem(*myList.First)
	fmt.Println(myList)
	fmt.Println(myList.First.Value)
}
