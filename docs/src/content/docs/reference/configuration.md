---
title: Configuration reference
description: JSON keys, defaults, and accepted values for jspackr configuration files.
---

Configuration files are JSON objects. The default file name is `jspackr.config.json` in the current working directory. Pass `-c` / `--config` to load another file.

| JSON key    | Type    | Default          | Accepted values / meaning                                                      |
| ----------- | ------- | ---------------- | ------------------------------------------------------------------------------ |
| `input`     | string  | none (required)  | Entry file path.                                                               |
| `output`    | string  | `dist/bundle.js` | Output file path.                                                              |
| `minify`    | boolean | `false`          | Enable all three esbuild minification modes.                                   |
| `report`    | boolean | `false`          | Print build size and module report.                                            |
| `sourcemap` | string  | `none`           | `none`, `l` (linked), or `in` (inline). Note the JSON spelling is `sourcemap`. |
| `format`    | string  | `iife`           | `iife`, `esm`, or `cjs`.                                                       |
| `watch`     | boolean | `false`          | Rebuild on watched file changes.                                               |
| `typeCheck` | boolean | `false`          | Run the local TypeScript compiler with `--noEmit` before each build. Requires TypeScript installed in the project. |
| `logLevel`  | string  | `info`           | `debug`, `info`, `warn`, or `error`.                                           |
| `force`     | boolean | `false`          | Skip overwrite confirmation.                                                   |
| `yes`       | boolean | `false`          | Automatically confirm overwrite prompts.                                       |
| `noConfirm` | boolean | `false`          | Skip all confirmation prompts.                                                 |

Example:

```json
{
  "input": "src/index.js",
  "output": "dist/app.js",
  "typeCheck": true,
  "minify": true,
  "sourcemap": "l",
  "format": "esm",
  "logLevel": "warn"
}
```

CLI strings override configuration strings when supplied. Boolean CLI flags only set true, so an enabled config option remains enabled when its flag is omitted. `--type-check` can enable checking from the CLI; `"typeCheck": true` enables it from this file. See [configuration guide](/guides/configuration/) for an example workflow.
