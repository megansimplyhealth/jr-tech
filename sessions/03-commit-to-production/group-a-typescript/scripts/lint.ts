// House lint rules. Small on purpose: each one is here because it bit us once.

import { readdirSync, readFileSync, statSync } from "node:fs";
import { join, relative } from "node:path";

interface Rule {
  id: string;
  folders: string[];
  pattern: RegExp;
  why: string;
}

const RULES: Rule[] = [
  {
    id: "no-console",
    folders: ["src"],
    pattern: /\bconsole\.(log|info|debug|warn|error)\s*\(/,
    why: "console output isn't redacted and can leak member data into our logs. Use log() from src/logger.ts.",
  },
  {
    id: "no-only",
    folders: ["test"],
    pattern: /\b(test|it|describe)\.only\s*\(/,
    why: ".only makes the runner skip every other test.",
  },
  {
    id: "no-debugger",
    folders: ["src", "test"],
    pattern: /^\s*debugger\b/,
    why: "debugger statements pause the service in production.",
  },
];

function files(dir: string): string[] {
  return readdirSync(dir).flatMap((name) => {
    const path = join(dir, name);
    return statSync(path).isDirectory() ? files(path) : path.endsWith(".ts") ? [path] : [];
  });
}

const root = process.cwd();
let problems = 0;
let checked = 0;
for (const rule of RULES) {
  for (const folder of rule.folders) {
    for (const file of files(join(root, folder))) {
      checked++;
      readFileSync(file, "utf8")
        .split("\n")
        .forEach((line, i) => {
          if (line.trim().startsWith("//")) return;
          const m = line.match(rule.pattern);
          if (!m) return;
          problems++;
          console.error(`${relative(root, file)}:${i + 1}:${(m.index ?? 0) + 1}  ${rule.id}  ${rule.why}`);
          console.error(`    ${line.trim()}`);
        });
    }
  }
}

if (problems) {
  console.error(`\n✗ ${problems} lint problem${problems > 1 ? "s" : ""}`);
  process.exit(1);
}
console.log(`✓ lint clean (${RULES.length} rules)`);
