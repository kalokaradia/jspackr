package cli

import (
	"fmt"

	"github.com/kalokaradia/jspackr/src/config"
)

// ProgressBar represents a simple progress bar
type ProgressBar struct {
	total     int
	width     int
	prefix    string
	fillChar  string
	emptyChar string
	fillColor interface{}
}

// NewProgressBar creates a new progress bar
func NewProgressBar(total int, prefix string) *ProgressBar {
	return &ProgressBar{
		total:     total,
		width:     40,
		prefix:    prefix,
		fillChar:  "█",
		emptyChar: "░",
	}
}

// Render draws the progress bar at current progress
func (p *ProgressBar) Render(current int) {
	if current > p.total {
		current = p.total
	}

	percent := float64(current) / float64(p.total)
	filled := int(float64(p.width) * percent)
	empty := p.width - filled

	bar := fmt.Sprintf("%s%s",
		fmt.Sprintf("%.*s", filled, p.fillChar),
		fmt.Sprintf("%.*s", empty, p.emptyChar))

	fmt.Printf("\r%s [%s] %3d%%", p.prefix, bar, int(percent*100))
}

// Finish completes the progress bar
func (p *ProgressBar) Finish(message string) {
	p.Render(p.total)
	fmt.Println()
	if message != "" {
		fmt.Printf("✓ %s\n", message)
	}
}

// PrintStatus prints a status message with consistent formatting
func PrintStatus(icon, message string) {
	fmt.Printf("%s %s\n", icon, message)
}

// PrintBuildSummary prints a summary of the build configuration
// Use Logger.PrintSection and Logger.PrintKeyValue instead
func PrintBuildSummary(cfg *config.Config) {
	// Kept for backward compatibility
	// New code should use Logger methods
}

// PrintBuildResult prints the result of a build operation
func PrintBuildResult(success bool, message string) {
	if success {
		fmt.Printf("✓ %s\n", message)
	} else {
		fmt.Printf("✗ %s\n", message)
	}
}

// PrintHelpInfo prints help/tip information
func PrintHelpInfo(tip string) {
	if tip == "" {
		return
	}
	fmt.Printf("  💡 %s\n", tip)
}

// NewLine prints a new line
func NewLine() {
	fmt.Println()
}
