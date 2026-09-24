import { test } from "node:test";
import assert from "node:assert/strict";
import { ClaimsEngine } from "../src/engine.ts";

/**
 * INC-2291 - write a test that reproduces the incident.
 *
 * A good reproduction test:
 *  - fails today, for the same reason production failed
 *  - goes through the same path production does (hint: engine.handle takes a raw request body)
 *  - will pass once the root cause is fixed, and stop it coming back
 *
 * Remove the `{ todo: ... }` option when you're ready to run it.
 */
test("INC-2291: member cannot be paid more than their annual dental limit", { todo: "write me" }, () => {
  const engine = new ClaimsEngine(undefined, () => {});

  // TODO: send the same sequence of requests that member M-1042 sent on 16 Sep
  // engine.handle(JSON.stringify({ ... }));

  // TODO: assert the total paid never exceeds £150
  assert.ok(true);
});
