# QE test plan — Mobile app v2.0 (new claims journey)

**Owner:** Mobile QE  **Status:** Signed off Fri 11 Sep

## In scope

| Area | Approach | Result |
|---|---|---|
| Claim form UI (iOS / Android) | Manual exploratory + Appium smoke | ✅ Pass |
| Field validation (amount, date, receipt upload) | Automated UI tests | ✅ Pass |
| Accessibility (screen reader, font scaling) | Manual | ✅ Pass (2 minor bugs raised) |
| API request is sent with correct fields | Mocked API — assert request body matches `claim-v2.json` | ✅ Pass |

## Out of scope

- **Claims engine decisioning / limits** — owned and tested by the claims squad
- **End-to-end with real claims-engine** — staging claims-engine not yet wired to mobile v2 (ticket MOB-91, backlog)
- Performance

## Risks

- ⚠️ No end-to-end test of mobile → claims-engine before go-live. Accepted by product owner because
  "the payload is the same as web".
