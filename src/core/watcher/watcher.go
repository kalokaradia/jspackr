package watcher

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/kalokaradia/jspackr/src/cli"
	"github.com/kalokaradia/jspackr/src/core/builder"
)

const debounceDelay = 300 * time.Millisecond

// WatchFiles watches the entry tree and rebuilds until interrupted by Ctrl+C.
func WatchFiles(entry string, opts builder.Options, logger *cli.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	return WatchFilesContext(ctx, entry, opts, logger)
}

// WatchFilesContext watches the entry tree until ctx is cancelled.
func WatchFilesContext(ctx context.Context, entry string, opts builder.Options, logger *cli.Logger) error {
	if logger == nil {
		logger = cli.New("info")
	}

	entryPath, err := filepath.Abs(entry)
	if err != nil {
		return fmt.Errorf("resolve entry path %q: %w", entry, err)
	}
	entryInfo, err := os.Stat(entryPath)
	if err != nil {
		return fmt.Errorf("inspect entry path %q: %w", entryPath, err)
	}
	root := filepath.Dir(entryPath)
	if entryInfo.IsDir() {
		root = entryPath
	}

	outputPath, err := filepath.Abs(opts.Output)
	if err != nil {
		return fmt.Errorf("resolve output path %q: %w", opts.Output, err)
	}

	fsw, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create file watcher: %w", err)
	}
	defer fsw.Close()

	if err := addTree(fsw, root); err != nil {
		return err
	}
	logger.PrintWatch(entryPath)

	var timer *time.Timer
	var rebuild <-chan time.Time
	defer func() {
		if timer != nil {
			timer.Stop()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case event, ok := <-fsw.Events:
			if !ok {
				return nil
			}
			if event.Op&fsnotify.Create != 0 {
				if info, statErr := os.Stat(event.Name); statErr == nil && info.IsDir() {
					if err := addTree(fsw, event.Name); err != nil {
						logger.Error("Watcher error: %v", err)
					}
				}
			}
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove|fsnotify.Rename) == 0 || isGeneratedOutput(event.Name, outputPath) {
				continue
			}
			if timer == nil {
				timer = time.NewTimer(debounceDelay)
			} else {
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(debounceDelay)
			}
			rebuild = timer.C
		case err, ok := <-fsw.Errors:
			if !ok {
				return nil
			}
			logger.Error("Watcher error: %v", err)
		case <-rebuild:
			rebuild = nil
			logger.PrintRebuild()
			if err := builder.Run(opts); err != nil {
				logger.Error("Build failed: %v", err)
			} else {
				logger.PrintSuccess()
			}
		}
	}
}

func addTree(fsw *fsnotify.Watcher, root string) error {
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return fmt.Errorf("walk watched directory %q: %w", path, walkErr)
		}
		if entry.IsDir() && path != root && (entry.Name() == ".git" || entry.Name() == "node_modules") {
			return filepath.SkipDir
		}
		if !entry.IsDir() {
			return nil
		}
		if err := fsw.Add(path); err != nil {
			return fmt.Errorf("watch directory %q: %w", path, err)
		}
		return nil
	})
	return err
}

func isGeneratedOutput(path, output string) bool {
	absolute, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	clean := filepath.Clean(absolute)
	if runtime.GOOS == "windows" {
		clean = strings.ToLower(clean)
		output = strings.ToLower(filepath.Clean(output))
	} else {
		output = filepath.Clean(output)
	}
	return clean == output || clean == output+".map"
}
