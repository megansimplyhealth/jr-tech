import assert from "node:assert/strict";
import { test } from "node:test";
import { ageBand, formatGBP, monthlyPremiumPence } from "../src/premium.ts";

test("18 to 39 is the youngest band", () => {
  assert.equal(ageBand(18), "18-39");
  assert.equal(ageBand(39), "18-39");
});

test("40 to 64 is the middle band", () => {
  assert.equal(ageBand(40), "40-64");
  assert.equal(ageBand(64), "40-64");
});

test("65 and over is the 65+ band", () => {
  assert.equal(ageBand(65), "65+");
  assert.equal(ageBand(66), "65+");
  assert.equal(ageBand(90), "65+");
});

test("won't quote under 18", () => {
  assert.throws(() => ageBand(17), RangeError);
});

test("prices for a 30 year old", () => {
  assert.equal(monthlyPremiumPence("essential", 30), 1200);
  assert.equal(monthlyPremiumPence("plus", 30), 2100);
  assert.equal(monthlyPremiumPence("complete", 30), 3400);
});

test("price for a 70 year old on complete", () => {
  assert.equal(monthlyPremiumPence("complete", 70), 6120);
});

test("formats pence as pounds", () => {
  assert.equal(formatGBP(2835), "£28.35");
  assert.equal(formatGBP(1200), "£12.00");
});
