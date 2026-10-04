---
title: Troubleshooting
description: Fix common installation, input, output, and build errors in jspackr.
---

## `input path ... does not exist`

Input paths are resolved from the current working directory. Run the command at the project root or supply a correct relative/absolute file path.

## `entry file is required`

Pass `--input` / `-i`, or set `input` in a JSON configuration file. A config file in the current directory is loaded automatically when present.

## Output overwrite prompt

Use `--force` for an intentional replacement in scripts. `--yes` also accepts overwrite prompts. `--no-confirm` skips both overwrite and missing-directory prompts.

## npm reports a missing binary

The published package artifacts currently include Linux x64/arm64 and macOS x64/arm64 binaries. Install from source with Go on Windows or another platform without a packaged binary.

## Type errors are not reported

Type checking is opt-in. Pass `--type-check` and install TypeScript in the project (`npm install --save-dev typescript`). jspackr looks only for the local compiler in `node_modules/.bin` (or `node_modules/.bin/tsc.cmd` on Windows), not a global installation.

The compiler uses `tsconfig.json` as-is. If no `tsconfig.json` exists, TypeScript's CLI displays its help and exits unsuccessfully because jspackr does not invent a file list or compiler configuration. Add a `tsconfig.json` selecting the project files.

When the compiler reports an error, jspackr exits unsuccessfully and does not start the esbuild bundle. In watch mode, a failed type check skips that rebuild and watching continues.

## Build fails on syntax or import

Check the first esbuild error and verify the referenced source file and import path exist. jspackr reports build failure and exits; it does not continue with a partial successful bundle. In watch mode, a failed rebuild is logged and watching continues.

## Output format and Node.js

For ESM, use an `.mjs` output file or set the containing package's `"type": "module"`. The default IIFE is a browser-oriented self-invoking script.
