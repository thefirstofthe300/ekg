package main

import (
	"flag"
	"log"
	"os"

	"github.com/thefirstofthe300/ekg/system/processes"
)

func main() {

	help := flag.Bool("help", false, "Display this help dialog and exit.")
	procs := flag.Bool("processes", false, "Pretty prints the currently running processes")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("I started and help doesn't exist.")

	if *procs {
		p, err := processes.New()

		if err != nil {
			log.Fatalf("could not get processes: %s", err)
		}

		p.Write(os.Stdout)
	}
}
