# PayFast API v1 — POST /v1/payments (extract)

Creates a payment to the beneficiary's bank account.

### Timeouts

Payment submission is processed **synchronously**. If your client disconnects before receiving a
response, **the payment is still processed**. Do not assume a timed-out request has failed.

### Idempotency

Send an `Idempotency-Key` header (any unique string, max 64 chars) to safely retry requests.
If PayFast receives a second request with the same key, it returns the original payment and
**does not create a new one**.

> ⚠️ Requests **without** an `Idempotency-Key` are treated as new payments every time they are received.

### Errors

| Status | Meaning | Payment created? |
|---|---|---|
| 201 | Payment created | Yes |
| 400 / 422 | Invalid request | No |
| 503 | Temporarily unavailable | No |
| *(timeout / no response)* | Unknown — check with `GET /v1/payments?claimId=` | **Maybe** |
