---
title: Source maps
description: Emit linked or inline source maps for jspackr bundles.
---

Source maps let debugging tools map generated bundle locations back to source files. jspackr supports `none` (default), `l` (linked), and `in` (inline).

Create an external map:

```sh
jspackr -i src/index.js -o dist/app.js --source l
```

This writes `dist/app.js` and `dist/app.js.map`; the bundle references the map. For an embedded map, use:

```sh
jspackr -i src/index.js -o dist/app.js --source in
```

The short alias `-s` accepts the same values. JSON configuration uses the key `sourcemap`:

```json
{ "input": "src/index.js", "sourcemap": "l" }
```

When watch mode is active, changes to the generated bundle and its `.map` file are ignored to avoid rebuild loops. See [watch mode](/guides/watch-mode/).
