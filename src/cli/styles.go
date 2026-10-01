package cli

// Styles defines color schemes for different UI elements
// Note: All styling is now handled by Logger in logger.go
// This file is kept for backward compatibility
type Styles struct {
	Title      interface{}
	Subtitle   interface{}
	Section    interface{}
	Key        interface{}
	Value      interface{}
	Path       interface{}
	Highlight  interface{}
	Dim        interface{}
	Stats      interface{}
	Badge      interface{}
	Warn       interface{}
	Error      interface{}
}

// DefaultStyles - kept for backward compatibility
// Use Logger methods instead for new code
var DefaultStyles = Styles{}
