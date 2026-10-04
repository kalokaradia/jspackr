---
title: CLI reference
description: Complete jspackr command line options, aliases, defaults, and interactions.
---

## Usage

```sh
jspackr [options]
```

`--help` prints the built-in usage text; `--version` prints the version. The CLI uses Go's standard flag parser, so values can be given as `--input file.js` or `--input=file.js`.

## Build options

| Option | Value / default | Description |
| --- | --- | --- |
| `-i`, `--input` | file; required | Entry file. Paths are resolved from the current working directory. |
| `-o`, `--out` | `dist/bundle.js` | Output file. Parent directories are created as needed. |
| `-c`, `--config` | auto-detect `jspackr.config.json` | Load a JSON config file. |
| `-m`, `--minify` | off | Enable whitespace, syntax, and identifier minification. |
| `-s`, `--source` | `none`; `none`, `l`, `in` | Source map mode: none, linked, or inline. |
| `--format` | `iife`; `iife`, `esm`, `cjs` | Output format. |
| `-r`, `--report` | off | Print sizes, module count, elapsed time, and up to five largest input contributors. |
| `-w`, `--watch` | off | Rebuild on changes under the entry directory tree. |
| `--type-check` | off | Enable TypeScript type checking. Also configurable as `"typeCheck": true`. Requires local TypeScript and runs `tsc --noEmit` before building. |
| `--log-level` | `info`; `debug`, `info`, `warn`, `error` | Set logger threshold. |

Example:

```sh
jspackr -i src/index.js -o dist/app.js --minify --source l --report --format esm
```

## Confirmation options

| Option | Description |
| --- | --- |
| `-f`, `--force` | Skip confirmation before overwriting an existing output file. |
| `-y`, `--yes` | Automatically accept overwrite prompts. |
| `-n`, `--no-confirm` | Skip all confirmation prompts, including creating a missing output directory. |

`force`, `yes`, and `no-confirm` all allow overwriting existing files. Only `no-confirm` also skips the prompt to create a missing output directory. The options are accepted in config as `force`, `yes`, and `noConfirm`.

## Information options

| Option | Description |
| --- | --- |
| `-h`, `--help` | Print help and exit. |
| `-v`, `--version` | Print version and exit. The implementation rejects combining `--version` with additional arguments. |

## Configuration interaction

`--type-check` turns type checking on from the CLI. It can also be enabled in `jspackr.config.json` with `"typeCheck": true`. In watch mode it checks before each build. See the [TypeScript and JSX guide](/guides/typescript-jsx/) for the `tsconfig.json` and local TypeScript behavior.

If no config path is supplied, jspackr checks the current directory for `jspackr.config.json`. Explicit `--config` selects a file. Non-empty string CLI values override config values. Boolean CLI flags only set `true`; they cannot turn a config boolean off. See [configuration](/guides/configuration/) and the [config reference](/reference/configuration/).
