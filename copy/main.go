package main

import (
	"flag"
	"log"
	"os"

	"github.com/tima-fey/otus-golang/copy/copy"
)

func main() {
	var offset int
	var sourceName string
	var destinationName string
	var limit int
	var isCustom bool
	flag.IntVar(&offset, "offset", 0, "offset in input file")
	flag.StringVar(&sourceName, "source", "", "osource file")
	flag.StringVar(&destinationName, "destination", "", "destination file")
	flag.IntVar(&limit, "limit", 0, "limit of copy")
	flag.BoolVar(&isCustom, "custom", false, "use custom copy func")
	flag.Parse()
	err := copy.Copy(offset, limit, sourceName, destinationName, isCustom)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
