package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.0.1"

var revision = "HEAD"

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}


func main() {
	var showVersion bool
	var showHelp bool

	// options
	fs := flag.NewFlagSet("go-cli-template", flag.ExitOnError)
	fs.BoolVar(&showVersion, "version", false, "show version")
	fs.BoolVar(&showVersion, "v", false, "show version")
	fs.BoolVar(&showHelp, "help", false, "show help")
	fs.BoolVar(&showHelp, "h", false, "show help")
	fs.Usage = func() {
		fmt.Println(`Usage:
 app <path>

Flags:`)
		fs.PrintDefaults()
		fmt.Println(`Repository:
  https://github.com/zztkm/app`)
	}

	fs.Parse(os.Args[1:])

	if showVersion {
		fmt.Printf("version: %s, revision: %s\n", version, revision)
		return
	}
	if showHelp {
		fs.Usage()
		return
	}
}
