/**
 * 5 Whys facilitator - an interactive terminal tool for running a blameless
 * root cause analysis as a group.
 *
 *   npm run whys
 *
 * It builds a "why tree" (you can branch when there's more than one cause),
 * nudges you when an answer blames a person or has no evidence, and writes a
 * markdown report you can paste into Confluence / the incident ticket.
 */
import { createInterface } from "node:readline";
import { writeFileSync } from "node:fs";

// ---------------------------------------------------------------- terminal io
const rl = createInterface({ input: process.stdin, terminal: false });
const lines = rl[Symbol.asyncIterator]();
let inputClosed = false;

const c = {
  bold: (s: string) => `\x1b[1m${s}\x1b[0m`,
  dim: (s: string) => `\x1b[2m${s}\x1b[0m`,
  yellow: (s: string) => `\x1b[33m${s}\x1b[0m`,
  green: (s: string) => `\x1b[32m${s}\x1b[0m`,
  cyan: (s: string) => `\x1b[36m${s}\x1b[0m`,
  red: (s: string) => `\x1b[31m${s}\x1b[0m`,
};

async function ask(prompt: string): Promise<string> {
  process.stdout.write(prompt);
  if (inputClosed) return "";
  const next = await lines.next();
  if (next.done) {
    inputClosed = true;
    process.stdout.write("\n");
    return "";
  }
  return next.value.trim();
}

async function askRequired(prompt: string): Promise<string> {
  for (;;) {
    const a = await ask(prompt);
    if (a || inputClosed) return a;
    console.log(c.dim("  (needs an answer)"));
  }
}

// ---------------------------------------------------------------- nudges
const BLAME_WORDS = [
  "forgot", "careless", "lazy", "stupid", "didn't bother", "didnt bother", "should have known",
  "human error", "user error", "incompetent", "sloppy", "their fault", "his fault", "her fault",
  "not paying attention", "wasn't careful", "wasnt careful", "mistake by",
];
const WEAK_ACTIONS = [
  "be more careful", "more careful", "remember to", "try harder", "pay more attention",
  "double check", "double-check", "be aware", "remind", "retrain", "tell people",
];

function blameNudge(text: string): string | undefined {
  const t = text.toLowerCase();
  const hit = BLAME_WORDS.find((w) => t.includes(w));
  if (hit) {
    return `"${hit}" sounds like it's about a person. Blame the process, not the people - ` +
      `what made it easy for this to happen, or hard to spot?`;
  }
  return undefined;
}

function weakActionNudge(text: string): string | undefined {
  const t = text.toLowerCase();
  const hit = WEAK_ACTIONS.find((w) => t.includes(w));
  if (hit) {
    return `"${hit}" relies on people remembering. Can you make it automatic, impossible, or detected instead?`;
  }
  return undefined;
}

// ---------------------------------------------------------------- the why tree
interface Node {
  id: number;
  parent?: Node;
  children: Node[];
  text: string;
  evidence: string; // "" = assumption
  depth: number; // 0 = problem statement
  rootCause: boolean;
}

interface Action {
  cause: Node;
  text: string;
  kind: "Prevent" | "Detect" | "Mitigate";
  owner: string;
}

const nodes: Node[] = [];
const actions: Action[] = [];
let blameNudges = 0;

function addNode(text: string, evidence: string, parent?: Node): Node {
  const n: Node = { id: nodes.length, parent, children: [], text, evidence, depth: parent ? parent.depth + 1 : 0, rootCause: false };
  parent?.children.push(n);
  nodes.push(n);
  return n;
}

const short = (s: string, n = 70) => (s.length > n ? s.slice(0, n - 1) + "…" : s);

/** Nodes in tree order (depth-first), so branches print under their parent. */
function inTreeOrder(): Node[] {
  const out: Node[] = [];
  const walk = (n: Node) => { out.push(n); n.children.forEach(walk); };
  if (nodes[0]) walk(nodes[0]);
  return out;
}

