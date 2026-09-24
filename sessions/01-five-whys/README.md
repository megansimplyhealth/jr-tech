# 01 — The 5 Whys

Finding the cause behind the cause — and why it's a process, not a person.

|            |                                                                              |
| ---------- | ---------------------------------------------------------------------------- |
| **Length** | 45 minutes                                                                   |
| **Who**    | Junior devs + QEs, in 2 mixed groups                                         |
| **Slides** | [The 5 Whys – Jr Session](https://claude.ai/artifact/PqDMuRNcview3XC26FkZwD) |
| **Needs**  | Group A: Node 22.6+ · Group B: Go 1.22+ · nothing to install                 |

## Plan

| Time      | What                                                      |
| --------- | --------------------------------------------------------- |
| 0–5 min   | What the 5 Whys is, the rules, two worked examples        |
| 5–10 min  | Split into groups, pick roles                             |
| 10–30 min | Group exercise: investigate an incident, build a why tree |
| 30–40 min | Playback — 3–4 min per group                              |
| 40–45 min | Wrap-up and take-home                                     |

## The groups

| Folder                                       | Incident                                     | Language   |
| -------------------------------------------- | -------------------------------------------- | ---------- |
| [`group-a-typescript/`](group-a-typescript/) | INC-2291 — claims paid over the annual limit | TypeScript |
| [`group-b-go/`](group-b-go/)                 | INC-2307 — members reimbursed 2–4 times      | Go         |

Each folder has the service code, a replay of the incident, the existing (passing) tests, a folder
of evidence, and an interactive 5 Whys tool that builds a why tree and writes a report.

## After the session

[`TAKE-HOME.md`](TAKE-HOME.md) — swap incidents, run a real 5 Whys, and log your evidence on your competency confluence page

## Further reading

- [Ministry of Testing — Two ways to use the 5 Whys method](https://www.ministryoftesting.com/insights/two-ways-to-use-the-5-whys-method-root-cause-of-bugs-and-identifying-continuous-improvements)
- [Atlassian — 5 Whys](https://www.atlassian.com/incident-management/postmortem/5-whys)
- [Miro — What is the 5 Whys framework](https://miro.com/root-cause-analysis/what-is-5-whys-framework/)
- [dev.to — Master the Five Whys method as a developer](https://dev.to/danilapetrova/master-the-five-whys-method-for-problem-solving-as-a-developer-3hi9)
- [Sithira Pathirana — Root cause analysis of software bugs using 5 Whys](https://www.linkedin.com/pulse/root-cause-analysis-software-bugs-using-5-whys-sithira-pathirana)
- [r/programming — Five Whys: Toyota's framework](https://www.reddit.com/r/programming/comments/1of9qf4/five_whys_toyotas_framework_for_finding_root/)

> The incidents, company and people in this session are fictional.
