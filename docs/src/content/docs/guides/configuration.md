---
title: Configuration
description: Set build options in jspackr.config.json and override them from the CLI.
---

jspackr reads JSON configuration. By default it looks for `jspackr.config.json` in the current working directory. Or pass a path with `--config` / `-c`.

```json title="jspackr.config.json"
{
  "input": "./src/index.js",
  "output": "./dist/bundle.js",
  "minify": true,
  "report": false,
  "sourcemap": "none",
  "format": "iife",
  "typeCheck": false,
  "watch": false,
  "logLevel": "info"
}
```

Run with the discovered config:

```sh
jspackr
```

Or specify it explicitly and override the output:

```sh
jspackr --config ./config/production.json --out ./public/app.js
```

Non-empty string options from the CLI override config values. Boolean flags such as `--minify`, `--report`, `--watch`, and `--type-check` turn a value on; omitting a flag does not turn a config value off. There are no negative boolean flags. See the [configuration reference](/reference/configuration/) for keys, defaults, and accepted values.

`force`, `yes`, and `noConfirm` are also valid config keys for confirmation behavior. See [CLI reference](/cli/reference/) for their distinction.
