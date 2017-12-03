// base256: a cli utility for base256 encode/decode. Run with --help for details.
package main // import "ekyu.moe/base256/cmd/base256"

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ekyu.moe/util/cli"
)

const (
	title = "base256"
	usage = "Usage: " + title + ` [OPTION]... <FILE>

Options:
   -d, --decode      Decode data.
   -o, --output      Use as output file. By default, it appends ".asc" to the input
                     filename under dump mode, and ".plain" under load mode. If the
                     input file is suffixed with ".asc", it will strip it. If "-"
                     is specified, write to stdout in raw. Default "-".
`
)

func main() {
	var decodeMode bool
	var outFilename string

	flag.BoolVar(&decodeMode, "d", false, "")
	flag.BoolVar(&decodeMode, "decode", false, "")
	flag.StringVar(&outFilename, "o", "-", "")
	flag.StringVar(&outFilename, "output", "-", "")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}

	flag.Parse()

	inFilename := flag.Arg(0)
	if len(inFilename) == 0 {
		inFilename = "-"
	}

	var err error
	if decodeMode {
		if len(outFilename) == 0 && inFilename != "-" {
			if strings.HasSuffix(strings.ToLower(inFilename), ".asc") {
				outFilename = strings.TrimSuffix(inFilename, ".asc")
			} else {
				outFilename = inFilename + ".plain"
			}
			outFilename = inFilename + ".asc"
		}
		err = load(outFilename, inFilename)
	} else {
		if len(outFilename) == 0 && inFilename != "-" {
			outFilename = inFilename + ".asc"
		}
		err = dump(outFilename, inFilename)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if err != cli.ErrAbortedByUser && outFilename != "-" {
			os.Remove(outFilename)
		}
		os.Exit(1)
	}
}
