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
	outputFlag    = flag.String("o", defaultOutput, "specify output folder")
	redirectsFlag = flag.String("r", defaultRedirects, "specify redirects file")
	versionFlag   = flag.Bool("v", false, "print version")
)

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Printf("%s version %s\n", os.Args[0], version)
		os.Exit(0)
	}

	redir, err := gosling.NewRedirects(*redirectsFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read redirects file: %s\n", err)
	}

	if err := gosling.BuildRedirects(redir, *outputFlag); err != nil {
		fmt.Fprintf(os.Stderr, "unable to build redirects: %s\n", err)
	}
}
