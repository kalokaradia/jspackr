package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kalokaradia/jspackr/internal/build"
)

func main() {
	var (
		input  string
		output string
		minify bool
		help   bool
		report bool
	)

	flag.StringVar(&input, "input", "", "Entry file")
	flag.StringVar(&input, "i", "", "Entry file (shorthand)")

	flag.StringVar(&output, "out", "dist/bundle.js", "Output file")
	flag.StringVar(&output, "o", "dist/bundle.js", "Output file (shorthand)")

	flag.BoolVar(&minify, "minify", false, "Minify output")
	flag.BoolVar(&minify, "m", false, "Minify output (shorthand)")

	flag.BoolVar(&report, "report", false, "Show build report")
	flag.BoolVar(&report, "r", false, "Show build report (shorthand)")

	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&help, "h", false, "Show help")

	flag.Usage = func() {
		fmt.Println(`jspackr v0.1.0

Usage:
  jspackr [entry] [options]

Options:
  -i, --input <file>     Entry file
  -o, --out <file>       Output file (default: dist/bundle.js)
  -m, --minify           Minify output
  -r, --report           Show build report
  -h, --help             Show help

Examples:
  jspackr src/index.js
  jspackr -i src/index.js -o dist/app.js -m --report`)
	}

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	// positional arg fallback
	if input == "" && flag.NArg() > 0 {
		input = flag.Arg(0)
	}

	if input == "" {
		fmt.Println("Error: input file is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if err := build.Run(build.Options{
		Input:  input,
		Output: output,
		Minify: minify,
		Report: report,
	}); err != nil {
		fmt.Println("Build failed:", err)
		os.Exit(1)
	}

	fmt.Println("Build finished")
}
