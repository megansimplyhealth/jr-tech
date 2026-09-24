// 5 Whys facilitator - an interactive terminal tool for running a blameless
// root cause analysis as a group.
//
//	go run ./cmd/whys
//
// It builds a "why tree" (you can branch when there's more than one cause),
// nudges you when an answer blames a person or has no evidence, and writes a
// markdown report you can paste into Confluence / the incident ticket.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// ------------------------------------------------------------------ terminal io

var (
	in          = bufio.NewScanner(os.Stdin)
	inputClosed bool
)

func bold(s string) string   { return "\x1b[1m" + s + "\x1b[0m" }
func dim(s string) string    { return "\x1b[2m" + s + "\x1b[0m" }
func yellow(s string) string { return "\x1b[33m" + s + "\x1b[0m" }
func green(s string) string  { return "\x1b[32m" + s + "\x1b[0m" }
func cyan(s string) string   { return "\x1b[36m" + s + "\x1b[0m" }

func ask(prompt string) string {
	fmt.Print(prompt)
	if inputClosed {
		return ""
	}
	if !in.Scan() {
		inputClosed = true
		fmt.Println()
		return ""
	}
	return strings.TrimSpace(in.Text())
}

func askRequired(prompt string) string {
	for {
		a := ask(prompt)
		if a != "" || inputClosed {
			return a
		}
		fmt.Println(dim("  (needs an answer)"))
	}
}

// ------------------------------------------------------------------ nudges

var blameWords = []string{
	"forgot", "careless", "lazy", "stupid", "didn't bother", "didnt bother", "should have known",
	"human error", "user error", "incompetent", "sloppy", "their fault", "his fault", "her fault",
	"not paying attention", "wasn't careful", "wasnt careful", "mistake by",
}

var weakActions = []string{
	"be more careful", "more careful", "remember to", "try harder", "pay more attention",
	"double check", "double-check", "be aware", "remind", "retrain", "tell people",
}

func firstHit(text string, words []string) string {
	t := strings.ToLower(text)
	for _, w := range words {
		if strings.Contains(t, w) {
			return w
		}
	}
	return ""
}

// ------------------------------------------------------------------ the why tree

type node struct {
	id        int
	parent    *node
	children  []*node
	text      string
	evidence  string // "" = assumption
	depth     int    // 0 = problem statement
	rootCause bool
}

type action struct {
	cause *node
	text  string
	kind  string
	owner string
}

var (
	nodes       []*node
	actions     []action
	blameNudges int
)

func addNode(text, evidence string, parent *node) *node {
	n := &node{id: len(nodes), parent: parent, text: text, evidence: evidence}
	if parent != nil {
		n.depth = parent.depth + 1
		parent.children = append(parent.children, n)
	}
	nodes = append(nodes, n)
	return n
}

func short(s string, n int) string {
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	return string([]rune(s)[:n-1]) + "…"
}

// inTreeOrder returns nodes depth-first, so branches print under their parent.
func inTreeOrder() []*node {
	var out []*node
	var walk func(*node)
	walk = func(n *node) {
		out = append(out, n)
		for _, c := range n.children {
			walk(c)
		}
	}
	if len(nodes) > 0 {
		walk(nodes[0])
	}
	return out
}

func printTree() {
	for _, n := range inTreeOrder() {
		tag := cyan(fmt.Sprintf("why %d", n.depth))
		if n.depth == 0 {
			tag = bold("PROBLEM")
		}
		var flags []string
		if n.depth > 0 && n.evidence == "" {
			flags = append(flags, yellow("assumption"))
		}
		if n.rootCause {
			flags = append(flags, green("ROOT CAUSE"))
		}
		fmt.Printf("  [%d] %s%s: %s %s\n", n.id, strings.Repeat("  ", n.depth), tag, short(n.text, 70), strings.Join(flags, " "))
	}
}

func pickBranchPoint() *node {
	fmt.Println(bold("\n  Branch from which statement? (things often have more than one cause)"))
	printTree()
	pick := ask("  Number: ")
	i, err := strconv.Atoi(pick)
	if err != nil || i < 0 || i >= len(nodes) {
		fmt.Println(dim("  (no branch)"))
		return nil
	}
	return nodes[i]
}

// ------------------------------------------------------------------ main flow

