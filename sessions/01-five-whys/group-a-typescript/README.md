# Group A — INC-2291 (TypeScript)

You're the incident review team. Money has gone out the door and nobody knows why.
Your job is **not** just to find the bug — it's to find out **why the bug was possible,
why it got to production, and why it took 2 days to notice.**

**Needs:** Node 22.6+ (`node -v`). No `npm install` needed.

## Pick roles (1 min)

| Role | Who | Does |
|---|---|---|
| **Facilitator** | anyone | Keeps asking "why?", keeps time, stops blame |
| **Scribe** | anyone | Drives `npm run whys` on the shared screen |
| **Code detectives** | devs | `src/`, `test/`, `npm run replay` |
| **Evidence detectives** | QEs | logs, PR, Slack, test plan, alerting |
| **Sceptic** | anyone | Asks "is that a fact or a guess? show me" |

## The 22 minutes

| Time | Do this |
|---|---|
| **0–5** | Read `INCIDENT.md`. Run `npm run replay`. Write a one-line **problem statement** (facts + numbers, no causes). |
| **5–17** | `npm run whys`. Ask why, back each answer with evidence from this folder. **Branch** when you find more than one cause — there's more than one here. |
| **17–22** | Agree countermeasures for each root cause. Prevent > Detect > Mitigate. Prep a 3-minute playback from the generated `whys-report-*.md`. |

## Commands

```bash
npm run replay                     # what happened on 16 Sep
npm run replay -- --member M-1042  # one member only
npm run replay -- --log            # raw engine logs
npm test                           # existing tests (they pass… so why didn't they catch it?)
npm run whys                       # the 5 Whys facilitator tool
```

## Questions to keep asking

- Why did the code do that? → Why was the code written/reviewed that way?
- Why didn't a **test** catch it? Why didn't **monitoring** catch it?
- Why did it take **2 days** and a **finance team** to notice?
- Is this a fact we can point to, or an assumption?

**Stretch (if you finish early):** make the `INC-2291` test in `test/incident.test.ts` reproduce the bug — then fix it.
