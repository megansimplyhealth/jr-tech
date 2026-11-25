import { payments } from "./config.ts";
import { log } from "./logger.ts";
import type { AgeBand, Plan } from "./plans.ts";
import { ageBand, formatGBP, monthlyPremiumPence } from "./premium.ts";

export interface QuoteRequest {
  memberName: string;
  dateOfBirth: string; // YYYY-MM-DD
  email: string;
  plan: Plan;
}

export interface Quote {
  plan: Plan;
  age: number;
  ageBand: AgeBand;
  monthlyPence: number;
  monthly: string;
  payLink: string;
}

export function ageOn(dateOfBirth: string, today: Date): number {
  const dob = new Date(`${dateOfBirth}T00:00:00Z`);
  if (Number.isNaN(dob.getTime())) throw new RangeError(`not a date: ${dateOfBirth}`);
  let age = today.getUTCFullYear() - dob.getUTCFullYear();
  const birthdayPassed =
    today.getUTCMonth() > dob.getUTCMonth() ||
    (today.getUTCMonth() === dob.getUTCMonth() && today.getUTCDate() >= dob.getUTCDate());
  if (!birthdayPassed) age--;
  return age;
}

export function createQuote(request: QuoteRequest, today: Date = new Date()): Quote {
  const age = ageOn(request.dateOfBirth, today);
  const monthlyPence = monthlyPremiumPence(request.plan, age);
  const quote: Quote = {
    plan: request.plan,
    age,
    ageBand: ageBand(age),
    monthlyPence,
    monthly: formatGBP(monthlyPence),
    payLink: `${payments.baseUrl}/checkout?plan=${request.plan}&amount=${monthlyPence}`,
  };
  log("quote.created", { ...request, ageBand: quote.ageBand, monthlyPence });
  return quote;
}
