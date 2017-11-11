package main

import (
	"flag"
	"log"
	"os"

	"github.com/thefirstofthe300/ekg/dns"
	"github.com/thefirstofthe300/ekg/fmt"
	"github.com/thefirstofthe300/ekg/processes"
)

func main() {

	var outptr *os.File

	help := flag.Bool("help", false, "Display this help dialog and exit.")
	procs := flag.Bool("processes", false, "Pretty prints the currently running processes")
	dnsdump := flag.Bool("dns", false, "Dumps the state of DNS")
	logfile := flag.String("log-file", "", "Log file to use")
	outfile := flag.String("out-file", "", "Output file for the gathered metrics")
	flag.Parse()

	if *logfile != "" {
		if _, err := os.Stat(*logfile); err == nil {
			os.Rename(*logfile, *logfile+string(".0"))
		}

		logptr, err := os.OpenFile(*logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

		if err != nil {
			log.Fatalf("unable to open log file for writing: %s", err)
		}
		defer logptr.Close()

		log.SetOutput(logptr)
	}

	if *outfile != "" {
		var err error
		if _, err := os.Stat(*outfile); err == nil {
			os.Rename(*outfile, *outfile+string(".0"))
		}

		outptr, err = os.OpenFile(*outfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

		if err != nil {
			log.Fatalf("unable to open log file for writing: %s", err)
		}
		defer outptr.Close()
	} else {
		outptr = os.Stdout
	}

	toFmt := fmt.FmtConfig{
		Processes: nil,
		DNS:       nil,
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

	if *dnsdump {
		resolvconf, err := dns.NewResolvConf()

		if err != nil {
			log.Fatalf("unable to generate resolvConf: %s", err)
		}

		dnsInfo, err := dns.NewDNSConfig(resolvconf, true)

		if err != nil {
			log.Fatalf("unable to generate DNS config: %s", err)
		}

		toFmt.DNS = dnsInfo
	}

	fmt.Printf(outptr, &toFmt)
}
