# jspackr

<img src='./logo.svg' width=150>

`jspackr` is a small JavaScript bundler CLI implemented in Go and powered by [esbuild](https://esbuild.github.io/). It bundles one entry file and can minify output, generate source maps, print a size report, and rebuild when files under the entry file's directory change.

## Features

- Bundle JavaScript with esbuild into ESM, IIFE, or CommonJS output.
- ESM output supports top-level `await`.
- Optional whitespace, syntax, and identifier minification.
- Linked or inline source maps.
- Build size report and watch mode.
- Optional TypeScript type checking with the project's local `tsc`.
- JSON configuration plus command line overrides.

## Requirements and supported platforms

The Go source uses Go 1.22 or newer. Go can compile the CLI for Windows, Linux, and macOS on amd64 and arm64. The npm package currently contains a Linux x64 binary only; its launcher reports a clear error on other platforms. For other targets, install from source with Go.

## Installation

Install with Go:

```sh
go install github.com/kalokaradia/jspackr/src/main@latest
```

The resulting `jspackr` executable must be on your `PATH`. The npm package is also available for Linux x64:

```sh
npm install -g jspackr
```

## Quick start

```sh
jspackr --input src/index.js --out dist/bundle.js
jspackr -i src/index.js -o dist/app.js --minify --report
jspackr -i src/index.js -o dist/app.js --watch
jspackr -i src/index.js -o dist/app.mjs --format esm
jspackr -i src/index.ts -o dist/app.js --type-check
```

The output directory is created when needed. Existing output files prompt before overwrite; use `--force`/`--yes` to skip that prompt. `--no-confirm` skips all confirmation prompts.
If no `--config`/`-c` path is provided, jspackr loads `jspackr.config.json` from the current directory when that file exists.

## CLI options

| Option               | Meaning                                                 |
| -------------------- | ------------------------------------------------------- |
| `-i`, `--input`      | Entry JavaScript file (required)                        |
| `-o`, `--out`        | Output file; defaults to `dist/bundle.js`               |
| `-c`, `--config`     | Read a JSON configuration file                          |
| `-m`, `--minify`     | Enable esbuild minification                             |
| `-s`, `--source`     | Source map mode: `none`, `l` (linked), or `in` (inline) |
| `--format`           | Output format: `esm`, `iife`, or `cjs` (default: `iife`) |
| `-r`, `--report`     | Print bundle size information                           |
| `-w`, `--watch`      | Rebuild after changes under the entry file's directory  |
| `--type-check`       | Run the local TypeScript compiler with `--noEmit` before building |
| `--log-level`        | `debug`, `info`, `warn`, or `error`                     |
| `-f`, `--force`      | Skip output overwrite confirmation                      |
| `-y`, `--yes`        | Automatically confirm overwrite prompts                 |
| `-n`, `--no-confirm` | Skip all confirmation prompts                           |
| `-h`, `--help`       | Show help                                               |
| `-v`, `--version`    | Show CLI version                                        |

Example `jspackr.config.json`:

```json
{
  "input": "./src/index.js",
  "output": "./dist/bundle.js",
  "minify": true,
  "report": true,
  "typeCheck": true,
  "sourcemap": "l",
  "format": "esm",
  "watch": false,
  "logLevel": "info"
}
```

CLI values override config values when supplied. Boolean CLI flags turn a config option on; they do not turn a config option off.

Use `--format esm` to preserve ES module output and bundle entry points that use top-level `await`. The default `iife` keeps the classic self-executing bundle behavior. `cjs` emits CommonJS. When running an ESM output file with Node.js, use an `.mjs` extension or configure the containing package with `"type": "module"`.

Configuration supports the same build and confirmation options: `input`, `output`, `minify`, `report`, `typeCheck`, `sourcemap`, `format`, `watch`, `logLevel`, `force`, `yes`, and `noConfirm`. Set `"typeCheck": true` to run the local TypeScript compiler with `--noEmit` before building; this requires TypeScript installed in the project and a `tsconfig.json` selecting files. Source map values are `none`, `l` (linked), and `in` (inline); log levels are `debug`, `info`, `warn`, and `error`.

## Type checking

esbuild transforms and bundles JavaScript, JSX, TypeScript, and TSX, but it does not check TypeScript types. Type checking is an independent, opt-in step:

```sh
npm install --save-dev typescript
jspackr -i src/index.ts -o dist/app.js --type-check
```

`--type-check` runs the project's local `node_modules/.bin/tsc` (or `tsc.cmd` on Windows) with `--noEmit` before esbuild. The JavaScript bundle is still produced by esbuild. If type checking fails, jspackr exits unsuccessfully and does not bundle. Without `--type-check` and `"typeCheck": true`, existing build behavior is unchanged.

Enable the same check through `jspackr.config.json` instead of the CLI flag:

```json
{
  "input": "src/index.ts",
  "output": "dist/app.js",
  "typeCheck": true
}
```

The CLI flag can turn checking on, but omitting it does not disable `typeCheck: true` from the config. In watch mode, type checking runs before each rebuild.

When a `tsconfig.json` is present, `tsc` reads it normally, including its file selection, `compilerOptions`, JSX settings, and JavaScript options. JavaScript is checked only when the TypeScript project configuration enables it (typically `allowJs` and `checkJs`). JSX/TSX type checking follows the same `tsconfig.json`; esbuild continues to perform the actual transformation and bundling.

If there is no `tsconfig.json`, jspackr invokes `tsc --noEmit` without inventing a project or passing a source-file list. In keeping with TypeScript CLI behavior, `tsc` displays its help and exits unsuccessfully because no project or input files were specified. Add a `tsconfig.json` to select the project files and configure checking.

For example, with this file selected by `tsconfig.json`, checking succeeds:

```ts
const name: string = "Kaloka";
const age: number = 15;
export {};
```

This assignment is a type error; jspackr prints the original TypeScript diagnostic and does not start the esbuild bundle:

```ts
const age: number = "hello";
export {};
```

The `typescript` package is a dependency of the project being built, not a runtime dependency of jspackr.

## Development

The implementation is organized into `src/main` (CLI orchestration), `src/config` (JSON configuration and validation), `src/core/builder` (esbuild integration), `src/core/watcher` (file watching), `src/cli` (terminal output), and `src/utils` (shared helpers).

```sh
npm install
go mod download
go build ./...
go test ./...
go vet ./...
go test -race ./...
```

## Build and release

Build locally with `go build -o jspackr ./src/main`. Set the version in release binaries with:

```sh
go build -ldflags "-X main.version=0.4.0" -o jspackr ./src/main
```

For cross compilation, set `GOOS` and `GOARCH`, for example `GOOS=windows GOARCH=amd64`, `GOOS=linux GOARCH=arm64`, or `GOOS=darwin GOARCH=arm64`. Release maintainers should keep the git tag, npm `version`, and injected CLI version aligned. npm binary artifacts are currently maintained for Linux x64 only.

## Troubleshooting

- `input path does not exist`: check the entry file path relative to the current working directory.
- npm reports an unsupported platform: use `go install` on Windows, macOS, or non-x64 Linux; those prebuilt binaries are not currently shipped in npm.
- Output overwrite prompt: pass `--force` in scripts that intentionally replace an existing bundle.

## Contributing

Please use `gofmt` for Go files, add focused tests for behavior changes, and run the development checks above. Keep the CLI and configuration compatible unless a change is explicitly documented.

## License

MIT. See [LICENSE](LICENSE).
