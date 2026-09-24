# INC-2307 — Members reimbursed two, three or four times

| | |
|---|---|
| **Severity** | SEV-2 (financial impact) |
| **Service** | `payout-service` (Go) → PayFast payments provider |
| **Detected** | Wed 23 Sep 2026, 09:15 — a **member phoned** to say they'd been paid three times for one claim |
| **Window** | Tue 22 Sep 12:05 → 12:45 |
| **Status** | Mitigated (lunchtime batches paused) — root cause **unknown** |

## What we know

- **38 members** were paid more than once for the same claim. Total over-payment: **£6,870**.
- Some were paid **twice**, some **three** times, and a few **four** times (e.g. M-2090 received £800 for a £200 claim).
- Our own dashboard says most of these payouts succeeded **once**. Some even say **FAILED**.
- PayFast's settlement file for 22 Sep had **214 payments**. We only approved **176**.
- The 09:05 and 16:05 batches the same day were fine.
- An alert **did** fire at 12:07. It was acknowledged and resolved within 3 minutes.

## Timeline

| When | What |
|---|---|
| Tue 18 Aug | INC-2250: PayFast outage. Hotfix PR #498 adds retries to the payout client. |
| Tue 15 Sep | PR #530: PayFast timeout reduced 10s → 2s |
| Tue 22 Sep 12:05 | Lunchtime payout batch starts |
| Tue 22 Sep 12:07 | Alert `PayoutFailuresHigh` fires → acked by on-call, resolved 12:10 |
| Tue 22 Sep 13:05 | Next batch re-sends the payouts marked FAILED at 12:05 |
| Wed 23 Sep 06:00 | PayFast daily settlement file arrives (nobody reads it until Friday) |
| Wed 23 Sep 09:15 | Member calls contact centre: "I've been paid three times" |

## Your evidence (all in this folder)

| Where | What it is |
|---|---|
| `go run ./cmd/replay` | Replays the 12:05 batch against a PayFast simulator (takes ~8s) |
| `go run ./cmd/replay -log` | Raw payout-service logs (also in `evidence/payout-service-2026-09-22.log`) |
| `internal/payout/` | The payout client code + tests |
| `internal/provider/` | The PayFast simulator (behaves like the real thing) |
| `evidence/git-log.txt` | Recent commits |
| `evidence/pr-498-hotfix.md`, `evidence/pr-530.md` | The two most recent PRs to the client |
| `evidence/PAY-771.md` | A backlog ticket |
| `evidence/payfast-latency.md` | PayFast response times: prod vs staging |
| `evidence/payfast-api-docs.md` | Extract from PayFast's API documentation |
| `evidence/qe-test-notes.md` | QE notes from testing PR #498 |
| `evidence/alerting.md` | What monitoring exists, and what happened with the 12:07 alert |
