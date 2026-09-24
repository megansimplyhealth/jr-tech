# claims-engine — monitoring & alerts (exported from dashboard config)

| Alert | Condition | Fired 16–18 Sep? |
|---|---|---|
| High error rate | 5xx > 2% for 5 min | No (0%) |
| High latency | p95 > 800ms for 10 min | No |
| Service down | health check fails x3 | No |

**Dashboards:** request rate, latency, error rate, CPU / memory.

**Business checks:** none. Payout vs limit is reconciled **manually by Finance every Friday**.

**Note from on-call:** the decision logs are structured JSON, but nobody queries them unless there's
an error. `remaining` is not graphed.
