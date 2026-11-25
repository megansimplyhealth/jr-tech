// Looks for passwords, keys and tokens that have been written into the code.
// Secrets belong in environment variables or a secrets store, never in git.

import { readdirSync, readFileSync, statSync } from "node:fs";
import { extname, join, relative } from "node:path";

const PATTERNS: { name: string; pattern: RegExp }[] = [
  { name: "BrightCare live API key", pattern: /bc_live_[A-Za-z0-9]{16,}/ },
  { name: "private key", pattern: new RegExp("-----BEGIN " + "[A-Z ]*PRIVATE KEY-----") },
  { name: "hard-coded password", pattern: /password\s*[:=]\s*["'][^"'\s]{6,}["']/i },
];
const SKIP = new Set(["node_modules", "dist", ".environments", ".git", "package-lock.json"]);
const TEXT = new Set([".ts", ".js", ".json", ".yml", ".yaml", ".md", ".env", ".txt", ".sh"]);

function files(dir: string): string[] {
  return readdirSync(dir).flatMap((name) => {
    if (SKIP.has(name)) return [];
    const path = join(dir, name);
    if (statSync(path).isDirectory()) return files(path);
    return TEXT.has(extname(name)) || name.startsWith(".env") ? [path] : [];
  });
}

const root = process.cwd();
const all = files(root);
let found = 0;
for (const file of all) {
  readFileSync(file, "utf8")
    .split("\n")
    .forEach((line, i) => {
      for (const { name, pattern } of PATTERNS) {
        const m = line.match(pattern);
        if (!m) continue;
        found++;
        console.error(`${relative(root, file)}:${i + 1}  ${name}  ${m[0].slice(0, 12)}…`);
      }
    });
}

if (found) {
  console.error(`\n✗ ${found} possible secret${found > 1 ? "s" : ""} found in ${all.length} files.`);
  console.error("  Read it from an environment variable instead (process.env.NAME).");
  console.error("  Anything that was ever committed is in git history, so it has to be rotated too.");
  process.exit(1);
}
console.log(`✓ no secrets found in ${all.length} files`);
