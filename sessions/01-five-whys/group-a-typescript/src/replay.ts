/**
 * Replays every claim received on the incident day (16 Sep 2026) through the
 * claims engine, exactly as production processed them.
 *
 *   npm run replay                     -> decision table + per-member totals
 *   npm run replay -- --member M-1042  -> only one member
 *   npm run replay -- --log            -> raw engine log lines instead
 */
import { readFileSync } from "node:fs";
import { ClaimsEngine } from "./engine.ts";
import { ANNUAL_LIMITS } from "./limits.ts";
import type { Decision } from "./types.ts";

const args = process.argv.slice(2);
const memberFilter = args.includes("--member") ? args[args.indexOf("--member") + 1] : undefined;
const rawLogs = args.includes("--log");

const lines = readFileSync(new URL("../fixtures/claims-2026-09-16.jsonl", import.meta.url), "utf8")
  .split("\n")
  .filter(Boolean);

const logLines: string[] = [];
const engine = new ClaimsEngine(undefined, (l) => logLines.push(l));

type Row = { time: string; channel: string; benefit: string; amount: unknown; d: Decision };
const rows: Row[] = [];

for (const body of lines) {
  const d = engine.handle(body);
  const raw = JSON.parse(body);
  rows.push({ time: raw.submittedAt.slice(11, 16), channel: raw.channel, benefit: raw.benefit, amount: raw.amount, d });
}

if (rawLogs) {
  for (const l of logLines) {
    if (!memberFilter || l.includes(`"${memberFilter}"`)) console.log(l);
  }
  process.exit(0);
}

const pad = (s: unknown, n: number) => String(s).padEnd(n);
console.log("\nClaims engine replay - 16 Sep 2026\n");
console.log(pad("time", 7) + pad("channel", 9) + pad("member", 9) + pad("claim", 11) + pad("benefit", 9) + pad("claimed", 10) + pad("outcome", 10) + "payout");
console.log("-".repeat(72));
for (const r of rows) {
  if (memberFilter && r.d.memberId !== memberFilter) continue;
  console.log(
    pad(r.time, 7) + pad(r.channel, 9) + pad(r.d.memberId, 9) + pad(r.d.claimId, 11) + pad(r.benefit, 9) +
      pad(`£${r.amount}`, 10) + pad(r.d.outcome, 10) + `£${r.d.payout}`,
  );
}

// Finance's view: total actually paid per member+benefit vs the annual limit
console.log("\nFinance reconciliation (paid today vs annual limit)\n");
const totals = new Map<string, { paid: number; limit: number }>();
for (const r of rows) {
  if (memberFilter && r.d.memberId !== memberFilter) continue;
  const k = `${r.d.memberId} ${r.benefit}`;
  const t = totals.get(k) ?? { paid: 0, limit: ANNUAL_LIMITS[r.benefit as keyof typeof ANNUAL_LIMITS] };
  t.paid += Number(r.d.payout);
  totals.set(k, t);
}
for (const [k, t] of totals) {
  const flag = t.paid > t.limit ? "  <-- OVER LIMIT" : "";
  console.log(pad(k, 18) + pad(`paid £${t.paid.toFixed(2)}`, 16) + `limit £${t.limit}` + flag);
}
console.log();