function printTree(): void {
  for (const n of inTreeOrder()) {
    const indent = "  ".repeat(n.depth);
    const tag = n.depth === 0 ? c.bold("PROBLEM") : c.cyan(`why ${n.depth}`);
    const flags = [n.depth > 0 && !n.evidence ? c.yellow("assumption") : "", n.rootCause ? c.green("ROOT CAUSE") : ""]
      .filter(Boolean).join(" ");
    console.log(`  [${n.id}] ${indent}${tag}: ${short(n.text)} ${flags}`);
  }
}

// ---------------------------------------------------------------- main flow
async function main() {
  console.log(c.bold("\n  5 WHYS - root cause facilitator\n"));
  console.log(c.dim("  Rules: facts over opinions · show your evidence · blame the process, not the people"));
  console.log(c.dim("  Tip: 5 is a guide, not a law. Stop when fixing the answer would stop it happening again.\n"));

  const group = (await ask("Group name: ")) || "group";
  const incident = (await ask("Incident ref (e.g. INC-2291): ")) || "incident";

  console.log(c.bold("\nStep 1 - Problem statement"));
  console.log(c.dim("  What happened, when, and what was the impact? Just the facts - no causes yet."));
  let problem = await askRequired("> ");
  for (let tries = 0; tries < 2 && !inputClosed; tries++) {
    const p = problem.toLowerCase();
    const warn: string[] = [];
    if (/\b(because|due to|caused by)\b/.test(p)) warn.push("it contains a cause ('because/due to') - save that for the whys");
    if (!/\d/.test(p)) warn.push("it has no numbers - when did it happen, how many customers, how much money?");
    if (!warn.length) break;
    console.log(c.yellow(`  Nudge: ${warn.join("; and ")}.`));
    const again = await ask("  Rewrite it (or press enter to keep): ");
    if (!again) break;
    problem = again;
  }
  const root = addNode(problem, "incident ticket");

  console.log(c.bold("\nStep 2 - Ask why"));
  console.log(c.dim("  After each answer: [enter] ask why again · [r] root cause found · [b] branch · [t] show tree\n"));

  let current: Node = root;
  let askAnotherReason = false;

  while (!inputClosed) {
    const whyNum = current.depth + 1;
    const q = askAnotherReason
      ? `What ELSE caused "${short(current.text, 50)}"?`
      : `Why #${whyNum}: why did "${short(current.text, 55)}" happen?`;
    askAnotherReason = false;
    console.log(c.bold(`\n${q}`));
    const answer = await askRequired("  because… ");
    if (inputClosed && !answer) break;

    const nudge = blameNudge(answer);
    let finalAnswer = answer;
    if (nudge) {
      blameNudges++;
      console.log(c.yellow(`  Nudge: ${nudge}`));
      const re = await ask("  Rephrase (or press enter to keep): ");
      if (re) finalAnswer = re;
    }

    const evidence = await ask("  Evidence - where did you see this? (file, log, test, commit, PR…) blank = it's a guess: ");
    if (!evidence) console.log(c.yellow("  Marked as an ASSUMPTION - try to verify it before you build on it."));

    current = addNode(finalAnswer, evidence, current);

    // what next?
    for (;;) {
      const choice = (await ask(c.dim("  [enter] why again · [r] root cause · [b] branch · [t] tree > "))).toLowerCase();
      if (choice === "t") { printTree(); continue; }
      if (choice === "" || choice === "w") break;
      if (choice === "r") {
        if (current.depth < 3) {
          console.log(c.yellow(`  Only ${current.depth} why(s) deep. Test: if you fixed JUST this, could the same problem happen again another way?`));
          const sure = (await ask("  Still a root cause? [y/N] ")).toLowerCase();
          if (sure !== "y") break;
        }
        if (!current.evidence) console.log(c.yellow("  Heads up: this root cause is still an assumption."));
        current.rootCause = true;
        console.log(c.green("  Root cause recorded."));
        const more = (await ask("  Other causes to explore? [b] branch · [enter] finish > ")).toLowerCase();
        if (more === "b") {
          const picked = await pickBranchPoint();
          if (picked) { current = picked; askAnotherReason = true; break; }
        }
        return finish(group, incident);
      }
      if (choice === "b") {
        const picked = await pickBranchPoint();
        if (picked) { current = picked; askAnotherReason = true; }
        break;
      }
    }
  }
  return finish(group, incident);
}

