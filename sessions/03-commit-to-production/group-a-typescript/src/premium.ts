import { BAND_MULTIPLIER, BASE_PENCE, type AgeBand, type Plan } from "./plans.ts";

export function ageBand(age: number): AgeBand {
  if (!Number.isInteger(age) || age < 18 || age > 120) {
    throw new RangeError(`can't quote for age ${age}, members must be 18 to 120`);
  }
  if (age >= 65) return "65+";
  if (age >= 40) return "40-64";
  return "18-39";
}

export function monthlyPremiumPence(plan: Plan, age: number): number {
  return Math.round(BASE_PENCE[plan] * BAND_MULTIPLIER[ageBand(age)]);
}

export function formatGBP(pence: number): string {
  return `£${(pence / 100).toFixed(2)}`;
}
