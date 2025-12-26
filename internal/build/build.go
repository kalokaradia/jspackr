package build

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/fatih/color"
)

// Options
type Options struct {
	Input     string
	Output    string
	Minify    bool
	Report    bool
	SourceMap string
}

// Report structure
type MetaFile struct {
	Inputs map[string]struct {
		Bytes int `json:"bytes"`
	} `json:"inputs"`
	Outputs map[string]struct {
		Bytes int `json:"bytes"`
	} `json:"outputs"`
}

func printReport(meta string) {
	var m MetaFile
	_ = json.Unmarshal([]byte(meta), &m)

	type item struct {
		Path  string
		Bytes int
	}

	var items []item
	for path, v := range m.Inputs {
		items = append(items, item{path, v.Bytes})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Bytes > items[j].Bytes
	})

	green := color.New(color.FgGreen).SprintFunc()
	fmt.Println("\nTop contributors:")

	limit := 5
	if len(items) < limit {
		limit = len(items)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("- %-40s %6s KB\n", items[i].Path, green(items[i].Bytes/1024))
	}
}


// Run build
func Run(opts Options) error {
	if opts.Input == "" {
		return errors.New("input file is required")
	}

	dir := filepath.Dir(opts.Output)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	var sm api.SourceMap
	switch opts.SourceMap {
	case "l":
		sm = api.SourceMapLinked
	case "in":
		sm = api.SourceMapInline
	default:
		sm = api.SourceMapNone
	}

	// Build
	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{opts.Input},
		Bundle:            true,
		MinifyWhitespace:  opts.Minify,
		MinifyIdentifiers: opts.Minify,
		MinifySyntax:      opts.Minify,
		Outfile:           opts.Output,
		Write:             true,
		Platform:          api.PlatformBrowser,
		Metafile:          opts.Report,
		Sourcemap:         sm,
	})

	if len(result.Errors) > 0 {
		return errors.New(result.Errors[0].Text)
	}

	if opts.Report && result.Metafile != "" {
		printReport(result.Metafile)
	}

	return nil
}
