#!/usr/bin/env node
"use strict";

const { existsSync, mkdirSync, copyFileSync, chmodSync } = require("fs");
const { join } = require("path");

const PLATFORM_PACKAGES = {
  "win32-x64": "@hostclube/hyper-win32-x64",
  "darwin-arm64": "@hostclube/hyper-darwin-arm64",
  "darwin-x64": "@hostclube/hyper-darwin-x64",
  "linux-x64": "@hostclube/hyper-linux-x64",
  "linux-arm64": "@hostclube/hyper-linux-arm64",
};

function main() {
  const platformKey = `${process.platform}-${process.arch}`;
  const pkgName = PLATFORM_PACKAGES[platformKey];

  if (!pkgName) {
    console.error(
      `Unsupported platform: ${platformKey}. Supported: ${Object.keys(PLATFORM_PACKAGES).join(", ")}`
    );
    process.exit(1);
  }

  const isWindows = process.platform === "win32";
  const binaryName = isWindows ? "hyper.exe" : "hyper";

  // Find the platform package in node_modules
  let platformBinary;
  try {
    platformBinary = require.resolve(`${pkgName}/bin/${binaryName}`);
  } catch (_) {
    console.error(
      `Platform package ${pkgName} not found. Try reinstalling:\n  npm install -g @hostclube/hyper`
    );
    process.exit(1);
  }

  const nativeDir = join(__dirname, "bin", "native");
  if (!existsSync(nativeDir)) {
    mkdirSync(nativeDir, { recursive: true });
  }

  const destPath = join(nativeDir, binaryName);

  if (existsSync(destPath)) {
    return;
  }

  copyFileSync(platformBinary, destPath);

  if (!isWindows) {
    chmodSync(destPath, 0o755);
  }

  console.log("Hyper installed successfully.");
}

main();
