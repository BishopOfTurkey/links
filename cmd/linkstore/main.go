package main

import (
	"flag"
	"linkstore"
	"log"
	"os"
)

var (
// pword *string = flag.String("code", "", "Code used verify identity")
// file  *string = flag.String("file", "links.csv", "CSV to store links in.")
)

func main() {
	flag.Parse()

	pword := os.Getenv("ACCESS_CODE")
	if pword != "" {
		log.Fatalln("Must set an ACCESS_CODE")
	}

	file := os.Getenv("LINKS_FILE")
	if file != "" {
		file = "links.csv"
	}

	linkstore.Server(pword, file)
}