var causeWords = regexp.MustCompile(`\b(because|due to|caused by)\b`)
var hasDigit = regexp.MustCompile(`\d`)

func main() {
	fmt.Println(bold("\n  5 WHYS - root cause facilitator\n"))
	fmt.Println(dim("  Rules: facts over opinions · show your evidence · blame the process, not the people"))
	fmt.Println(dim("  Tip: 5 is a guide, not a law. Stop when fixing the answer would stop it happening again.\n"))

	group := ask("Group name: ")
	if group == "" {
		group = "group"
	}
	incident := ask("Incident ref (e.g. INC-2307): ")
	if incident == "" {
		incident = "incident"
	}

	fmt.Println(bold("\nStep 1 - Problem statement"))
	fmt.Println(dim("  What happened, when, and what was the impact? Just the facts - no causes yet."))
	problem := askRequired("> ")
	for tries := 0; tries < 2 && !inputClosed; tries++ {
		p := strings.ToLower(problem)
		var warn []string
		if causeWords.MatchString(p) {
			warn = append(warn, "it contains a cause ('because/due to') - save that for the whys")
		}
		if !hasDigit.MatchString(p) {
			warn = append(warn, "it has no numbers - when did it happen, how many customers, how much money?")
		}
		if len(warn) == 0 {
			break
		}
		fmt.Println(yellow("  Nudge: " + strings.Join(warn, "; and ") + "."))
		again := ask("  Rewrite it (or press enter to keep): ")
		if again == "" {
			break
		}
		problem = again
	}
	root := addNode(problem, "incident ticket", nil)

	fmt.Println(bold("\nStep 2 - Ask why"))
	fmt.Println(dim("  After each answer: [enter] ask why again · [r] root cause found · [b] branch · [t] show tree\n"))

	current := root
	askAnother := false

loop:
	for !inputClosed {
		q := fmt.Sprintf("Why #%d: why did %q happen?", current.depth+1, short(current.text, 55))
		if askAnother {
			q = fmt.Sprintf("What ELSE caused %q?", short(current.text, 50))
		}
		askAnother = false
		fmt.Println(bold("\n" + q))
		answer := askRequired("  because… ")
		if inputClosed && answer == "" {
			break
		}

		if hit := firstHit(answer, blameWords); hit != "" {
			blameNudges++
			fmt.Println(yellow(fmt.Sprintf("  Nudge: %q sounds like it's about a person. Blame the process, not the people - what made it easy for this to happen, or hard to spot?", hit)))
			if re := ask("  Rephrase (or press enter to keep): "); re != "" {
				answer = re
			}
		}

		evidence := ask("  Evidence - where did you see this? (file, log, test, commit, PR…) blank = it's a guess: ")
		if evidence == "" {
			fmt.Println(yellow("  Marked as an ASSUMPTION - try to verify it before you build on it."))
		}
		current = addNode(answer, evidence, current)

		for {
			choice := strings.ToLower(ask(dim("  [enter] why again · [r] root cause · [b] branch · [t] tree > ")))
			switch choice {
			case "t":
				printTree()
				continue
			case "", "w":
				continue loop
			case "r":
				if current.depth < 3 {
					fmt.Println(yellow(fmt.Sprintf("  Only %d why(s) deep. Test: if you fixed JUST this, could the same problem happen again another way?", current.depth)))
					if strings.ToLower(ask("  Still a root cause? [y/N] ")) != "y" {
						continue loop
					}
				}
				if current.evidence == "" {
					fmt.Println(yellow("  Heads up: this root cause is still an assumption."))
				}
				current.rootCause = true
				fmt.Println(green("  Root cause recorded."))
				if strings.ToLower(ask("  Other causes to explore? [b] branch · [enter] finish > ")) == "b" {
					if picked := pickBranchPoint(); picked != nil {
						current, askAnother = picked, true
						continue loop
					}
				}
				break loop
			case "b":
				if picked := pickBranchPoint(); picked != nil {
					current, askAnother = picked, true
				}
				continue loop
			}
		}
	}
	finish(group, incident)
}

