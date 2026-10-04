package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	binaryPath  string
	binaryDir   string
	projectRoot string
	buildOnce   sync.Once
	buildErr    error
)

func TestMain(m *testing.M) {
	_, source, _, _ := runtime.Caller(0)
	projectRoot = filepath.Clean(filepath.Join(filepath.Dir(source), "..", ".."))
	code := m.Run()
	if binaryDir != "" {
		_ = os.RemoveAll(binaryDir)
	}
	os.Exit(code)
}

func cliBinary(t *testing.T) string {
	t.Helper()
	buildOnce.Do(func() {
		tempDir, err := os.MkdirTemp("", "jspackr-e2e-")
		if err != nil {
			buildErr = err
			return
		}
		binaryDir = tempDir
		binaryPath = filepath.Join(tempDir, "jspackr")
		if runtime.GOOS == "windows" {
			binaryPath += ".exe"
		}
		cmd := exec.Command("go", "build", "-o", binaryPath, "./src/main")
		cmd.Dir = projectRoot
		if output, err := cmd.CombinedOutput(); err != nil {
			buildErr = fmt.Errorf("build jspackr: %w\n%s", err, output)
		}
	})
	if buildErr != nil {
		t.Fatal(buildErr)
	}
	return binaryPath
}

func newProject(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "src"), 0o755); err != nil {
		t.Fatal(err)
	}
	return dir
}

