# jspackr

A minimal JavaScript bundler powered by esbuild.
Fast, simple, and focused on doing one thing well.

## Why jspackr?

jspackr exists because sometimes you just want to:

- bundle a JavaScript entry file
- optionally minify it
- see which files actually make your bundle big

No config files.
No plugins.
No framework cosplay.

Run it, get your bundle, move on.

## Features

- ⚡ Extremely fast bundling via esbuild
- ✂️ Optional minification
- 📊 Simple build report to see top contributors
- 🧠 Sensible CLI UX with shorthand flags

## Installation

For now, build from source:

```bash
npm install -g jspackr
```

## Usage

-> Basic usage:

```bash
jspackr src/index.js
```

-> With options:

```bash
jspackr -i src/index.js -o dist/app.js -m -r
```

## CLI Options

| Flag | Long Form       | Description                           |
| ---- | --------------- | ------------------------------------- |
| -i   | --input \<file> | Entry file                            |
| -o   | --out \<file>   | Output file (default: dist/bundle.js) |
| -m   | --minify        | Minify output                         |
| -r   | --report        | Show build report                     |
| -h   | --help          | Show help                             |

---

Created with ❤️ by Kaloka Radia Nanda  
Licensed under the MIT License.
