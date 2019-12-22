package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPushFront(t *testing.T) {
	var myList List
	myList.PushFront(2)
	require.Equal(t, 2, myList.First.Value)
	require.Equal(t, 2, myList.Last.Value)
	require.Equal(t, 1, myList.Len)
	myList.PushFront(3)
	require.Equal(t, 3, myList.First.Value)
	require.Equal(t, 2, myList.Last.Value)
	require.Equal(t, 2, myList.Len)
	myList.PushFront(4)
	require.Equal(t, 4, myList.First.Value)
	require.Equal(t, 2, myList.Last.Value)
	require.Equal(t, 3, myList.Len)
}

func TestPushBack(t *testing.T) {
	var myList List
	myList.PushBack(2)
	require.Equal(t, 2, myList.First.Value)
	require.Equal(t, 2, myList.Last.Value)
	require.Equal(t, 1, myList.Len)
	myList.PushBack(3)
	require.Equal(t, 2, myList.First.Value)
	require.Equal(t, 3, myList.Last.Value)
	require.Equal(t, 2, myList.Len)
	myList.PushBack(4)
	require.Equal(t, 2, myList.First.Value)
	require.Equal(t, 4, myList.Last.Value)
	require.Equal(t, 3, myList.Len)
}

func TestRemoveItem(t *testing.T) {
	var myList List
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.RemoveItem(*myList.Last)
	require.Equal(t, 1, myList.First.Value)
	require.Equal(t, 2, myList.Last.Value)
	require.Equal(t, 2, myList.Len)
	myList.RemoveItem(*myList.First)
	require.Equal(t, 2, myList.First.Value)
	require.Equal(t, 2, myList.Last.Value)
	require.Equal(t, 1, myList.Len)
	myList.RemoveItem(*myList.First)
	var zeroList List
	require.Equal(t, zeroList.First, myList.First)
	require.Equal(t, zeroList.First, *myList.Last)
	require.Equal(t, 0, myList.Len)
}