func installLocalTSC(t *testing.T, project string) {
	t.Helper()
	binDir := filepath.Join(project, "node_modules", ".bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	tscJS := filepath.Join(projectRoot, "node_modules", "typescript", "bin", "tsc")
	var path, contents string
	if runtime.GOOS == "windows" {
		path = filepath.Join(binDir, "tsc.cmd")
		contents = "@echo off\r\nnode \"" + tscJS + "\" %*\r\n"
	} else {
		path = filepath.Join(binDir, "tsc")
		contents = "#!/bin/sh\nexec node \"" + tscJS + "\" \"$@\"\n"
	}
	if err := os.WriteFile(path, []byte(contents), 0o755); err != nil {
		t.Fatal(err)
	}
}

func writeFile(t *testing.T, path, contents string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
}

func runCLI(t *testing.T, dir string, args ...string) (int, string) {
	t.Helper()
	cmd := exec.Command(cliBinary(t), args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err == nil {
		return 0, string(output)
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode(), string(output)
	}
	t.Fatalf("run jspackr: %v\n%s", err, output)
	return 0, ""
}

func addTSConfig(t *testing.T, project, config string) {
	t.Helper()
	writeFile(t, filepath.Join(project, "tsconfig.json"), config)
}

func TestTypeCheckValidTypeScript(t *testing.T) {
	project := newProject(t)
	installLocalTSC(t, project)
	writeFile(t, filepath.Join(project, "src", "index.ts"), `const name: string = "Kaloka"; const age: number = 15; export {};`)
	addTSConfig(t, project, `{"files":["src/index.ts"]}`)

	code, output := runCLI(t, project, "-i", "src/index.ts", "-o", "dist/app.js", "--no-confirm", "--type-check")
	if code != 0 {
		t.Fatalf("jspackr exit = %d, want 0:\n%s", code, output)
	}
	if _, err := os.Stat(filepath.Join(project, "dist", "app.js")); err != nil {
		t.Fatalf("bundle was not produced: %v", err)
	}
	if _, err := os.Stat(filepath.Join(project, "src", "index.js")); !os.IsNotExist(err) {
		t.Fatalf("tsc emitted JavaScript despite --noEmit (stat error: %v)", err)
	}
}

func TestTypeCheckInvalidTypeScriptFailsBeforeBundling(t *testing.T) {
	project := newProject(t)
	installLocalTSC(t, project)
	writeFile(t, filepath.Join(project, "src", "index.ts"), `const age: number = "hello";`)
	addTSConfig(t, project, `{"files":["src/index.ts"]}`)

	code, output := runCLI(t, project, "-i", "src/index.ts", "-o", "dist/app.js", "--no-confirm", "--type-check")
	if code == 0 {
		t.Fatalf("jspackr exit = 0, want failure:\n%s", output)
	}
	if !strings.Contains(output, "TS2322") || !strings.Contains(output, "src/index.ts") {
		t.Fatalf("compiler diagnostics were not preserved:\n%s", output)
	}
	if _, err := os.Stat(filepath.Join(project, "dist", "app.js")); !os.IsNotExist(err) {
		t.Fatalf("bundle exists after failed type check (stat error: %v)", err)
	}
}

func TestTypeCheckRequiresLocalCompiler(t *testing.T) {
	project := newProject(t)
	writeFile(t, filepath.Join(project, "src", "index.ts"), `const value = 1;`)
	addTSConfig(t, project, `{"files":["src/index.ts"]}`)

	code, output := runCLI(t, project, "-i", "src/index.ts", "-o", "dist/app.js", "--no-confirm", "--type-check")
	if code == 0 {
		t.Fatalf("jspackr exit = 0, want failure:\n%s", output)
	}
	if !strings.Contains(output, "local TypeScript compiler not found") ||
		!strings.Contains(output, "npm install --save-dev typescript") {
		t.Fatalf("missing actionable compiler error:\n%s", output)
	}
}

func TestTypeCheckUsesTypeScriptDefaultWithoutTsconfig(t *testing.T) {
	project := newProject(t)
	installLocalTSC(t, project)
	writeFile(t, filepath.Join(project, "src", "index.ts"), `const value: number = 1;`)

	code, output := runCLI(t, project, "-i", "src/index.ts", "-o", "dist/app.js", "--no-confirm", "--type-check")
	if code == 0 {
		t.Fatalf("tsc without a tsconfig or explicit files unexpectedly succeeded:\n%s", output)
	}
	if !strings.Contains(output, "tsc") || !strings.Contains(strings.ToLower(output), "help") {
		t.Fatalf("expected TypeScript CLI's no-project help output:\n%s", output)
	}
}

func TestTypeCheckRespectsTsconfigFileSelection(t *testing.T) {
	project := newProject(t)
	installLocalTSC(t, project)
	writeFile(t, filepath.Join(project, "src", "index.ts"), `const value: string = "valid";`)
	writeFile(t, filepath.Join(project, "src", "excluded.ts"), `const value: number = "invalid";`)
	addTSConfig(t, project, `{"compilerOptions":{"strict":true},"include":["src/index.ts"],"exclude":["src/excluded.ts"]}`)

	code, output := runCLI(t, project, "-i", "src/index.ts", "-o", "dist/app.js", "--no-confirm", "--type-check")
	if code != 0 {
		t.Fatalf("jspackr did not honor tsconfig include/exclude: exit %d:\n%s", code, output)
	}
	if _, err := os.Stat(filepath.Join(project, "dist", "app.js")); err != nil {
		t.Fatalf("bundle was not produced: %v", err)
	}
}

func TestJavaScriptCheckingFollowsTsconfig(t *testing.T) {
	for _, checkJS := range []bool{false, true} {
		t.Run(fmt.Sprintf("checkJs_%t", checkJS), func(t *testing.T) {
			project := newProject(t)
			installLocalTSC(t, project)
			writeFile(t, filepath.Join(project, "src", "index.js"), `/** @type {string} */ const value = 1;`)
			addTSConfig(t, project, fmt.Sprintf(
				`{"compilerOptions":{"allowJs":true,"checkJs":%t},"include":["src/**/*.js"]}`,
				checkJS,
			))

			code, output := runCLI(t, project, "-i", "src/index.js", "-o", "dist/app.js", "--no-confirm", "--type-check")
			if checkJS && code == 0 {
				t.Fatalf("jspackr exit = 0 with checkJs enabled:\n%s", output)
			}
			if !checkJS && code != 0 {
				t.Fatalf("jspackr failed with checkJs disabled: exit %d:\n%s", code, output)
			}
		})
	}
}

func TestBuildWithoutTypeCheckRemainsUnchanged(t *testing.T) {
	project := newProject(t)
	writeFile(t, filepath.Join(project, "src", "index.ts"), `const age: number = "hello";`)

	code, output := runCLI(t, project, "-i", "src/index.ts", "-o", "dist/app.js", "--no-confirm")
	if code != 0 {
		t.Fatalf("build without --type-check failed: exit %d:\n%s", code, output)
	}
	if _, err := os.Stat(filepath.Join(project, "dist", "app.js")); err != nil {
		t.Fatalf("bundle was not produced: %v", err)
	}
}

func TestTypeCheckCanBeEnabledFromJspackrConfig(t *testing.T) {
	project := newProject(t)
	installLocalTSC(t, project)
	writeFile(t, filepath.Join(project, "src", "index.ts"), `const age: number = "hello";`)
	addTSConfig(t, project, `{"files":["src/index.ts"]}`)
	writeFile(t, filepath.Join(project, "jspackr.config.json"),
		`{"input":"src/index.ts","output":"dist/app.js","typeCheck":true}`)

	code, output := runCLI(t, project, "--no-confirm")
	if code == 0 {
		t.Fatalf("type checking from config unexpectedly succeeded:\n%s", output)
	}
	if !strings.Contains(output, "TS2322") {
		t.Fatalf("expected TypeScript diagnostic from config option:\n%s", output)
	}
	if _, err := os.Stat(filepath.Join(project, "dist", "app.js")); !os.IsNotExist(err) {
		t.Fatalf("bundle exists after type-check failed (stat error: %v)", err)
	}
}

func TestJavaScriptJSXAndTSXBuild(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		source   string
	}{
		{name: "javascript", filename: "index.js", source: `console.log("ok");`},
		{name: "jsx", filename: "index.jsx", source: `const App = () => <div>ok</div>;`},
		{name: "tsx", filename: "index.tsx", source: `const App = (): JSX.Element => <div>ok</div>;`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			project := newProject(t)
			writeFile(t, filepath.Join(project, "src", tt.filename), tt.source)

			code, output := runCLI(t, project, "-i", filepath.Join("src", tt.filename), "-o", "dist/app.js", "--no-confirm")
			if code != 0 {
				t.Fatalf("build failed: exit %d:\n%s", code, output)
			}
			if _, err := os.Stat(filepath.Join(project, "dist", "app.js")); err != nil {
				t.Fatalf("bundle was not produced: %v", err)
			}
		})
	}
}

