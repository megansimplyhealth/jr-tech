# payout-service — monitoring & alerts

| Alert | Condition | Fired 22 Sep? |
|---|---|---|
| `PayoutFailuresHigh` | > 5 payouts FAILED in 10 min | **Yes, 12:07** |
| `PayoutServiceDown` | health check fails x3 | No |

**Not monitored:** number of payments PayFast made vs number of payouts we approved.
Reconciliation against PayFast's daily settlement file is done **manually by Finance on Fridays**.

## #payments-oncall (Slack) — Tue 22 Sep

**Alertbot** 12:07 🔴 `PayoutFailuresHigh` — 7 failures in 10 min

**On-call engineer** 12:09
PayFast is slow again, it's their lunchtime peak 🙄 retries will sort it and anything that fails
gets picked up in the next batch. Acking.

**On-call engineer** 12:10
✅ resolved — failures back to 0

*(Runbook for `PayoutFailuresHigh`: "Usually PayFast being slow. Failed payouts are retried in the next batch automatically. No action needed unless it persists > 1 hour.")*
