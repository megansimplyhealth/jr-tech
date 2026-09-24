import type { Claim } from "./types.ts";

/**
 * Turns an incoming HTTP request body into a Claim.
 *
 * Web and mobile both POST to /claims. Mobile v2 was switched over to this
 * endpoint in PR #412 - "same shape as web", so no extra handling needed.
 */
export function parseClaim(body: string): Claim {
  const raw = JSON.parse(body);

  if (!raw.id || !raw.memberId || !raw.benefit) {
    throw new Error("Invalid claim: missing id, memberId or benefit");
  }

  return raw as Claim;
}
