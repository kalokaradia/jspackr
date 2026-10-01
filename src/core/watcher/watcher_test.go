package watcher

import (
	"path/filepath"
	"testing"
)

func TestIsGeneratedOutput(t *testing.T) {
	dir := t.TempDir()
	output := filepath.Join(dir, "bundle.js")
	tests := []struct {
		name string
		path string
		want bool
	}{
		{name: "bundle", path: output, want: true},
		{name: "source map", path: output + ".map", want: true},
		{name: "input", path: filepath.Join(dir, "index.js"), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGeneratedOutput(tt.path, output); got != tt.want {
				t.Fatalf("isGeneratedOutput(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
