---
title: Installation
description: Install jspackr with Go or use the npm package where a binary is published.
---

## Requirements

Building from source requires Go 1.26 or newer, as specified by the repository's `go.mod`. The npm launcher requires Node.js 18 or newer. jspackr produces browser bundles; Node is only needed to run the optional npm launcher.

## Install with Go

Install the latest module version and ensure Go's binary directory is on your `PATH`:

```sh
go install github.com/kalokaradia/jspackr/src/main@latest
```

## Install from npm

The npm package contains prebuilt binaries for Linux x64, Linux arm64, macOS x64, and macOS arm64. It does not currently contain Windows binaries, although the launcher recognizes Windows architecture names. On Windows or an unsupported platform, install from source with Go.

```sh
npm install --global jspackr
```

## Verify

```sh
jspackr --version
jspackr --help
```

The CLI currently reports version `0.5.0` by default. Release builds can inject another version; `--version` prints the version embedded in the executable.
