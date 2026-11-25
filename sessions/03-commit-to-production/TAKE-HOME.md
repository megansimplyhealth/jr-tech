# Take-home: Commit to Production

**Time:** about 2 hours, in your learning time, over the next 2 weeks
**Bring to:** your next 1:1 with your buddy or team lead
**Why bother?** Every part below produces **evidence for your competencies**. The table at the end tells you which ones and gives you a sentence template.

---

## Part 1 — Swap services (30 min)

Do all three missions on the **other** group's service and workflow:
`s03-quote-calculator.yml` if you were in Group B, `s03-clinic-slots.yml` if you were in Group A.
You don't need to know the language well. The Actions logs tell you where to look.

Work on your own branch and **draft** PR, and don't merge it:

```bash
git checkout -b takehome/03-commit-to-production/<your-name>
```

**Developer track**

1. Finish all three missions, then do one stretch goal (the matrix, or explain why production
   is skipped on a PR).
2. Write the PR description: what was wrong with the workflow, what you changed, and what
   you changed in the code.

**QE track**

1. Before you fix anything in Mission 2, list every job in the workflow, whether it ran on your
   PR, and what kind of problem it would catch.
2. Find one kind of problem that **none** of the jobs would catch, and write a test for it.

---

## Part 2 — Read your squad's real workflows (45 min)

Find your service's workflows: `.github/workflows/` in its repo (ask someone in your squad if
it uses something other than GitHub Actions). For each workflow, write down:

- What starts it (`on:`), and on which branches and paths.
- Every job, what it checks, and what it `needs`.
- Anything that can be **skipped** (`if:`) or **allowed to fail** (`continue-on-error`). Why?
- Where secrets come in, and which `environment:` needs an approval.

**Developer track:** pick one job and explain every line of it in plain English.
**QE track:** mark where each kind of testing runs (unit, integration, end-to-end) and on which
event (every PR, merge to main, nightly). Note one gap or risk.

> Keep your notes in your team's space, not in a public tool.

---

## Part 3 — Write a workflow from scratch (30–45 min)

Do the free [Hello GitHub Actions](https://github.com/skills/hello-github-actions) course
first. Then, in a **personal** repo of your own, copy in one of the session services and write
its workflow yourself, without copying the session's:

1. Run on `pull_request` and `workflow_dispatch`.
2. A `test` job that checks out, sets up Node or Go **with caching**, and runs the tests.
3. A `strategy.matrix` that runs the tests on two versions.
4. A `build` job that `needs` the tests and uploads the build as an artifact.
5. A step that writes a short report to `$GITHUB_STEP_SUMMARY`.

Break it on purpose once, then fix it. Note one thing that surprised you.

---

## Part 4 — Log your competency evidence (15 min)

Copy the rows that fit what you actually did into your competencies sheet. Mapped to the
[Government Digital and Data Profession Capability Framework](https://understand-digital-data-roles-skills.service.gov.uk/).
If your sheet uses different headings, match on the description.

### Developers — Junior developer

| Skill (level expected)                     | What in this exercise shows it                                     | Framework wording it matches                                                                              |
| ------------------------------------------ | ------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------- |
| **Programming and build** (Working)        | Fixing a failing workflow from its logs, then the bug behind BC-5521 / BC-5530 | "design, code, test, correct and document simple programs or scripts under the direction of others"       |
| **Modern development standards** (Working) | Changing the workflow so tests run on every PR and block the build | "apply modern development standards principles under guidance"                                            |
| **Information security** (Working)         | Adding the secret scan, moving the key out of the code and passing it in with `${{ secrets }}` | "use information security practices and available security controls to contribute to protecting solutions and services" |
| **Service support** (Working)              | Finding and fixing a fault that reached members                    | "help fix service faults following agreed procedures"                                                     |

### QEs — Associate quality assurance test analyst

| Skill                             | What in this exercise shows it                                              | Framework wording it matches                                         |
| --------------------------------- | --------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| **Test engineering**              | Explaining what each job checks and when it runs; reviewing the PR          | "explain why testing processes, environments and tools are important" |
| **Designing and executing tests** | Part 1 — writing a test for a gap no check covered                          | "explain the value of automation within testing"                     |
| **Test analysis**                 | Part 2 — the risk or gap in your squad's workflows                         | "explain what might be a risk in achieving quality goals"            |
| **Test and quality planning**     | Mission 2 — showing how a PR went green while its tests failed or never ran | "explain how to measure the effectiveness of quality testing approaches, plans and strategies, and why it's important" |

### Evidence sentence template

> **Situation:** In the Jr Commit to Production session / on [task], [one line on what needed doing].
> **What I did:** I [read the Actions log and fixed… / changed the workflow so… / mapped where our tests run…].
> **Result:** [what changed, what it now catches, who I shared it with].
> **What I learned:** [one line].

Keep it to 3–4 lines per competency. A link to your PR, your notes or your workflow run is
worth more than a paragraph.

---

### Checklist to bring to your 1:1

- [ ] Part 1 — a link to your draft PR, plus its description or your list of jobs
- [ ] Part 2 — notes on your squad's real workflows with one gap or risk marked
- [ ] Part 3 — a link to a run of the workflow you wrote yourself
- [ ] Part 4 — at least 2 competency rows filled in with the template