func TestWatchTypeChecksBeforeTriggeredRebuild(t *testing.T) {
	project := newProject(t)
	installLocalTSC(t, project)
	input := filepath.Join(project, "src", "index.ts")
	writeFile(t, input, `const age: number = 15; export {};`)
	addTSConfig(t, project, `{"files":["src/index.ts"]}`)

	cmd := exec.Command(cliBinary(t), "-i", "src/index.ts", "-o", "dist/app.js", "--no-confirm", "--type-check", "--watch")
	cmd.Dir = project
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	waitDone := make(chan struct{})
	go func() {
		_ = cmd.Wait()
		close(waitDone)
	}()
	t.Cleanup(func() {
		select {
		case <-waitDone:
		default:
			_ = cmd.Process.Kill()
			<-waitDone
		}
	})

	lines := make(chan string, 32)
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()
	waitForOutput(t, lines, "Watching", 20*time.Second)

	writeFile(t, input, `const age: number = "invalid"; export {};`)
	waitForOutput(t, lines, "TS2322", 20*time.Second)

	select {
	case <-waitDone:
		t.Fatal("watch process exited after type error")
	case <-time.After(250 * time.Millisecond):
	}
	if _, err := os.Stat(filepath.Join(project, "dist", "app.js")); !os.IsNotExist(err) {
		t.Fatalf("watcher bundled after failed type check (stat error: %v)", err)
	}
}

func waitForOutput(t *testing.T, lines <-chan string, fragment string, timeout time.Duration) {
	t.Helper()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	for {
		select {
		case line, ok := <-lines:
			if !ok {
				t.Fatalf("process output ended before %q appeared", fragment)
			}
			if strings.Contains(line, fragment) {
				return
			}
		case <-timer.C:
			t.Fatalf("timed out waiting for %q in process output", fragment)
		}
	}
}
