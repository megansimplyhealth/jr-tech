import type { Benefit, Claim } from "./types.ts";

/**
 * In-memory stand-in for the claims database.
 * Stores every claim that resulted in a payout, keyed by member + benefit.
 */
export class ClaimStore {
  private paid: Map<string, Claim[]> = new Map();

  private key(memberId: string, benefit: Benefit): string {
    return `${memberId}:${benefit}`;
  }

  history(memberId: string, benefit: Benefit): Claim[] {
    return this.paid.get(this.key(memberId, benefit)) ?? [];
  }

  recordPayout(claim: Claim, payout: number): void {
    const k = this.key(claim.memberId, claim.benefit);
    const list = this.paid.get(k) ?? [];
    list.push({ ...claim, amount: payout });
    this.paid.set(k, list);
  }
}
