// Builds the release bundle in dist/ and stamps it with a version.

import { cpSync, mkdirSync, readFileSync, rmSync, writeFileSync } from "node:fs";

const pkg = JSON.parse(readFileSync("package.json", "utf8"));
const version = `${pkg.version}-run${process.env.GITHUB_RUN_NUMBER ?? "local"}`;

rmSync("dist", { recursive: true, force: true });
mkdirSync("dist");
cpSync("src", "dist", { recursive: true });
writeFileSync(
  "dist/release.json",
  JSON.stringify({ service: "quote-calculator", version, builtAt: new Date().toISOString() }, null, 2) + "\n",
);
console.log(`✓ built quote-calculator ${version} into dist/`);
