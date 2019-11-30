package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time_my, err := ntp.Time("01.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(3)
	}
	fmt.Printf("%v\n", time_my)
}
