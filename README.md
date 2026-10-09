# jr-tech

Hands-on sessions for the junior tech group — the material, code and take-home
exercises for each session, in one place.

## Sessions

| #   | Session                              | Topic                                          | Languages      | Length |
| --- | ------------------------------------ | ---------------------------------------------- | -------------- | ------ |
| 01  | [The 5 Whys](sessions/01-five-whys/) | Root cause analysis, blameless incident review | TypeScript, Go | 45 min |
| 02  | [Plan, Gate, Build](sessions/02-plan-gate-build/) | Working with AI: plan first, human review gate | Any (browser AI) | 45 min |

## How this repo works

- Each session lives in `sessions/NN-short-name/` with its own `README.md`.
- Group exercises are self-contained folders — no shared dependencies, so you can `cd` in and run.
- Take-home exercises are in each session's `TAKE-HOME.md` and map back to competencies
  (UK Government Digital and Data Profession Capability Framework).
- Starting a new session? Copy [`sessions/_template/`](sessions/_template/).

## Getting started

```bash
git clone <this repo>
cd jr-tech/sessions/01-five-whys
```

Then follow that session's README.

## Doing the take-home

Work on a branch so you don't step on each other:

```bash
git checkout -b takehome/01-five-whys/<your-name>
```

Open a PR when you're done if you'd like feedback — link it in your competencies sheet as evidence.
