package config

// Merge merges configuration values from override into base
func Merge(base, override *Config) {
	if override.Input != "" {
		base.Input = override.Input
	}
	if override.Output != "" {
		base.Output = override.Output
	}
	if override.SourceMap != "" {
		base.SourceMap = override.SourceMap
	}
	if override.Format != "" {
		base.Format = override.Format
	}
	if override.LogLevel != "" {
		base.LogLevel = override.LogLevel
	}
	if override.Minify {
		base.Minify = true
	}
	if override.Report {
		base.Report = true
	}
	if override.Watch {
		base.Watch = true
	}
	// Merge force flags
	if override.Force {
		base.Force = true
	}
	if override.Yes {
		base.Yes = true
	}
	if override.NoConfirm {
		base.NoConfirm = true
	}
}
