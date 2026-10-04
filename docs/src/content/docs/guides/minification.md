---
title: Minification
description: Reduce bundle whitespace, syntax, and identifier names with --minify.
---

Minification can reduce the size of production output. jspackr maps `--minify` to esbuild's whitespace, syntax, and identifier minification options together.

```sh
jspackr -i src/index.js -o dist/app.js --minify
```

The result is still a bundle in the selected output format; names and formatting may be harder to read. Use an unminified build while debugging, or add `"minify": true` to the JSON config for a repeatable build.

There are no separate flags for each minification mode and no CLI option to disable minification from a config file. See [configuration precedence](/guides/configuration/) and [CLI reference](/cli/reference/).
