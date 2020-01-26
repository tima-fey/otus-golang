package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Helper() (string, func()) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	return dir, func() {
		defer os.RemoveAll(dir)
	}
}

func TestReadDir(t *testing.T) {
	dir, cleanup := Helper()
	defer cleanup()
	tempFile, err := os.Create(dir + "/" + "TEST")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	tempFile.WriteString("TOST")
	envs, _ := ReadDir(dir)
	require.Equal(t, "TOST", envs["TEST"])
}
func TestReadDirEmptyFile(t *testing.T) {
	dir, cleanup := Helper()
	defer cleanup()
	tempFile, err := os.Create(dir + "/" + "TEST")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	envs, _ := ReadDir(dir)
	require.Equal(t, "", envs["TEST"])
}
