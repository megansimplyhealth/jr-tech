import assert from "node:assert/strict";
import { test } from "node:test";
import { ageOn, createQuote } from "../src/quote.ts";

process.env.QUIET_LOGS = "1";
const today = new Date("2026-12-03T10:00:00Z");

test("works out age from date of birth", () => {
  assert.equal(ageOn("1990-12-03", today), 36);
  assert.equal(ageOn("1990-12-04", today), 35);
});

test("creates a quote with a pay link", () => {
  const q = createQuote({ memberName: "Test Member", dateOfBirth: "1980-01-15", email: "t@example.com", plan: "plus" }, today);
  assert.equal(q.age, 46);
  assert.equal(q.ageBand, "40-64");
  assert.equal(q.monthly, "£28.35");
  assert.match(q.payLink, /^https:\/\/pay\.brightcare\.example\/checkout\?plan=plus&amount=2835$/);
});
