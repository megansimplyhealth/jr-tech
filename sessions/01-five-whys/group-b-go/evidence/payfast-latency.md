# PayFast response time (p95) — prod vs staging, Tue 22 Sep

| Time (UTC) | Prod p95 | Staging sandbox p95 |
|---|---|---|
| 09:00–09:30 | 340 ms | 170 ms |
| 11:00–11:30 | 410 ms | 180 ms |
| **12:00–12:45** | **2,450 ms** | 175 ms |
| 13:00–13:30 | 620 ms | 180 ms |
| 16:00–16:30 | 380 ms | 170 ms |

**Note from PayFast status page:** *"Payment submission latency is typically higher between 12:00 and 13:00
UK time due to peak volumes from UK payroll and benefits providers. Integrators should allow a timeout of at
least 10 seconds."*

**Note:** The staging sandbox returns a canned response and never calls a bank, so it is always fast.
