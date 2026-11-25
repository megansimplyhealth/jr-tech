# Group B — clinic-slots (Go)

You're the team that owns BrightCare's **clinic-slots**, the service that lists the appointment
times a clinic can offer members. It has a GitHub Actions workflow,
[`.github/workflows/s03-clinic-slots.yml`](../../../.github/workflows/s03-clinic-slots.yml).
Today that workflow is broken, and fixing it is only the start.

**Needs:** a GitHub account that can push branches to this repo. Go 1.22+ is handy for
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
git checkout -b s03/group-b-<your-team-name>
```

Add your names to a new file, `TEAM.md`, in this folder. Then push and open a **draft** PR:

```bash
git add TEAM.md && git commit -m "s03: group B team"
git push -u origin HEAD
gh pr create --draft --fill        # or open the PR on github.com
```

Open the PR on github.com and go to its **Checks** tab. That's where you'll spend the session.

## Mission 1 — Make it run

Look at your PR's checks. Your pipeline, **s03 clinic-slots**, isn't there.

- Open the workflow file. When does it run? When doesn't it?
- Change it so it runs on every pull request that changes **your service or the workflow
  itself**, and nothing else. Push, and watch it appear on the PR.

## Mission 2 — Green, but wrong

A ticket has just come in from the Leeds clinic:

> **BC-5530** · On Sunday 29 March the booking app didn't offer a 09:00 appointment, and
> the first slot was 10:00. On Sunday 25 October a member was booked in at **08:00**, an
> hour before we open. Every check on the PR was green.

- Look at the jobs in your run. Did every job run? What does the grey icon mean?
- **Why didn't the tests catch it?** Find the line in the workflow that let it through.
- Is there anything else in the workflow that means a build could go out without tests?
- Change the workflow so tests run on every PR and nothing builds until they pass.
  Push it and watch it go red. **Then** fix the bug in `internal/slots/` and get back to green.

## Mission 3 — Add a gate

The security team want every workflow to check for secrets before anything is built.
There's already a scanner: `go run ./cmd/secretscan`.

- Add a new `secret-scan` job to the workflow. Copy the shape of the `vet` job.
  Make `build` wait for it.
- Push. What did it find? Change the code so the value comes from an environment variable.
- Open the run's **Summary** page. The staging deploy says the token isn't set. The real token
  is stored as a repository secret called `S03_SMS_GATEWAY_TOKEN`. Pass it into the deploy step
  as the `SMS_GATEWAY_TOKEN` environment variable, and check the warning goes away.
- Get the PR green, then ask your PR reviewer to approve it.

**Stretch:**
- Why does **Deploy to production** always say *skipped* on your PR? Find the line.
- Add a `strategy.matrix` to the test job so it runs on Go 1.22 *and* Go 1.23.

## Useful commands

```bash
gh pr checks                         # the checks on your PR
gh run watch                         # follow the latest run live
gh run view --log-failed             # just the logs of the steps that failed
go test ./...                        # run a job's command locally to see more
go run ./cmd/secretscan
go run ./cmd/slots -date 2026-03-29  # try the service by hand
```

## Playback (3 minutes)

1. Show your workflow file as it was at the start, and as it is now (the PR's **Files** tab).
2. What was wrong in each mission, and where in the Actions UI did you spot it?
3. Which was worse: the missing run, or the green one?
