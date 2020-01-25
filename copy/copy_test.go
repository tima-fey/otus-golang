package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCase struct {
	offset int
	limit  int
	custom bool
	answer string
}

var TestCases = []TestCase{
	TestCase{
		offset: 0,
		limit:  0,
		custom: false,
		answer: "0123456",
	},
	TestCase{
		offset: 0,
		limit:  0,
		custom: true,
		answer: "0123456",
	},
	TestCase{
		offset: 1,
		limit:  0,
		custom: false,
		answer: "123456",
	},
	TestCase{
		offset: 1,
		limit:  0,
		custom: true,
		answer: "123456",
	},
	TestCase{
		offset: 1,
		limit:  2,
		custom: false,
		answer: "12",
	},
	TestCase{
		offset: 1,
		limit:  2,
		custom: true,
		answer: "12",
	},
}

func testHelper() (*os.File, *os.File, func()) {
	file1, err := ioutil.TempFile("", "test")
	if err != nil {
		fmt.Println(err)
	}
	// defer os.Remove(file1.Name())
	file2, err := ioutil.TempFile("", "test")
	if err != nil {
		fmt.Println(err)
	}
	// defer os.Remove(file2.Name())
	file1D, err := os.Create(file1.Name())
	if err != nil {
		fmt.Println(err)
	}
	defer file1D.Close()
	file1D.WriteString("0123456")
	return file1, file2, func() {
		os.Remove(file1.Name())
		os.Remove(file2.Name())
	}
}

func TestBaseCopy(t *testing.T) {
	for _, aCase := range TestCases {
		file1, file2, cleanup := testHelper()
		defer cleanup()
		Copy(aCase.offset, aCase.limit, file1.Name(), file2.Name(), aCase.custom)
		file2Data, _ := ioutil.ReadFile(file2.Name())
		require.Equal(t, aCase.answer, string(file2Data))
	}
}
