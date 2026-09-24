# PR #498 — fix: retry PayFast calls on failure (INC-2250 hotfix)

**Opened:** Tue 18 Aug 15:20 **Merged:** Tue 18 Aug 15:41 **Labels:** `hotfix`, `incident`

## Description

PayFast is returning intermittent 503s (INC-2250). Payouts are failing and members are
calling. This retries any failed call up to 3 times with a short backoff.

- Retry on network errors and 5xx
- Don't retry on 4xx
- Unit test for the 503 → retry path

## Review

> **Reviewer 1 (15:28):** Is it safe to retry a payment? If the first call actually went through
> we'd pay twice. PayFast supports an `Idempotency-Key` header — should we send one?

> **Author (15:31):** Good shout. 503 means they didn't process it so we're fine for this outage.
> I've raised **PAY-771** to add idempotency keys properly as a follow-up.

> **Reviewer 1 (15:33):** 👍 approving so we can stop the bleeding. Let's make sure PAY-771 gets done.

> **QE (15:39):** Retested in staging with the sandbox's "force 503" switch — retries work, one payment made. ✅
