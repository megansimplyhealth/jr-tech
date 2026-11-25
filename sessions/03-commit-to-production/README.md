# 03 — Commit to Production

Read, fix and extend real GitHub Actions workflows: why a run goes red, why a green one can
still be wrong, and how to add a gate that keeps secrets out of the code.

|            |                                                                                                                                   |
| ---------- | --------------------------------------------------------------------------------------------------------------------------------- |
| **Length** | 45 minutes                                                                                                                        |
| **Who**    | Junior devs + QEs, in 2 mixed groups                                                                                              |
| **Slides** | [Commit to Production.pptx](Commit%20to%20Production.pptx)                                                                        |
| **Needs**  | A GitHub account that can push branches to this repo · optional: Node 22.6+ (Group A) or Go 1.22+ (Group B) to run things locally |

## Plan

| Time      | What                                                                      |
| --------- | ------------------------------------------------------------------------- |
| 0–8 min   | Anatomy of a GitHub Actions workflow, and where to look when one fails    |
| 8–35 min  | Group exercise: open a draft PR, then three missions on a broken workflow |
| 35–42 min | Playback, 3 minutes per group                                             |
| 42–45 min | Wrap-up and take-home                                                     |

## The groups

| Folder                                       | Service                                          | Workflow                                                                       |
| -------------------------------------------- | ------------------------------------------------ | ------------------------------------------------------------------------------ |
| [`group-a-typescript/`](group-a-typescript/) | quote-calculator: prices a plan for a new member | [`s03-quote-calculator.yml`](../../.github/workflows/s03-quote-calculator.yml) |
| [`group-b-go/`](group-b-go/)                 | clinic-slots: lists bookable appointment times   | [`s03-clinic-slots.yml`](../../.github/workflows/s03-clinic-slots.yml)         |

Each group works on its own branch and **draft** PR, so every push runs the workflow on
GitHub for real. Nobody merges: the PRs are closed after the session, so the next group
starts from the same broken workflows.

## The three missions

| Mission                 | What happens                                                      | GitHub Actions you'll use                                    |
| ----------------------- | ----------------------------------------------------------------- | ------------------------------------------------------------ |
| **1. Make it run**      | The workflow fails, or doesn't run at all                         | `on:` triggers and `paths`, `actions/checkout`, reading logs |
| **2. Green, but wrong** | A bug reached members even though every check on the PR was green | `needs`, `if`, `continue-on-error`, skipped jobs             |
| **3. Add a gate**       | Add a secret-scanning job, then pass a real secret in safely      | adding a job, `${{ secrets.* }}`, `env`, the run summary     |

## Workflow cheat sheet

| Key                                     | What it does                                                                                                                |
| --------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `on:`                                   | What starts a run: `pull_request`, `push`, `workflow_dispatch` (a Run button), with optional `branches` and `paths` filters |
| `jobs:`                                 | The stages. Each job gets a fresh machine and runs in parallel unless it `needs` another                                    |
| `runs-on:`                              | Which kind of machine (runner) the job gets, e.g. `ubuntu-latest`                                                           |
| `needs:`                                | Jobs that must succeed before this one starts                                                                               |
| `if:`                                   | A condition. When it's false the job or step is **skipped**, and skipped still counts as fine                               |
| `steps:`                                | Run in order. `uses:` runs a ready-made action, `run:` runs a shell command                                                 |
| `actions/checkout`                      | Puts your repo on the runner. Without it, the job has no code                                                               |
| `continue-on-error:`                    | Lets a step or job fail without failing the run                                                                             |
| `environment:`                          | A named place you deploy to. It can hold its own secrets and require approval                                               |
| `${{ secrets.NAME }}`                   | A secret stored in the repo's settings. Hidden in logs                                                                      |
| `upload-artifact` / `download-artifact` | Pass the build's output from one job to the next                                                                            |
| `$GITHUB_STEP_SUMMARY`                  | A file: anything a step writes to it shows on the run's Summary page                                                        |

## After the session

[`TAKE-HOME.md`](TAKE-HOME.md) — swap services, read your squad's real workflows, build one in
your own repo, and log your evidence on your competency confluence page.

## Further reading

- [GitHub Skills — Hello GitHub Actions](https://github.com/skills/hello-github-actions): a free, hands-on course that runs in your own repo
- [GitHub Docs — Workflow syntax for GitHub Actions](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)
- [GitHub Docs — Using secrets in GitHub Actions](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions)
- [GitHub Docs — Using environments for deployment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment)
- [GitHub Docs — Storing workflow data as artifacts](https://docs.github.com/en/actions/using-workflows/storing-workflow-data-as-artifacts)
- [OWASP — Secrets Management Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Secrets_Management_Cheat_Sheet.html)

> The services, company, tickets and people in this session are fictional.
