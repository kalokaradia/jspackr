package config

// Config represents the jspackr configuration
type Config struct {
	Input     string `json:"input"`
	Output    string `json:"output"`
	Minify    bool   `json:"minify"`
	Report    bool   `json:"report"`
	SourceMap string `json:"sourcemap"`
	Format    string `json:"format"`
	Watch     bool   `json:"watch"`
	LogLevel  string `json:"logLevel"`
	TypeCheck bool   `json:"typeCheck"`
	// Force flags for non-interactive mode
	Force     bool `json:"force"`     // Skip overwrite confirmation
	Yes       bool `json:"yes"`       // Auto-confirm overwrite
	NoConfirm bool `json:"noConfirm"` // Skip all confirmations
}

// Default returns the default configuration
func Default() *Config {
	return &Config{
		Output:    "dist/bundle.js",
		SourceMap: "none",
		Format:    "iife",
		LogLevel:  "info",
	}
}
