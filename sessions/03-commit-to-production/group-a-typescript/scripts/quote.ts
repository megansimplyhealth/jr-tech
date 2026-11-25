// Try the service by hand:  npm run quote -- 1961-03-10 complete

import { PLANS, type Plan } from "../src/plans.ts";
import { createQuote } from "../src/quote.ts";

const [dob = "1990-05-01", plan = "plus"] = process.argv.slice(2);
if (!PLANS.includes(plan as Plan)) {
  console.error(`plan must be one of: ${PLANS.join(", ")}`);
  process.exit(2);
}
process.env.QUIET_LOGS = "1";
const q = createQuote({ memberName: "Example Member", dateOfBirth: dob, email: "member@example.com", plan: plan as Plan });
console.log(`${q.plan} plan, age ${q.age} (${q.ageBand} band): ${q.monthly} a month`);
