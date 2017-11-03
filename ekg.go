package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	help := flag.Bool("help", false, "Display this help dialog and exit.")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("I started and help doesn't exist.")
}
