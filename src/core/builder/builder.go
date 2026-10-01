package builder

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/kalokaradia/jspackr/src/cli"
)

// Options defines build options
type Options struct {
	Input     string
	Output    string
	Minify    bool
	Report    bool
	SourceMap string
	Format    string
	Logger    *cli.Logger
}

// Run execute the build process with given options
func Run(opts Options) error {
	if opts.Input == "" {
		return errors.New("input file is required")
	}

	// Validate input path exists
	info, err := os.Stat(opts.Input)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("input path does not exist: " + opts.Input)
		}
		return err
	}
	if info.IsDir() {
		entries, err := os.ReadDir(opts.Input)
		if err != nil {
			return err
		}
		if len(entries) == 0 {
			return errors.New("input directory is empty: " + opts.Input)
		}
	}

	// make sure output directory exists
	if dir := filepath.Dir(opts.Output); dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	start := time.Now()

	// execute build
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
		Sourcemap:         MapSourceMap(opts.SourceMap),
		Format:            MapFormat(opts.Format),
	})

	if len(result.Errors) > 0 {
		return errors.New(result.Errors[0].Text)
	}

	elapsed := time.Since(start)

	// Build report
	buildResult := BuildResult{
		OutputPath:  opts.Output,
		InputSize:   GetInputSize(result.Metafile),
		OutputSize:  GetOutputSize(opts.Output),
		ModuleCount: GetModuleCount(result.Metafile),
		Elapsed:     elapsed,
		Metafile:    result.Metafile,
	}

	// Print report using logger if available
	if opts.Logger != nil && opts.Report {
		PrintReport(opts.Logger, buildResult)
	} else if opts.Report {
		// Fallback: print report without logger
		PrintReport(nil, buildResult)
	}

	return nil
}
