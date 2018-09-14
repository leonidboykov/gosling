package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leonidboykov/gosling"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

const (
	defaultOutput    = "public"
	defaultRedirects = "redirects.json"
)

var (
	outputFlag    string
	redirectsFlag string
)

func main() {
	// Parse flags
	flag.StringVar(&outputFlag, "o", defaultOutput, "specify output folder")
	flag.StringVar(&redirectsFlag, "r", defaultRedirects, "specify redirects file")
	flag.Parse()

	redir, err := gosling.NewRedirects(redirectsFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read redirects file: %s", err)
	}

	if err := gosling.BuildRedirects(redir, outputFlag); err != nil {
		fmt.Fprintf(os.Stderr, "unable to build redirects: %s", err)
	}
}
