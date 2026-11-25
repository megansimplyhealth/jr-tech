// Structured logging. Anything that could identify a member is redacted
// before it's written, because these logs are shipped to a third-party tool.

const PERSONAL = ["memberName", "dateOfBirth", "email", "phone", "address"];

export function log(event: string, fields: Record<string, unknown> = {}): void {
  const safe: Record<string, unknown> = {};
  for (const [key, value] of Object.entries(fields)) {
    safe[key] = PERSONAL.includes(key) ? "[redacted]" : value;
  }
  if (process.env.QUIET_LOGS) return;
  process.stdout.write(JSON.stringify({ at: new Date().toISOString(), event, ...safe }) + "\n");
}
