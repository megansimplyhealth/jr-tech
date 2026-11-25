// Smoke test: is the deployed service alive and giving sensible answers?
// It checks one ordinary day, not every rule. That's the unit tests' job.
//
//	go run ./cmd/smoke staging
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	env := "staging"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	binary := ""
	for _, name := range []string{"slots", "slots.exe"} {
		p := filepath.Join(".environments", env, name)
		if _, err := os.Stat(p); err == nil {
			binary, _ = filepath.Abs(p)
		}
	}
	if binary == "" {
		fail("nothing deployed to %s", env)
	}

	health, err := exec.Command(binary, "-health").Output()
	if err != nil || strings.TrimSpace(string(health)) != "ok" {
		fail("%s health check: %v %s", env, err, health)
	}
	out, err := exec.Command(binary, "-date", "2026-06-01", "-clinic", "Leeds").Output()
	if err != nil {
		fail("%s: listing slots failed: %v", env, err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if lines[0] != "09:00–09:30 at Leeds" {
		fail("%s: first slot on 1 June should be 09:00–09:30 at Leeds, got %q", env, lines[0])
	}
	fmt.Printf("✓ smoke test on %s: healthy, %d slots on 1 June starting %s\n", env, len(lines), lines[0])
}

func fail(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "✗ smoke test: "+format+"\n", a...)
	os.Exit(1)
}
