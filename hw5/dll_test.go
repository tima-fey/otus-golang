package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPushFront(t *testing.T) {
	var myList List
	myList.PushFront(2)
	require.Equal(t, 2, myList.first.value)
	require.Equal(t, 2, myList.last.value)
	require.Equal(t, 1, myList.len)
	myList.PushFront(3)
	require.Equal(t, 3, myList.first.value)
	require.Equal(t, 2, myList.last.value)
	require.Equal(t, 2, myList.len)
	myList.PushFront(4)
	require.Equal(t, 4, myList.first.value)
	require.Equal(t, 2, myList.last.value)
	require.Equal(t, 3, myList.len)
}

func TestPushBack(t *testing.T) {
	var myList List
	myList.PushBack(2)
	require.Equal(t, 2, myList.first.value)
	require.Equal(t, 2, myList.last.value)
	require.Equal(t, 1, myList.len)
	myList.PushBack(3)
	require.Equal(t, 2, myList.first.value)
	require.Equal(t, 3, myList.last.value)
	require.Equal(t, 2, myList.len)
	myList.PushBack(4)
	require.Equal(t, 2, myList.first.value)
	require.Equal(t, 4, myList.last.value)
	require.Equal(t, 3, myList.len)
}

func TestRemoveItem(t *testing.T) {
	var myList List
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	require.Nil(t, myList.RemoveItem(*myList.last))
	require.Equal(t, 1, myList.first.value)
	require.Equal(t, 2, myList.last.value)
	require.Equal(t, 2, myList.len)
	require.Nil(t, myList.RemoveItem(*myList.first))
	require.Equal(t, 2, myList.first.value)
	require.Equal(t, 2, myList.last.value)
	require.Equal(t, 1, myList.len)
	require.Nil(t, myList.RemoveItem(*myList.first))
	var zeroList List
	require.Equal(t, zeroList.first, myList.first)
	require.Equal(t, zeroList.first, myList.last)
	require.Equal(t, 0, myList.len)
	var oneMoreList List
	oneMoreList.PushBack(1)
	myList.PushBack(1)
	require.NotNil(t, myList.RemoveItem(*oneMoreList.last))

}

func TestPrev(t *testing.T) {
	var myList List
	myList.PushBack(1)
	myList.PushBack(2)
	require.Equal(t, 1, myList.Last().Prev().Value())
}
func TestNext(t *testing.T) {
	var myList List
	myList.PushBack(1)
	myList.PushBack(2)
	require.Equal(t, 2, myList.First().Next().Value())
}

func TestValue(t *testing.T) {
	var myList List
	myList.PushBack(1)
	require.Equal(t, 1, myList.First().Value())
}

func TestLen(t *testing.T) {
	var myList List
	myList.PushBack(1)
	require.Equal(t, 1, myList.Len())
}
func TestFirst(t *testing.T) {
	var myList List
	myList.PushBack(1)
	myList.PushBack(2)
	require.Equal(t, 1, myList.First().Value())
}
func TestLast(t *testing.T) {
	var myList List
	myList.PushBack(1)
	myList.PushBack(2)
	require.Equal(t, 2, myList.Last().Value())
}
