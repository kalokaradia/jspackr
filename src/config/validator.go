package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Validate validates the configuration
func Validate(cfg *Config) error {
	if cfg.Input == "" {
		return errors.New("entry file is required")
	}
	if cfg.LogLevel != "" {
		switch cfg.LogLevel {
		case "debug", "info", "warn", "error":
		default:
			return fmt.Errorf("invalid log level %q: use debug, info, warn, or error", cfg.LogLevel)
		}
	}
	switch cfg.Format {
	case "iife", "esm", "cjs":
	default:
		return errors.New("invalid format: use iife, esm, or cjs")
	}

	switch cfg.SourceMap {
	case "none", "l", "in":
		return nil
	default:
		return errors.New("invalid sourcemap mode: use none, l, or in")
	}
}

// ValidateInputPath checks if the input path exists
func ValidateInputPath(input string) error {
	info, err := os.Stat(input)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("input path %q does not exist: %w", input, err)
		}
		return fmt.Errorf("inspect input path %q: %w", input, err)
	}
	if info.IsDir() {
		// For directory input, check if it's empty
		entries, err := os.ReadDir(input)
		if err != nil {
			return fmt.Errorf("read input directory %q: %w", input, err)
		}
		if len(entries) == 0 {
			return fmt.Errorf("input directory is empty: %s", input)
		}
	}
	return nil
}

// ValidateOutputPath checks if the output parent directory exists
// Returns the parent directory path and an error if parent doesn't exist
func ValidateOutputPath(output string) (string, error) {
	dir := filepath.Dir(output)
	if dir == "." {
		// Output is in current directory, always valid
		return dir, nil
	}
	info, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return dir, fmt.Errorf("output directory does not exist: %s", dir)
		}
		return dir, fmt.Errorf("inspect output directory %q: %w", dir, err)
	}
	if !info.IsDir() {
		return dir, errors.New("output path is not a directory: " + dir)
	}
	return dir, nil
}
