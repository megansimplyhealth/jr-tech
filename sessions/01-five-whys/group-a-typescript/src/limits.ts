import type { Benefit, Claim, Decision } from "./types.ts";

/** Annual benefit limits per member, in pounds. Reset on policy anniversary. */
export const ANNUAL_LIMITS: Record<Benefit, number> = {
  dental: 150,
  optical: 100,
  physio: 200,
};

/** Total already paid out to the member for this benefit this policy year. */
export function usedThisYear(history: Claim[]): number {
  return history.reduce((total, claim) => total + claim.amount, 0);
}

/**
 * Decide how much of a claim to pay.
 *  - nothing left in the pot  -> REJECTED
 *  - claim bigger than pot    -> PARTIAL (pay what's left)
 *  - otherwise                -> APPROVED in full
 */
export function decide(claim: Claim, history: Claim[]): Decision {
  const limit = ANNUAL_LIMITS[claim.benefit];
  const remaining = limit - usedThisYear(history);

  if (remaining <= 0) {
    return {
      claimId: claim.id,
      memberId: claim.memberId,
      outcome: "REJECTED",
      payout: 0,
      remaining,
      reason: `Annual ${claim.benefit} limit of £${limit} reached`,
    };
  }

  if (claim.amount > remaining) {
    return {
      claimId: claim.id,
      memberId: claim.memberId,
      outcome: "PARTIAL",
      payout: remaining,
      remaining,
      reason: `Only £${remaining} of £${limit} ${claim.benefit} limit left`,
    };
  }

  return {
    claimId: claim.id,
    memberId: claim.memberId,
    outcome: "APPROVED",
    payout: claim.amount,
    remaining,
    reason: "Within limit",
  };
}
