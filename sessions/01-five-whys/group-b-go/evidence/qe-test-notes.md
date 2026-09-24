# QE notes — PR #498 retry hotfix (18 Aug)

**Tested in:** staging, PayFast sandbox

| Scenario | How | Result |
|---|---|---|
| 503 on first call → retry succeeds | Sandbox "force 503" switch | ✅ 1 payment |
| 503 on every call → gives up after 3 | Sandbox "force 503" switch | ✅ 0 payments, marked FAILED |
| 400 bad request → no retry | Sent invalid sort code | ✅ |
| **Slow response / timeout** | — | ⏭️ **Not tested** — sandbox has no way to simulate slow responses |

**Risks noted:** Timeout behaviour untested. Raised verbally at stand-up; agreed to cover under PAY-771.

*(PR #530 — timeout change — was not sent to QE: labelled `perf`, "config only".)*
