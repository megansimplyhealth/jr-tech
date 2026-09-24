# Group B — INC-2307 (Go)

You're the incident review team. Members have been paid two, three, even four times —
and our own dashboards say everything worked. Your job is **not** just to find the bug — it's
to find out **why it was possible, why it got to production, and why our alert didn't save us.**

**Needs:** Go 1.22+ (`go version`). No dependencies to download.

## Pick roles (1 min)

| Role | Who | Does |
|---|---|---|
| **Facilitator** | anyone | Keeps asking "why?", keeps time, stops blame |
| **Scribe** | anyone | Drives `go run ./cmd/whys` on the shared screen |
| **Code detectives** | devs | `internal/payout/`, `internal/provider/`, the replay |
| **Evidence detectives** | QEs | logs, PRs, backlog ticket, latency, API docs, test notes, alerting |
| **Sceptic** | anyone | Asks "is that a fact or a guess? show me" |

## The 22 minutes

| Time | Do this |
|---|---|
| **0–5** | Read `INCIDENT.md`. Run `go run ./cmd/replay`. Write a one-line **problem statement** (facts + numbers, no causes). |
| **5–17** | `go run ./cmd/whys`. Ask why, back each answer with evidence from this folder. **Branch** when you find more than one cause — there's more than one here. |
| **17–22** | Agree countermeasures for each root cause. Prevent > Detect > Mitigate. Prep a 3-minute playback from the generated `whys-report-*.md`. |

## Commands

```bash
go run ./cmd/replay                  # what happened in the 12:05 batch (~8s)
go run ./cmd/replay -member M-2090   # one member only
go run ./cmd/replay -log             # raw payout-service logs
go test ./...                        # existing tests (they pass… so why didn't they catch it?)
go run ./cmd/whys                    # the 5 Whys facilitator tool
```

## Questions to keep asking

- Why did the code do that? → Why was the code written/reviewed that way?
- Why didn't a **test** catch it? Why didn't the **alert** that fired lead anywhere?
- Why did it take a **member phoning us** to notice?
- Is this a fact we can point to, or an assumption?

**Stretch (if you finish early):** make `TestINC2307_MemberIsPaidExactlyOnce` in `internal/payout/incident_test.go` reproduce the bug — then fix it.
