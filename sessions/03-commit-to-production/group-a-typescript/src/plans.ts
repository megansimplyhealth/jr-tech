// BrightCare cash plans and how age changes the price.

export type Plan = "essential" | "plus" | "complete";
export type AgeBand = "18-39" | "40-64" | "65+";

// Monthly price for the youngest band, in pence.
export const BASE_PENCE: Record<Plan, number> = {
  essential: 1200,
  plus: 2100,
  complete: 3400,
};

// Agreed with the pricing team, see PRICING-2026.md in the product wiki.
export const BAND_MULTIPLIER: Record<AgeBand, number> = {
  "18-39": 1,
  "40-64": 1.35,
  "65+": 1.8,
};

export const PLANS = Object.keys(BASE_PENCE) as Plan[];
