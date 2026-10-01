package builder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/kalokaradia/jspackr/src/cli"
)

// MetaFile represents the structure of the metadata file
type MetaFile struct {
	Inputs  map[string]struct{ Bytes int `json:"bytes"` } `json:"inputs"`
	Outputs map[string]struct{ Bytes int `json:"bytes"` } `json:"outputs"`
}

// BuildResult holds the build information for reporting
type BuildResult struct {
	OutputPath  string
	InputSize   int64
	OutputSize  int64
	ModuleCount int
	Elapsed     time.Duration
	Metafile    string
}

// formatBytes formats bytes to human readable string
func formatBytes(bytes int64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	}
	if bytes < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(bytes)/1024)
	}
	return fmt.Sprintf("%.1f MB", float64(bytes)/(1024*1024))
}

// PrintReport prints a detailed build report using CLI styles
func PrintReport(logger *cli.Logger, result BuildResult) {
	// Output path (relative to current directory)
	relPath, _ := filepath.Rel(".", result.OutputPath)

	fmt.Println()

	// Success header
	if logger != nil {
		logger.Success("Build succeeded")
	} else {
		fmt.Println("✓ Build succeeded")
	}

	// Output
	if logger != nil {
		logger.PrintKeyValue("Output", relPath, 1)
	} else {
		fmt.Printf("  Output: %s\n", relPath)
	}

	// Size comparison
	inputSizeStr := formatBytes(result.InputSize)
	outputSizeStr := formatBytes(result.OutputSize)
	percent := float64(result.OutputSize) / float64(result.InputSize) * 100
	sizeStr := fmt.Sprintf("%s → %s (%.0f%%)", inputSizeStr, outputSizeStr, percent)
	if logger != nil {
		logger.PrintStat("Size:", sizeStr)
	} else {
		fmt.Printf("  Size: %s\n", sizeStr)
	}

	// Module count
	modulesStr := fmt.Sprintf("%d", result.ModuleCount)
	if result.ModuleCount == 1 {
		modulesStr += " module"
	} else {
		modulesStr += " modules"
	}
	if logger != nil {
		logger.PrintStat("Modules:", modulesStr)
	} else {
		fmt.Printf("  Modules: %s\n", modulesStr)
	}

	// Build time
	timeStr := fmt.Sprintf("%dms", result.Elapsed.Milliseconds())
	if logger != nil {
		logger.PrintStat("Time:", timeStr)
	} else {
		fmt.Printf("  Time: %s\n", timeStr)
	}

	// Print detailed contributors if report flag is enabled
	if result.Metafile != "" {
		contributors := getContributors(result.Metafile)
		if len(contributors) > 0 {
			fmt.Println()
			if logger != nil {
				logger.PrintSection("Top contributors")
			} else {
				fmt.Println("Top contributors:")
			}
			for i := 0; i < len(contributors) && i < 5; i++ {
				item := contributors[i]
				sizeKB := float64(item.Bytes) / 1024
				sizeStr := fmt.Sprintf("%.1f KB", sizeKB)
				if logger != nil {
					logger.GetColors().Printf("  %-50s ", item.Path)
					logger.GetSuccessColor().Printf("%10s\n", sizeStr)
				} else {
					fmt.Printf("  %-50s %10s\n", item.Path, sizeStr)
				}
			}
		}
	}

	fmt.Println()
}

// contributorItem represents a single contributor item
type contributorItem struct {
	Path  string
	Bytes int
}

// getContributors extracts and sorts contributors from metadata
func getContributors(meta string) []contributorItem {
	var m MetaFile
	_ = json.Unmarshal([]byte(meta), &m)

	items := make([]contributorItem, 0, len(m.Inputs))
	for path, v := range m.Inputs {
		items = append(items, contributorItem{Path: path, Bytes: v.Bytes})
	}

	if len(items) == 0 {
		return items
	}

	// sort by size descending
	sort.Slice(items, func(i, j int) bool {
		return items[i].Bytes > items[j].Bytes
	})

	return items
}

// GetInputSize returns the total size of input files
func GetInputSize(meta string) int64 {
	var m MetaFile
	if err := json.Unmarshal([]byte(meta), &m); err != nil {
		return 0
	}

	var total int64
	for _, v := range m.Inputs {
		total += int64(v.Bytes)
	}
	return total
}

// GetOutputSize returns the size of the output file
func GetOutputSize(outputPath string) int64 {
	info, err := os.Stat(outputPath)
	if err != nil {
		return 0
	}
	return info.Size()
}

// GetModuleCount returns the number of modules from metadata
func GetModuleCount(meta string) int {
	var m MetaFile
	if err := json.Unmarshal([]byte(meta), &m); err != nil {
		return 0
	}
	return len(m.Inputs)
}
