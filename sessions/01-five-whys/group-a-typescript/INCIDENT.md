# INC-2291 — Dental & physio claims paid over annual limit

| | |
|---|---|
| **Severity** | SEV-2 (financial impact, customer impact) |
| **Service** | `claims-engine` (TypeScript) |
| **Detected** | Fri 18 Sep 2026, 14:30 — by the **Finance reconciliation team**, not by monitoring |
| **Window** | Wed 16 Sep 09:00 → Fri 18 Sep 14:30 |
| **Status** | Mitigated (mobile claims paused) — root cause **unknown** |

## What we know

- **143 members** were paid more than their annual benefit limit. Total over-payment: **£11,420**.
- **9 members** had claims **rejected** even though they still had money left in their allowance.
  One of them (M-1103) phoned the contact centre: *"I had £20 left on dental and my claim was declined."*
- Only claims submitted through the **mobile app** seem affected. Web claims look fine.
- Not every mobile member was affected. Nobody can explain the pattern yet.
- No alerts fired. Error rate on `claims-engine` was 0% the whole time.

## Timeline

| When | What |
|---|---|
| Mon 14 Sep 16:40 | Mobile app v2.0 released to app stores (new claims journey) |
| Tue 15 Sep 11:05 | PR #412 merged & deployed: `claims-engine` accepts mobile v2 claims on `/claims` |
| Wed 16 Sep 09:00 | Mobile v2 traffic starts hitting `claims-engine` (feature flag on) |
| Fri 18 Sep 14:30 | Finance flags members paid £200+ on a £150 dental limit |
| Fri 18 Sep 15:10 | Mobile claims paused via feature flag. Incident opened. |

## Your evidence (all in this folder)

| Where | What it is |
|---|---|
| `npm run replay` | Replays a sample of 20 real claims from 16 Sep through the engine |
| `npm run replay -- --log` | The raw engine logs for those claims (also in `evidence/claims-engine-2026-09-16.log`) |
| `fixtures/claims-2026-09-16.jsonl` | The raw request bodies received |
| `src/` | The claims engine code |
| `test/` | The existing tests (they all pass…) |
| `evidence/git-log.txt` | Recent commits |
| `evidence/pr-412.md` | The PR that switched mobile onto this endpoint |
| `evidence/slack-mobile-squad.md` | A thread from the mobile squad's channel |
| `evidence/qe-test-plan-mobile-v2.md` | The QE test plan for the mobile v2 release |
| `evidence/alerting.md` | What monitoring exists for `claims-engine` |
