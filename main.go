package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const version = "0.0.1"

var revision = "HEAD"

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func main() {
	var showVersion bool
	var showHelp bool
	var output string

	// options
	fs := flag.NewFlagSet("go-cli-template", flag.ExitOnError)
	fs.BoolVar(&showVersion, "version", false, "show version")
	fs.BoolVar(&showVersion, "v", false, "show version")
	fs.BoolVar(&showHelp, "help", false, "show help")
	fs.BoolVar(&showHelp, "h", false, "show help")
	fs.StringVar(&output, "o", "", "write output to file")
	fs.StringVar(&output, "output", "", "write output to file")
	fs.Usage = func() {
		fmt.Println(`Usage:
 sjis2utf8 [options] [<input>]

Flags:`)
		fs.PrintDefaults()
		fmt.Println(`Repository:
  https://github.com/zztkm/sjis2utf8`)
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

	if fs.NArg() < 1 {
		log.Fatal("[ERROR] path is required")
	}

	path := fs.Arg(0)

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
	defer f.Close()

	var writer io.Writer
	if output != "" {
		f, err := os.Create(output)
		if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}
		defer f.Close()
		writer = f
	} else {
		writer = os.Stdout
	}

	r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	w := csv.NewWriter(writer)
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}
		if err := w.Write(records); err != nil {
			log.Fatalf("[ERROR] %s", err)
		}
	}
}
