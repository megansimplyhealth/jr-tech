// Smoke test: is the deployed service alive and giving sensible answers?
// It checks one happy path, not every rule. That's the unit tests' job.

import { existsSync } from "node:fs";
import { resolve } from "node:path";
import { pathToFileURL } from "node:url";

const env = process.argv[2] ?? "staging";
const entry = resolve(`.environments/${env}/quote.ts`);
if (!existsSync(entry)) {
  console.error(`✗ nothing deployed to ${env}`);
  process.exit(1);
}

process.env.QUIET_LOGS = "1";
const { createQuote } = await import(pathToFileURL(entry).href);
const q = createQuote(
  { memberName: "Smoke Test", dateOfBirth: "1990-05-01", email: "smoke@example.com", plan: "plus" },
  new Date("2026-06-01T09:00:00Z"),
);
if (q.monthly !== "£21.00" || !q.payLink.startsWith("https://")) {
  console.error(`✗ smoke test on ${env}: expected £21.00 with a pay link, got ${JSON.stringify(q)}`);
  process.exit(1);
}
console.log(`✓ smoke test on ${env}: plus plan, age ${q.age} → ${q.monthly}/month`);
