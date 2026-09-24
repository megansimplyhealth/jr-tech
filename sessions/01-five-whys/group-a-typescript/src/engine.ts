import { parseClaim } from "./intake.ts";
import { decide } from "./limits.ts";
import { ClaimStore } from "./store.ts";
import type { Decision } from "./types.ts";

export type Logger = (line: string) => void;

/** The claims engine: POST /claims -> decision -> payout instruction. */
export class ClaimsEngine {
  private store: ClaimStore;
  private log: Logger;

  constructor(store: ClaimStore = new ClaimStore(), log: Logger = console.log) {
    this.store = store;
    this.log = log;
  }

  handle(body: string): Decision {
    const claim = parseClaim(body);
    const history = this.store.history(claim.memberId, claim.benefit);
    const decision = decide(claim, history);

    if (decision.payout > 0) {
      this.store.recordPayout(claim, decision.payout);
    }

    this.log(
      JSON.stringify({
        ts: claim.submittedAt,
        level: "info",
        msg: "claim.decided",
        channel: claim.channel,
        ...decision,
      }),
    );
    return decision;
  }
}
