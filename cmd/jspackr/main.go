package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kalokaradia/jspackr/internal/build"
	"github.com/kalokaradia/jspackr/internal/watch"
)

func main() {
	var (
		input     string
		output    string
		minify    bool
		report    bool
		help      bool
		sourceMap string
		isWatch 		bool
	)

	// --- CLI flags ---
	flag.BoolVar(&isWatch, "watch", false, "Enable watch mode")
	flag.BoolVar(&isWatch, "w", false, "Enable watch mode (shorthand)")


	flag.StringVar(&input, "input", "", "Entry file")
	flag.StringVar(&input, "i", "", "Entry file (shorthand)")

	flag.StringVar(&output, "out", "dist/bundle.js", "Output file")
	flag.StringVar(&output, "o", "dist/bundle.js", "Output file (shorthand)")

	flag.BoolVar(&minify, "minify", false, "Minify output")
	flag.BoolVar(&minify, "m", false, "Minify output (shorthand)")

	flag.BoolVar(&report, "report", false, "Show build report")
	flag.BoolVar(&report, "r", false, "Show build report (shorthand)")

	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&help, "h", false, "Show help (shorthand)")

	flag.StringVar(&sourceMap, "source", "none", "Source map: none, l (linked), in (inline)")
	flag.StringVar(&sourceMap, "s", "none", "Source map: none, l (linked), in (inline) (shorthand)")

	flag.Usage = func() {
		fmt.Println(`jspackr v0.2.0

Usage:
  jspackr [entry] [options]

Options:
  -i, --input <file>       Entry file
  -o, --out <file>         Output file (default: dist/bundle.js)
  -m, --minify             Minify output
  -r, --report             Show build report
  -s, --source <mode>      Source map: none, l (linked), in (inline)
	-w, --watch							 Enable watch mode
  -h, --help               Show help

Examples:
  jspackr src/index.js
  jspackr -i src/index.js -o dist/app.js -m -r -s l`)
	}

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if input == "" && flag.NArg() > 0 {
		input = flag.Arg(0)
	}

	if input == "" {
		color.New(color.FgRed).Println("Error: input file is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if isWatch {
			color.New(color.FgBlue).Println("👀 Watch mode enabled...")
			if err := watch.WatchFiles(input, build.Options{
					Input:     input,
					Output:    output,
					Minify:    minify,
					Report:    report,
					SourceMap: sourceMap,
			}); err != nil {
					color.New(color.FgRed).Println("❌ Watcher error:", err)
					os.Exit(1)
			}
			return
	}


	// --- Run build ---
	err := build.Run(build.Options{
		Input:     input,
		Output:    output,
		Minify:    minify,
		Report:    report,
		SourceMap: sourceMap,
	})

	if err != nil {
		color.New(color.FgRed).Printf("❌ Build failed: %s\n", err)
		os.Exit(1)
	}

	// Success
	color.New(color.FgGreen).Println("✅ Build finished")
}
