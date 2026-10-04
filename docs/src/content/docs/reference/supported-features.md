---
title: Supported features
description: Current jspackr build behavior and the limits of its CLI surface.
---

## Build behavior

| Capability | Current behavior |
| --- | --- |
| Entry points | One entry point per invocation. |
| Bundling | Always enabled; imported modules are bundled into one output file. |
| Platform | esbuild browser platform. |
| Output formats | IIFE (default), ESM, CommonJS. |
| Minification | Optional whitespace, syntax, and identifier minification together. |
| Source maps | None (default), linked, or inline. |
| Reporting | Sizes, module count, elapsed build time, and largest input contributors. |
| Watch | Rebuild on file changes in the entry directory tree. |
| Configuration | JSON file with CLI overrides. |
| TypeScript / JSX syntax | Transformed by esbuild according to file extension. Optional type checking uses local `tsc --noEmit`, enabled with `--type-check` or config key `typeCheck`. |

The esbuild dependency handles source parsing and transforms; jspackr does not expose all esbuild controls. For example, there are no CLI/config controls for target, external packages, define values, plugins, multiple entries, splitting, or JSX runtime configuration.

Type checking uses the project's `tsconfig.json` directly and does not override its file selection or compiler settings. JavaScript is checked only when enabled by TypeScript configuration. In watch mode, type checking runs before each rebuild triggered by a watched change. The repository has watcher, configuration validation, and CLI end-to-end tests. It does not provide a separate documented plugin API or a generated config schema.

## Report details

`--report` enables esbuild metadata and prints the total input bytes, output bytes, percentage size ratio, module count, elapsed time, and up to five largest input contributors. It does not write a standalone report file.
