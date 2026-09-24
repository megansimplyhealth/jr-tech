# Take-home: 5 Whys in practice

**Time:** about 2–3 hours, in your learning time, over the next 2 weeks
**Bring to:** your next 1:1 with your buddy or team lead
**Why bother?** Every part below produces **evidence for your competencies** — the table at the end tells you exactly which ones and gives you a sentence template, so you're not left wondering what counts.

---

## Part 1 — Swap incidents (60–90 min)

Pick up the **other group's** incident pack (TypeScript ↔ Go). Read `INCIDENT.md`, run the replay, then:

**Developer track**

1. **Reproduce it.** Finish the incident test (`test/incident.test.ts` or `internal/payout/incident_test.go`) so it **fails** for the same reason production failed. Go through the same path production does — raw request bodies / the real `NewClient` config.
2. **Fix the root cause, not the symptom.** "Wrap it in a try/catch" or "`Number()` it in `limits.ts`" fixes a symptom. Fix it where the problem _entered_ the system. Your test should now pass.
3. **Add a _detect_ control.** Something that would have told us within minutes if it happened again (a guard that throws, a log line an alert could query, an invariant check).
4. Write a short PR description: what broke, root cause, fix, how you tested it.

**QE track**

1. **Test gap analysis.** For each root cause in the pack, write down which test (unit, contract, integration, end-to-end, exploratory, non-functional) would have caught it, **at what point** in the delivery process, and **who** would own it.
2. **Reproduce it.** Write the failing incident test — pair with a dev if you want — or write a clear manual reproduction in a defect report (steps, expected, actual, evidence, impact).
3. **Test charter.** Write a 30-minute exploratory test charter for the fixed version: _"Explore [area] with [resources] to discover [information]."_ Include at least one risk you'd test that isn't in the incident.
4. **Test environment note.** Both incidents involve an environment that doesn't behave like production (hand-built fixtures; a sandbox that's always fast). Suggest one change to make our test environments more realistic.

---

## Part 2 — A real 5 Whys (45 min)

Pick a **real bug** from your squad in the last few months (ask your team lead if you can't think of one — it doesn't need to be big, and it doesn't need to be yours).

1. Run it through the tool: `npm run whys` or `go run ./cmd/whys`.
2. Aim for: **at least one branch**, **80%+ evidence-backed**, **zero blame**, and at least one **Prevent** and one **Detect** action.
3. Share the generated report with your buddy or team lead. If an action is worth doing, raise a ticket for it.

**Blameless means blameless.** If the chain leads to a person, ask _"what made it easy for a smart person to make that mistake?"_ and keep going.

---

## Part 3 — Explain it to a non-techie (15 min)

Write **5 sentences max** explaining one of the incidents to someone in Finance or the contact centre:
what happened, who was affected, why (in plain English), what we're doing so it doesn't happen again.
No jargon — if you write "idempotency" or "NaN", try again. 🙂

---

## Part 4 — Log your competency evidence (15 min)

Copy the rows that fit what you actually did into your competencies sheet. Mapped to the
[Government Digital and Data Profession Capability Framework](https://ddat-capability-framework.service.gov.uk/) —
if your sheet uses different headings, match on the description.

### Developers — Junior developer

| Skill (level expected)                               | What in this exercise shows it                                                     | Framework wording it matches                                                                                            |
| ---------------------------------------------------- | ---------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| **Programming and build** (Working)                  | Repro test + root cause fix + PR                                                   | "design, code, test, correct and document simple programs or scripts under the direction of others"                     |
| **Service support** (Working)                        | Running a structured RCA on a live incident / real bug                             | "help fix service faults following agreed procedures"                                                                   |
| **Systems integration** (Awareness)                  | Both incidents are failures _between_ systems (mobile ↔ engine, payouts ↔ PayFast) | "describe challenges of designing, building and testing interfaces between systems"                                     |
| **Modern development standards** (Working)           | Test-first reproduction; review/PR checklist improvements                          | "apply modern development standards principles under guidance"                                                          |
| **Information security** (Working)                   | Validating untrusted input at a service boundary (TS)                              | "use information security practices and available security controls to contribute to protecting solutions and services" |
| **Availability and capacity management** (Awareness) | Timeouts, provider latency, retries (Go)                                           | "explain availability and capacity management processes"                                                                |
| **Systems design** (Awareness)                       | Idempotency / contract design in your countermeasures                              | "assist as part of a team in the design of components of larger systems"                                                |
| **User focus** (Awareness)                           | Part 3 — impact on members, plain-English explanation                              | "show an awareness or understanding of user experience analysis and its principles"                                     |

### QEs — Quality assurance test analyst

| Skill                                                     | What in this exercise shows it                                                 | Framework wording it matches                                                                                     |
| --------------------------------------------------------- | ------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------- |
| **Test analysis**                                         | Test gap analysis — spotting risk in PRs, test plans, "out of scope" decisions | "explain what might be a risk in achieving quality goals"                                                        |
| **Designing and executing tests**                         | Repro test / defect reproduction + exploratory charter                         | "design, build and maintain simple tests that align to user needs and requirements"                              |
| **Managing, reporting and resolving defects**             | Defect report with evidence; retesting the fix                                 | "report, communicate and maintain defects with appropriate information… retest and escalate defects when needed" |
| **Test and quality planning**                             | Where in the process each test should happen, and who owns it                  | "explain the value of quality testing approaches, plans and strategies"                                          |
| **Test engineering**                                      | Test environment realism suggestion (fixtures, latency injection)              | "explain why testing processes, environments and tools are important"                                            |
| **Communicating between the technical and non-technical** | Part 3                                                                         | "explain why it's important to communicate technical concepts in non-technical language"                         |

### Evidence sentence template

> **Situation:** In the Jr 5 Whys session / on [bug ref], [one line on the problem and impact].
> **What I did:** I [reproduced it with a failing test / led the why tree / found the evidence that…].
> **Result:** [root cause found, fix merged, action raised as TICKET-123, report shared with…].
> **What I learned:** [one line].

Keep it to 3–4 lines per competency. A link to your PR, report or ticket is worth more than a paragraph.

---

### Checklist to bring to your 1:1

- [ ] Part 1 — PR / test / fix (dev) or gap analysis / defect report / charter (QE)
- [ ] Part 2 — a `whys-report-*.md` from a real bug, shared with someone
- [ ] Part 3 — 5-sentence plain-English summary
- [ ] Part 4 — at least 3 competency rows filled in with the template
