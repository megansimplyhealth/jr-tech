// "Deploys" the binary in dist/ by copying it into .environments/<env>/.
// On a real team this would push to a server or cloud service instead.
//
//	go run ./cmd/deploy staging
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const version = "2.3.0"

func main() {
	if len(os.Args) != 2 || (os.Args[1] != "staging" && os.Args[1] != "production") {
		fmt.Fprintln(os.Stderr, "usage: go run ./cmd/deploy staging|production")
		os.Exit(2)
	}
	env := os.Args[1]
	binary := ""
	for _, name := range []string{"slots", "slots.exe"} {
		if _, err := os.Stat(filepath.Join("dist", name)); err == nil {
			binary = name
		}
	}
	if binary == "" {
		fmt.Fprintln(os.Stderr, "✗ nothing to deploy: dist/ is empty. The build has to run first.")
		os.Exit(1)
	}

	target := filepath.Join(".environments", env)
	check(os.RemoveAll(target))
	check(os.MkdirAll(target, 0o755))
	data, err := os.ReadFile(filepath.Join("dist", binary))
	check(err)
	check(os.WriteFile(filepath.Join(target, binary), data, 0o755))

	run := os.Getenv("GITHUB_RUN_NUMBER")
	if run == "" {
		run = "local"
	}
	release := map[string]any{
		"service":     "clinic-slots",
		"version":     version + "-run" + run,
		"environment": env,
		"deployedAt":  time.Now().UTC(),
		"binary":      binary,
	}
	if by := os.Getenv("GITHUB_ACTOR"); by != "" {
		release["deployedBy"] = by
	}
	out, _ := json.MarshalIndent(release, "", "  ")
	check(os.WriteFile(filepath.Join(target, "release.json"), append(out, '\n'), 0o644))
	fmt.Printf("✓ deployed clinic-slots %s to %s\n", release["version"], env)

	notes := fmt.Sprintf("### Deployed clinic-slots `%s` to %s\n", release["version"], env)
	if os.Getenv("SMS_GATEWAY_TOKEN") == "" {
		fmt.Println("⚠ SMS_GATEWAY_TOKEN isn't set, so appointment reminders won't send in this environment.")
		notes += "\n⚠️ `SMS_GATEWAY_TOKEN` isn't set, so reminders won't send here.\n"
	}
	// On GitHub Actions, anything written to this file shows on the run's summary page.
	if summary := os.Getenv("GITHUB_STEP_SUMMARY"); summary != "" {
		f, err := os.OpenFile(summary, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		check(err)
		_, _ = f.WriteString(notes)
		f.Close()
	}
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "✗", err)
		os.Exit(1)
	}
}
