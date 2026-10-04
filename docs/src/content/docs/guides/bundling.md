---
title: Bundling
description: Bundle one JavaScript entry point and its imported modules for browser use.
---

jspackr always enables esbuild bundling and selects the browser platform. Imports reachable from the entry point are included in the output file.

```text title="Project files"
src/
├── index.js
└── greeting.js
```

```js title="src/index.js"
import { greet } from './greeting.js';
greet('Ada');
```

```js title="src/greeting.js"
export function greet(name) {
  console.log(`Hello, ${name}`);
}
```

```sh
jspackr -i src/index.js -o dist/app.js
```

The output contains the entry and imported code in one bundle. jspackr accepts one entry point per invocation and writes one output file. It does not expose code splitting or multi-entry builds.

## Output format

The default `iife` output runs as a self-invoking script. Use ESM when consumers need module imports, or CommonJS for a CommonJS loader:

```sh
jspackr -i src/index.js -o dist/app.mjs --format esm
jspackr -i src/index.js -o dist/app.cjs --format cjs
```

With Node.js, use `.mjs` or a package configured with `"type": "module"` for ESM output. Formats are described in the [supported features reference](/reference/supported-features/).
