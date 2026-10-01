#!/usr/bin/env node

"use strict";

const { spawnSync } = require("node:child_process");
const fs = require("node:fs");
const path = require("node:path");

const platformMap = {
  linux: "linux",
  darwin: "darwin",
  win32: "win32",
};

const archMap = {
  x64: "x64",
  arm64: "arm64",
};

const platform = platformMap[process.platform];
const arch = archMap[process.arch];

if (!platform || !arch) {
  console.error(
    `jspackr: unsupported platform or architecture: ${process.platform}/${process.arch}`,
  );
  console.error(
    "Supported platforms: Linux, macOS, and Windows on x64 or arm64.",
  );
  process.exit(1);
}

const extension = platform === "win32" ? ".exe" : "";

const binaryName = `jspackr-${platform}-${arch}${extension}`;
const binary = path.join(__dirname, "bin", binaryName);

if (!fs.existsSync(binary)) {
  console.error(`jspackr: bundled binary is missing: ${binary}`);
  console.error("Reinstall the package or install jspackr manually with Go:");
  console.error("go install github.com/kalokaradia/jspackr/src/main@latest");
  process.exit(1);
}

if (platform !== "win32") {
  try {
    fs.chmodSync(binary, 0o755);
  } catch (error) {
    console.error(
      `jspackr: cannot make bundled binary executable: ${error.message}`,
    );
    process.exit(1);
  }
}

const result = spawnSync(binary, process.argv.slice(2), {
  stdio: "inherit",
});

if (result.error) {
  console.error(
    `jspackr: could not start bundled binary: ${result.error.message}`,
  );
  process.exit(1);
}

process.exit(result.status ?? 1);
