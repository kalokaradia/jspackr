---
title: Quick start
description: Build and run your first jspackr JavaScript bundle.
---

## 1. Create an entry file

```js title="src/index.js"
import { message } from "./message.js";

console.log(message);
```

```js title="src/message.js"
export const message = "Bundled with jspackr";
```

## 2. Build

From the project directory, run:

```sh
jspackr --input src/index.js --out dist/bundle.js
```

jspackr follows the import and writes a single browser bundle. It creates the output directory when needed. The default format is an immediately invoked function expression (IIFE).

## 3. Run the output

For a quick check in Node.js, run the IIFE bundle:

```sh
node dist/bundle.js
```

Expected output:

```text
Bundled with jspackr
```

This example uses browser platform build defaults. Browser globals such as `window` are available when the output runs in a browser; Node-specific package behavior is not selected by jspackr.

Continue with [configuration](/guides/configuration/) or review [all CLI options](/cli/reference/).
