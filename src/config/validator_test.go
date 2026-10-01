package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateSourceMap(t *testing.T) {
	tests := []struct {
		name string
		mode string
		wantErr bool
	}{
		{name: "none", mode: "none"},
		{name: "linked", mode: "l"},
		{name: "inline", mode: "in"},
		{name: "unknown", mode: "external", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Default()
			cfg.Input = "entry.js"
			cfg.SourceMap = tt.mode
			err := Validate(cfg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLogLevel(t *testing.T) {
	for _, level := range []string{"debug", "info", "warn", "error"} {
		t.Run(level, func(t *testing.T) {
			cfg := Default()
			cfg.Input = "entry.js"
			cfg.LogLevel = level
			if err := Validate(cfg); err != nil {
				t.Fatalf("Validate() error = %v", err)
			}
		})
	}

	cfg := Default()
	cfg.Input = "entry.js"
	cfg.LogLevel = "verbose"
	if err := Validate(cfg); err == nil {
		t.Fatal("Validate() accepted an unknown log level")
	}
}

func TestValidateInputPath(t *testing.T) {
	dir := t.TempDir()
	if err := ValidateInputPath(dir); err == nil {
		t.Fatal("ValidateInputPath() accepted an empty directory")
	}

	entry := filepath.Join(dir, "entry.js")
	if err := os.WriteFile(entry, []byte("console.log('ok')"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := ValidateInputPath(entry); err != nil {
		t.Fatalf("ValidateInputPath() rejected a file: %v", err)
	}
	if err := ValidateInputPath(filepath.Join(dir, "missing.js")); err == nil {
		t.Fatal("ValidateInputPath() accepted a missing path")
	}
}

func TestValidateOutputPath(t *testing.T) {
	dir := t.TempDir()
	out := filepath.Join(dir, "bundle.js")
	got, err := ValidateOutputPath(out)
	if err != nil {
		t.Fatalf("ValidateOutputPath() error = %v", err)
	}
	if got != dir {
		t.Fatalf("ValidateOutputPath() dir = %q, want %q", got, dir)
	}

	if err := os.WriteFile(out, []byte("file"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := ValidateOutputPath(filepath.Join(out, "nested.js")); err == nil {
		t.Fatal("ValidateOutputPath() accepted a file as a parent directory")
	}
}
