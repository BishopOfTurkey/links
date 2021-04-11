package main

import (
	"flag"
	"linkstore"
	"os"
)

var (
	pword *string = flag.String("code", "", "Code used verify identity")
	file  *string = flag.String("file", "links.csv", "CSV to store links in.")
)

func main() {
	flag.Parse()
	if *pword == "" {
		flag.Usage()
		os.Exit(1)
	}
	linkstore.Server(*pword, *file)
}
