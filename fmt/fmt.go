package fmt

import (
	"fmt"
	"io"
	"text/template"

	"github.com/thefirstofthe300/ekg/dns"
	"github.com/thefirstofthe300/ekg/processes"
)

// FmtConfig is the data struct to be used when passing data to the output template
type FmtConfig struct {
	Processes *processes.Processes
	DNS       *dns.DNSConfig
}

// Printf prints the data to the whatever writer it is passed
func Printf(w io.Writer, fc *FmtConfig) error {
	tmpl := template.Must(template.ParseGlob("fmt/templates/*"))

	err := tmpl.Execute(w, fc)

	if err != nil {
		return fmt.Errorf("unable to execute templates: %s", err)
	}

	return nil
}
