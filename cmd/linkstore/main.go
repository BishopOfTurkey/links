package main

import (
	"flag"
	"linkstore"
	"os"
)

var (
	pword *string = flag.String("code", "", "Code used verify identity")
)

func main() {
	flag.Parse()
	if *pword == "" {
		flag.Usage()
		os.Exit(1)
	}
	linkstore.Server(*pword)
}
