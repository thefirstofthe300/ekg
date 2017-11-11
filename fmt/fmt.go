package fmt

import (
	"fmt"
	"io"
	"text/template"

	"github.com/thefirstofthe300/ekg/dns"
	"github.com/thefirstofthe300/ekg/processes"
)

// Config is the data struct to be used when passing data to the output template
type Config struct {
	Processes *processes.Processes
	DNS       *dns.Config
}

// Printf prints the data to the whatever writer it is passed
func Printf(w io.Writer, fc *Config) error {
	funcs := template.FuncMap{
		"trunc": trunc,
	}
	tmpl := template.Must(template.New("").Funcs(funcs).ParseFiles("fmt/templates/processes.tmpl"))

	err := tmpl.ExecuteTemplate(w, "processes.tmpl", fc)

	if err != nil {
		return fmt.Errorf("unable to execute templates: %s", err)
	}

	return nil
}

func trunc(text string) string {
	if len(text) > 50 {
		text = text[:50]
	}
	return text
}
