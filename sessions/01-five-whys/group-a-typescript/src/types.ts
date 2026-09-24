export type Benefit = "dental" | "optical" | "physio";
export type Channel = "web" | "mobile";

export interface Claim {
  id: string;
  memberId: string;
  benefit: Benefit;
  /** Amount claimed, in pounds (e.g. 45.5 = £45.50) */
  amount: number;
  submittedAt: string;
  channel: Channel;
}

export type Outcome = "APPROVED" | "PARTIAL" | "REJECTED";

export interface Decision {
  claimId: string;
  memberId: string;
  outcome: Outcome;
  payout: number;
  remaining: number;
  reason: string;
}
