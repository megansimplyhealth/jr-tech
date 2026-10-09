# What could we make?

Read this before you start [the ticket](TICKET.md). You have about **8 minutes of build time**, so the best ideas are small: one problem, one screen or one page.

These are suggestions, not a list to choose from. If you have a better idea, go with it — you just need to justify it at the human gate.

## A small app

Ask for **one self-contained HTML file** (HTML, CSS and JavaScript together). Most AI tools will preview it in the chat; if yours doesn't, save it as `index.html` and open it in your browser.

| Idea                   | Smallest useful version                                          |
| ---------------------- | ---------------------------------------------------------------- |
| First-week checklist   | A tick list grouped by day, with a progress bar                  |
| Jargon buster          | Search box over 15–20 acronyms and what they mean                |
| "Who do I ask?" finder | Pick a topic, get the role or team to go to                      |
| Day-one quiz           | Five multiple-choice questions with a score at the end           |
| Setup tracker          | List of tools and access to request, each marked done or waiting |

## A document

Good if you want to spend your time on the content rather than the code.

| Idea                      | Smallest useful version                                              |
| ------------------------- | -------------------------------------------------------------------- |
| First-week guide          | One page: day-by-day, what to do and who to meet                     |
| "Things I wish I'd known" | Ten short tips from your own first weeks                             |
| Buddy guide               | One page for the buddy: what to cover on each of the first five days |
| First ticket walkthrough  | From picking up a ticket to a merged PR, in ten steps                |
| Welcome FAQ               | The ten questions every new starter asks, with answers               |

## A small API

You won't have time to host anything, so aim for the **design plus the code** — something a teammate could read and run later.

| Idea                | Smallest useful version                                                   |
| ------------------- | ------------------------------------------------------------------------- |
| New starter FAQ API | `GET /faqs` and `GET /faqs?topic=`, with 8–10 made-up FAQs in a JSON file |
| Glossary API        | `GET /terms/{term}` returning a definition and an example                 |
| Checklist API       | `GET /checklist/{day}` returning that day's tasks                         |

Ask for the endpoints, example requests and responses, and the code in one file in a language you can read.

## Something else

A decision tree, a one-page cheat sheet, a welcome message template, a set of flashcards — anything a new starter could use on day one.

## Keeping it small

- Solve **one** problem. If your plan has "and" in it, cut half.
- Use made-up content. Nothing real from work.
- No logins, databases or hosting.
- If the AI's plan has more than about five steps, it's too big for 8 minutes — ask it to cut it down.
