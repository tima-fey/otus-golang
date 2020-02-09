package main

import (
	"fmt"
	"os"

	"github.com/tima-fey/otus-golang/emvdir/internal/envdir"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error. Specify env dir and command to execute")
		os.Exit(1)
	}
	exitCode, err := envdir.Envdir(os.Args[1], os.Args[2:])
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitCode)
}
