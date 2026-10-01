package main

import (
	"errors"
	"os"

	"github.com/kalokaradia/jspackr/src/cli"
	"github.com/kalokaradia/jspackr/src/config"
	"github.com/kalokaradia/jspackr/src/core/builder"
	"github.com/kalokaradia/jspackr/src/core/watcher"
	"github.com/kalokaradia/jspackr/src/utils"
)

// version is set by release builds with -ldflags. Local builds report dev.
var version = "0.4.0"

func main() {
	flagCfg, configPath, showVersion, help := utils.ParseFlags()

	// Handle version flag
	if err := utils.ValidateVersionFlag(showVersion); err != nil {
		// Create logger with default level for error output
		logger := cli.New("info")
		logger.Fatal(err.Error())
	}

	if showVersion {
		utils.ShowVersion(version)
		return
	}

	if help {
		utils.ShowUsage(version)
		return
	}

	// Load configuration
	finalCfg := config.Default()

	if configPath == "" {
		defaultConfig, _ := utils.FindConfigFile()
		if defaultConfig != "" {
			configPath = defaultConfig
		}
	}

	if configPath != "" {
		fileCfg, err := config.Load(configPath)
		if err != nil {
			logger := cli.New("info")
			logger.Fatal("Failed to load config: " + err.Error())
		}
		finalCfg = fileCfg
	}

	config.Merge(finalCfg, flagCfg)

	if err := config.Validate(finalCfg); err != nil {
		logger := cli.New("info")
		logger.Fatal(err.Error())
	}

	logger := cli.New(finalCfg.LogLevel)

	// Print welcome banner
	logger.PrintTitle()

	// Print full build configuration summary
	printBuildSummary(logger, finalCfg)

	// Validate input path exists
	if err := config.ValidateInputPath(finalCfg.Input); err != nil {
		logger.FatalErr(err, "Invalid input path")
	}

	// Validate output path and handle directory creation
	outDir := utils.GetOutputParent(finalCfg.Output)
	if outDir != "." {
		if _, err := config.ValidateOutputPath(finalCfg.Output); err != nil {
			// Output directory doesn't exist, ask user to create it
			// Skip confirmation if noConfirm flag is set
			if !finalCfg.NoConfirm {
				if !logger.ConfirmCreateDir(outDir) {
					logger.Warn("Build cancelled")
					return
				}
			}
			if err := utils.CreateDir(outDir); err != nil {
				logger.FatalErr(err, "Failed to create directory")
			}
			logger.PrintDirCreated(outDir)
		}
	}

	if err := utils.ValidateOutputFile(finalCfg.Output); err != nil {
		logger.FatalErr(err, "Invalid output path")
	}

	// Check if we should overwrite existing file
	// Skip confirmation if force, yes, or noConfirm flags are set
	confirmed, err := confirmOverwrite(logger, finalCfg.Output, finalCfg.Force, finalCfg.Yes, finalCfg.NoConfirm)
	if err != nil {
		logger.FatalErr(err, "Cannot inspect output path")
	}
	if !confirmed {
		logger.Warn("Build cancelled")
		return
	}

	if outDir != "." {
		if notEmpty, _ := utils.DirNotEmpty(outDir); notEmpty {
			logger.WarnWithTip(
				"Output directory not empty: "+outDir,
				"Existing files may be overwritten",
			)
		}
	}

	opts := builder.Options{
		Input:     finalCfg.Input,
		Output:    finalCfg.Output,
		Minify:    finalCfg.Minify,
		Report:    finalCfg.Report,
		SourceMap: finalCfg.SourceMap,
		Format:    finalCfg.Format,
		Logger:    logger,
	}

	if finalCfg.Watch {
		logger.Info("Watch mode enabled")
		if err := watcher.WatchFiles(finalCfg.Input, opts, logger); err != nil {
			logger.FatalErr(err, "Watch failed")
		}
		return
	}

	// Start build
	logger.PrintBuildStart()
	spinner := logger.NewSpinner("Bundling...")
	spinner.Start()

	if err := builder.Run(opts); err != nil {
		spinner.Stop(false)
		logger.FatalErr(err, "Build failed")
	}

	spinner.Stop(true)

	logger.PrintSuccess()
}

// printBuildSummary prints a summary of the build configuration
func printBuildSummary(logger *cli.Logger, cfg *config.Config) {
	logger.PrintSection("Build Configuration")

	logger.PrintKeyValue("Input", cfg.Input, 0)
	logger.PrintKeyValue("Output", cfg.Output, 0)
	logger.PrintKeyValue("Minify", boolToStr(cfg.Minify), 0)
	logger.PrintKeyValue("Source Map", cfg.SourceMap, 0)
	logger.PrintKeyValue("Format", cfg.Format, 0)
	logger.PrintKeyValue("Report", boolToStr(cfg.Report), 0)
	logger.PrintKeyValue("Watch Mode", boolToStr(cfg.Watch), 0)
	logger.PrintKeyValue("Log Level", cfg.LogLevel, 0)

	logger.PrintDivider()
}

// confirmOverwrite checks if we should overwrite existing file
func confirmOverwrite(logger *cli.Logger, output string, force, yes, noConfirm bool) (bool, error) {
	if force || yes || noConfirm {
		return true, nil
	}

	// Check if file exists
	if _, err := os.Stat(output); errors.Is(err, os.ErrNotExist) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return logger.Confirm("File already exists. Overwrite?", false), nil
}

// boolToStr converts bool to string
func boolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
