---
title: Watch mode
description: Rebuild a bundle after changes in the entry file's directory tree.
---

Watch mode is useful during development. It watches directories beneath the entry file's directory (or the entry directory if an existing directory is passed), then rebuilds after file writes, creates, removals, or renames.

```sh
jspackr -i src/index.js -o dist/app.js --watch
```

Watch mode starts watching without performing an initial build. Run a normal build first when you need the output immediately, then start the watcher. Changes are debounced for 300 ms. Newly created directories are added to the watch. `.git` and `node_modules` subdirectories are skipped. Generated output and the corresponding `.map` file are ignored. A failed rebuild is logged and the watcher continues. Press Ctrl+C to stop.

Set `"watch": true` in the config to enable it. CLI flags can enable watch mode but cannot disable it when enabled in config.
