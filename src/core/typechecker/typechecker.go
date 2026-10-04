package typechecker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Check runs the project's local TypeScript compiler without emitting JavaScript.
func Check(projectDir string) error {
	absoluteDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("resolve project directory %q: %w", projectDir, err)
	}

	compiler, err := findCompiler(absoluteDir)
	if err != nil {
		return err
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" && filepath.Ext(compiler) == ".cmd" {
		cmd = exec.Command("cmd.exe", "/d", "/c", compiler, "--noEmit")
	} else {
		cmd = exec.Command(compiler, "--noEmit")
	}
	cmd.Dir = absoluteDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("TypeScript compiler: %w", err)
	}
	return nil
}

func findCompiler(projectDir string) (string, error) {
	binDir := filepath.Join(projectDir, "node_modules", ".bin")
	names := []string{"tsc"}
	if runtime.GOOS == "windows" {
		names = []string{"tsc.cmd", "tsc"}
	}

	for _, name := range names {
		path := filepath.Join(binDir, name)
		info, err := os.Stat(path)
		if err == nil && !info.IsDir() {
			return path, nil
		}
		if err != nil && !os.IsNotExist(err) {
			return "", fmt.Errorf("inspect local TypeScript compiler %q: %w", path, err)
		}
	}

	return "", fmt.Errorf(
		"local TypeScript compiler not found in %q; install TypeScript in this project (for example, npm install --save-dev typescript)",
		binDir,
	)
}
