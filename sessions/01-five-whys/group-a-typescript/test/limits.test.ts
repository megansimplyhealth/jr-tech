import { test } from "node:test";
import assert from "node:assert/strict";
import { decide, usedThisYear } from "../src/limits.ts";
import type { Claim } from "../src/types.ts";

const claim = (amount: number, overrides: Partial<Claim> = {}): Claim => ({
  id: "CLM-1",
  memberId: "M-1",
  benefit: "dental",
  amount,
  submittedAt: "2026-06-01T10:00:00Z",
  channel: "web",
  ...overrides,
});

test("usedThisYear sums previous claims", () => {
  assert.equal(usedThisYear([claim(60), claim(50)]), 110);
});

test("approves a claim within the limit", () => {
  const d = decide(claim(40), [claim(60)]);
  assert.equal(d.outcome, "APPROVED");
  assert.equal(d.payout, 40);
});

test("partially pays a claim that exceeds what is left", () => {
  const d = decide(claim(80), [claim(60), claim(50)]);
  assert.equal(d.outcome, "PARTIAL");
  assert.equal(d.payout, 40);
});

test("rejects a claim when the limit is used up", () => {
  const d = decide(claim(20), [claim(100), claim(50)]);
  assert.equal(d.outcome, "REJECTED");
  assert.equal(d.payout, 0);
});

test("limits are per benefit", () => {
  const d = decide(claim(90, { benefit: "optical" }), []);
  assert.equal(d.outcome, "APPROVED");
});
