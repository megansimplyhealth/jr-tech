# 02 — Plan, Gate, Build

Get a better result out of any AI by making it plan first and be your sparing partner — and agreeing/adjusting that plan before it builds anything.

|            |                                                                                         |
| ---------- | --------------------------------------------------------------------------------------- |
| **Length** | 45 minutes                                                                              |
| **Who**    | Junior devs + QEs, in pairs or threes                                                   |
| **Slides** | [Plan, Gate, Build.pptx](Plan,%20Gate,%20Build.pptx)                                    |
| **Needs**  | A browser and any AI chat tool (Claude, ChatGPT, Gemini, Copilot…) · no cli or installs |

## Plan

| Time      | What                                                                           |
| --------- | ------------------------------------------------------------------------------ |
| 0–10 min  | Recap of Tech Day and Navigator, AI pros and cons, why we plan before we build |
| 10–35 min | Hands-on: plan, gate, build and review the ticket                              |
| 35–45 min | Show and reflect, take-home                                                    |

## Before you start

1. Read the ticket: [`TICKET.md`](TICKET.md).
2. **Read [`SUGGESTIONS.md`](SUGGESTIONS.md)** — it has small, realistic ideas for what you could make. You have about 8 minutes of build time, so pick something small.
3. Open whichever AI tool you like in your browser.
   - The **free tier is more than good enough**.
   - It **doesn't need a "plan mode"** — normal chat is fine. You make it plan with your prompt.
   - One screen per team, so you're all looking at the same conversation.

> **Don't paste anything from work into a public AI tool** — no customer or member data, no secrets, no internal code. Everything you need for this session is in this folder.

## The task (25 minutes)

### Step 1 — Spar for a plan (suggest 5/8 min)

This is the sparring part: you and the AI go back and forth until the plan is good. **No building yet.**

This is the prompt most people would send:

```text
Build something for this ticket: <paste the ticket>
```

Don't send it. Adjust it first so the AI goes into a planning and question mode and you get the best design for the ticket before anything is built. Things worth adding:

- Tell it to **plan only** and not to build anything until you say so.
- Ask it to **ask you questions first**, then answer them.
- Say **who it's for** and what **done** looks like.
- Give it your **constraints**: 8 minutes to build, runs in a browser chat, nothing to install.
- Ask for **two or three options with trade-offs**, then pick one.
- Ask for the final plan as a **short numbered list**.

Push back on it. If the plan is too big, too vague or not what you meant, say so and ask again.

### Step 2 — Human gate (suggest 5 min)

Stop and read the plan as a team. Nothing gets built until you all agree to it.

- Does it actually deliver the ticket?
- Can it be built in 8 minutes?
- What has the AI assumed that nobody told it?
- What would you cut?

Adjust the plan where it needs it — edit it yourselves or tell the AI what to change. Then post in the **Teams chat**:

```text
Team:
AI tool:

AI's plan:
<paste>

Our plan (only if you changed it):
<paste>
```

### Step 3 — Build (suggest 8/10 min)

Ask the AI to build the plan you agreed — and only that plan. Paste your final plan back in if you edited it outside the chat.

### Step 4 — Review the build (suggest 5/8 min)

Check it yourselves before you show anyone.

- Does it match the plan you agreed? What's missing or extra?
- Does it work? Open it, run it, read it properly.
- Is anything wrong, made up or out of date?
- Would you be happy to put your name on it?

## Show and reflect (10 min)

Each team shows:

- What did you make, and with which AI tool?
- What did you change at the human gates?
- Where did the AI help most?
- Where did it go rouge or get in the way?
- What would you do differently next time?

## After the session

[`TAKE-HOME.md`](TAKE-HOME.md) — run it again a different way, try it on something real, and log your evidence on your competency confluence page.

> The ticket in this session is fictional.