async function pickBranchPoint(): Promise<Node | undefined> {
  console.log(c.bold("\n  Branch from which statement? (things often have more than one cause)"));
  printTree();
  const pick = await ask("  Number: ");
  const n = nodes[Number(pick)];
  if (!pick || !n) { console.log(c.dim("  (no branch)")); return undefined; }
  return n;
}

async function finish(group: string, incident: string) {
  // make sure at least the deepest node on each leaf is considered
  let causes = nodes.filter((n) => n.rootCause);
  if (!causes.length) {
    causes = nodes.filter((n) => n.depth > 0 && n.children.length === 0);
    if (causes.length) console.log(c.yellow("\n  No root cause marked - using the deepest answer(s) on each branch."));
  }

  if (causes.length && !inputClosed) {
    console.log(c.bold("\nStep 3 - Countermeasures"));
    console.log(c.dim("  For each root cause: what will stop it happening again (or catch it sooner)?"));
  }
  for (const cause of causes) {
    if (inputClosed) break;
    console.log(c.bold(`\n  Root cause: ${cause.text}`));
    let text = await ask("  Action: ");
    if (!text) continue;
    const weak = weakActionNudge(text);
    if (weak) {
      console.log(c.yellow(`  Nudge: ${weak}`));
      const re = await ask("  Improve it (or press enter to keep): ");
      if (re) text = re;
    }
    const k = (await ask("  Does it [p]revent, [d]etect or [m]itigate? ")).toLowerCase();
    const kind = k.startsWith("d") ? "Detect" : k.startsWith("m") ? "Mitigate" : "Prevent";
    const owner = (await ask("  Owner (a role or team, not a name to blame): ")) || "TBC";
    actions.push({ cause, text, kind, owner });
  }

  // ---- scorecard
  const whys = nodes.filter((n) => n.depth > 0);
  const evidenced = whys.filter((n) => n.evidence).length;
  const maxDepth = Math.max(0, ...whys.map((n) => n.depth));
  const branches = nodes.filter((n) => n.children.length > 1).length;
  const score = {
    "Whys asked": whys.length,
    "Deepest chain": maxDepth,
    "Branch points": branches,
    "Root causes": causes.length,
    "Evidence-backed": whys.length ? `${Math.round((evidenced / whys.length) * 100)}%` : "-",
    "Blame nudges": blameNudges,
    "Actions": actions.length,
  };

  console.log(c.bold("\n  Scorecard"));
  for (const [k, v] of Object.entries(score)) console.log(`   ${k.padEnd(16)} ${v}`);

  // ---- markdown report
  const md: string[] = [];
  md.push(`# 5 Whys - ${incident} (${group})`, "", `_Generated ${new Date().toISOString()}_`, "");
  md.push("## Problem statement", "", nodes[0]?.text ?? "(none)", "");
  md.push("## Why tree", "");
  for (const n of inTreeOrder().slice(1)) {
    const flags = [n.rootCause ? "**ROOT CAUSE**" : "", n.evidence ? `_evidence: ${n.evidence}_` : "⚠️ _assumption_"].filter(Boolean).join(" · ");
    md.push(`${"  ".repeat(n.depth - 1)}- **Why ${n.depth}:** ${n.text} — ${flags}`);
  }
  md.push("", "## Countermeasures", "", "| Root cause | Action | Type | Owner |", "|---|---|---|---|");
  for (const a of actions) md.push(`| ${a.cause.text} | ${a.text} | ${a.kind} | ${a.owner} |`);
  md.push("", "## Scorecard", "", ...Object.entries(score).map(([k, v]) => `- ${k}: ${v}`), "");

  const stamp = new Date().toTimeString().slice(0, 8).replace(/:/g, "");
  const file = `whys-report-${group.replace(/[^a-z0-9]+/gi, "-").toLowerCase()}-${stamp}.md`;
  writeFileSync(file, md.join("\n"));
  console.log(c.green(`\n  Report written to ${file} - use it for your playback.\n`));
  rl.close();
}

main();
