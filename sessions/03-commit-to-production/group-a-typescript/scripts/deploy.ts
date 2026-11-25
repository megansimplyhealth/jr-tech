// "Deploys" the build in dist/ by copying it into .environments/<env>/.
// On a real team this would push to a server or cloud service instead.

import { appendFileSync, cpSync, existsSync, readFileSync, rmSync, writeFileSync } from "node:fs";

const env = process.argv[2];
if (env !== "staging" && env !== "production") {
  console.error("usage: npm run deploy -- staging|production");
  process.exit(2);
}
if (!existsSync("dist/release.json")) {
  console.error("✗ nothing to deploy: dist/ is empty. The build has to run first.");
  process.exit(1);
}

const release = JSON.parse(readFileSync("dist/release.json", "utf8"));
const target = `.environments/${env}`;
rmSync(target, { recursive: true, force: true });
cpSync("dist", target, { recursive: true });
writeFileSync(
  `${target}/release.json`,
  JSON.stringify(
    { ...release, environment: env, deployedAt: new Date().toISOString(), deployedBy: process.env.GITHUB_ACTOR ?? "local" },
    null,
    2,
  ) + "\n",
);
console.log(`✓ deployed quote-calculator ${release.version} to ${env}`);

const notes = [`### Deployed quote-calculator \`${release.version}\` to ${env}`];
if (!process.env.PAYMENTS_API_KEY) {
  console.log("⚠ PAYMENTS_API_KEY isn't set, so pay links won't work in this environment.");
  notes.push("⚠️ `PAYMENTS_API_KEY` isn't set, so pay links won't work here.");
}
// On GitHub Actions, anything written to this file shows on the run's summary page.
if (process.env.GITHUB_STEP_SUMMARY) appendFileSync(process.env.GITHUB_STEP_SUMMARY, notes.join("\n\n") + "\n");
