# Group A — quote-calculator (TypeScript)

You're the team that owns BrightCare's **quote-calculator**, the service that prices a cash
plan for a new member. Every pull request runs its GitHub Actions workflow,
[`.github/workflows/s03-quote-calculator.yml`](../../../.github/workflows/s03-quote-calculator.yml).
Today that workflow is broken, and fixing it is only the start.

**Needs:** a GitHub account that can push branches to this repo. Node 22.6+ is handy for
running commands locally, but everything important happens on GitHub.

> **Don't merge your PR.** Leave it open as a draft. The facilitator closes it after the session.

## Pick roles (1 min)

| Role            | Who    | Does                                                                                     |
| --------------- | ------ | ---------------------------------------------------------------------------------------- |
| **Driver**      | anyone | Edits files, commits, pushes                                                             |
| **Log reader**  | dev    | Has the PR's **Checks** tab open. Reads out which job and step failed, and why           |
| **PR reviewer** | QE     | Reviews the PR on GitHub. **Only approves if they can say which checks ran and passed**  |
| **Sceptic**     | anyone | Asks "it's green, but did it actually check anything?"                                   |

## Get set up (3 min)

```bash
git checkout main && git pull
git checkout -b s03/group-a-<your-team-name>
```

Add your names to a new file, `TEAM.md`, in this folder. Then push and open a **draft** PR:

```bash
git add TEAM.md && git commit -m "s03: group A team"
git push -u origin HEAD
gh pr create --draft --fill        # or open the PR on github.com
```

Open the PR on github.com and go to its **Checks** tab. That's where you'll spend the session.

## Mission 1 — Make it run

Your pipeline, **s03 quote-calculator**, has gone red.

- Which job failed? Which step? Open the log and read the error properly.
- Why does this job fail when the jobs after it would be fine? Compare it with them.
- Fix the **workflow file**, push, and watch it run again.

## Mission 2 — Green, but wrong

A ticket has just come in from the contact centre:

> **BC-5521** · A member who turned 65 last week was quoted the **40–64** price for the
> complete plan. Our price list says 65 and over should be in the 65+ band. The change went
> out last Tuesday. Every check on the PR was green.

- Open the **Unit tests** job. Did it pass? Look at every step, not just the tick.
- **Why was the PR green?** Find the line in the workflow that let it through.
- Is there anything else in the workflow that means a build could go out without tests?
- Change the workflow so a failing test fails the PR and nothing builds until tests pass.
  Push it and watch it go red. **Then** fix the bug in `src/` and get back to green.

## Mission 3 — Add a gate

The security team want every workflow to check for secrets before anything is built.
There's already a scanner: `npm run secret-scan`.

- Add a new `secret-scan` job to the workflow. Copy the shape of the `lint` job.
  Make `build` wait for it.
- Push. What did it find? Change the code so the value comes from an environment variable.
- Open the run's **Summary** page. The staging deploy says the key isn't set. The real key is
  stored as a repository secret called `S03_PAYMENTS_API_KEY`. Pass it into the deploy step
  as the `PAYMENTS_API_KEY` environment variable, and check the warning goes away.
- Get the PR green, then ask your PR reviewer to approve it.

**Stretch:**
- Why does **Deploy to production** always say *skipped* on your PR? Find the line.
- Add a `strategy.matrix` to the test job so it runs on Node 22 *and* Node 24.

## Useful commands

```bash
gh pr checks                       # the checks on your PR
gh run watch                       # follow the latest run live
gh run view --log-failed           # just the logs of the steps that failed
npm test                           # run a job's command locally to see more
npm run secret-scan
npm run quote -- 1961-03-10 complete   # try the service by hand
```

## Playback (3 minutes)

1. Show your workflow file as it was at the start, and as it is now (the PR's **Files** tab).
2. What was wrong in each mission, and where in the Actions UI did you spot it?
3. Which was worse: the red run, or the green one?