func finish(group, incident string) {
	var causes []*node
	for _, n := range nodes {
		if n.rootCause {
			causes = append(causes, n)
		}
	}
	if len(causes) == 0 {
		for _, n := range nodes {
			if n.depth > 0 && len(n.children) == 0 {
				causes = append(causes, n)
			}
		}
		if len(causes) > 0 {
			fmt.Println(yellow("\n  No root cause marked - using the deepest answer(s) on each branch."))
		}
	}

	if len(causes) > 0 && !inputClosed {
		fmt.Println(bold("\nStep 3 - Countermeasures"))
		fmt.Println(dim("  For each root cause: what will stop it happening again (or catch it sooner)?"))
	}
	for _, cause := range causes {
		if inputClosed {
			break
		}
		fmt.Println(bold("\n  Root cause: " + cause.text))
		text := ask("  Action: ")
		if text == "" {
			continue
		}
		if hit := firstHit(text, weakActions); hit != "" {
			fmt.Println(yellow(fmt.Sprintf("  Nudge: %q relies on people remembering. Can you make it automatic, impossible, or detected instead?", hit)))
			if re := ask("  Improve it (or press enter to keep): "); re != "" {
				text = re
			}
		}
		k := strings.ToLower(ask("  Does it [p]revent, [d]etect or [m]itigate? "))
		kind := "Prevent"
		if strings.HasPrefix(k, "d") {
			kind = "Detect"
		} else if strings.HasPrefix(k, "m") {
			kind = "Mitigate"
		}
		owner := ask("  Owner (a role or team, not a name to blame): ")
		if owner == "" {
			owner = "TBC"
		}
		actions = append(actions, action{cause, text, kind, owner})
	}

	// ---- scorecard
	whys := inTreeOrder()[1:]
	evidenced, maxDepth, branches := 0, 0, 0
	for _, n := range whys {
		if n.evidence != "" {
			evidenced++
		}
		if n.depth > maxDepth {
			maxDepth = n.depth
		}
	}
	for _, n := range nodes {
		if len(n.children) > 1 {
			branches++
		}
	}
	evPct := "-"
	if len(whys) > 0 {
		evPct = fmt.Sprintf("%d%%", evidenced*100/len(whys))
	}
	score := [][2]string{
		{"Whys asked", fmt.Sprint(len(whys))},
		{"Deepest chain", fmt.Sprint(maxDepth)},
		{"Branch points", fmt.Sprint(branches)},
		{"Root causes", fmt.Sprint(len(causes))},
		{"Evidence-backed", evPct},
		{"Blame nudges", fmt.Sprint(blameNudges)},
		{"Actions", fmt.Sprint(len(actions))},
	}
	fmt.Println(bold("\n  Scorecard"))
	for _, s := range score {
		fmt.Printf("   %-16s %s\n", s[0], s[1])
	}

	// ---- markdown report
	var md strings.Builder
	fmt.Fprintf(&md, "# 5 Whys - %s (%s)\n\n_Generated %s_\n\n", incident, group, time.Now().UTC().Format(time.RFC3339))
	fmt.Fprintf(&md, "## Problem statement\n\n%s\n\n## Why tree\n\n", nodes[0].text)
	for _, n := range whys {
		var flags []string
		if n.rootCause {
			flags = append(flags, "**ROOT CAUSE**")
		}
		if n.evidence != "" {
			flags = append(flags, "_evidence: "+n.evidence+"_")
		} else {
			flags = append(flags, "⚠️ _assumption_")
		}
		fmt.Fprintf(&md, "%s- **Why %d:** %s — %s\n", strings.Repeat("  ", n.depth-1), n.depth, n.text, strings.Join(flags, " · "))
	}
	md.WriteString("\n## Countermeasures\n\n| Root cause | Action | Type | Owner |\n|---|---|---|---|\n")
	for _, a := range actions {
		fmt.Fprintf(&md, "| %s | %s | %s | %s |\n", a.cause.text, a.text, a.kind, a.owner)
	}
	md.WriteString("\n## Scorecard\n\n")
	for _, s := range score {
		fmt.Fprintf(&md, "- %s: %s\n", s[0], s[1])
	}

	slug := strings.Trim(regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(strings.ToLower(group), "-"), "-")
	file := fmt.Sprintf("whys-report-%s-%s.md", slug, time.Now().Format("150405"))
	if err := os.WriteFile(file, []byte(md.String()), 0o644); err != nil {
		fmt.Println(yellow("  Could not write report: " + err.Error()))
		fmt.Println(md.String())
		return
	}
	fmt.Println(green(fmt.Sprintf("\n  Report written to %s - use it for your playback.\n", file)))
}
