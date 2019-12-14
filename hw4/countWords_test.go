package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrepareString(t *testing.T) {
	require.Equal(t, []string{"test"}, prepareStrings("test"))
	require.Equal(t, []string{"test"}, prepareStrings("teSt"))
	require.Equal(t, []string{"one", "two"}, prepareStrings(" one two"))
	require.Equal(t, []string{"one", "two"}, prepareStrings("one  two"))
	require.Equal(t, []string{"one", "two"}, prepareStrings("one, two"))
	require.Equal(t, []string{"one", "two"}, prepareStrings("one - two"))
	require.Equal(t, []string{"two", "three", "three", "one", "two", "three"}, prepareStrings(" two  three, three one Two - three"))
}

func TestCountWords(t *testing.T) {
	testData := []wordCount{{"two", 2}, {"one", 1}}
	require.Equal(t, testData, countWords([]string{"one", "two", "two"}))
	testData = []wordCount{{"three", 3}, {"two", 2}, {"one", 1}}
	require.Equal(t, testData, countWords([]string{"two", "three", "three", "one", "two", "three"}))
}

func TestGetTopCommon(t *testing.T) {
	testData := []wordCount{{"three", 3}, {"two", 2}}
	require.Equal(t, testData, getTopCommon(" two  three, three one Two - three", 2))
}
