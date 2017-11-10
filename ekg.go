package main

import (
	"flag"
	"log"
	"os"

	"github.com/thefirstofthe300/ekg/fmt"

	"github.com/thefirstofthe300/ekg/processes"
)

func main() {

	help := flag.Bool("help", false, "Display this help dialog and exit.")
	procs := flag.Bool("processes", false, "Pretty prints the currently running processes")
	flag.Parse()

	toFmt := fmt.FmtConfig{
		Processes: nil,
	}

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

		toFmt.Processes = p
	}

	fmt.Printf(os.Stdout, &toFmt)
}
