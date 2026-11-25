// Looks for passwords, keys and tokens that have been written into the code.
// Secrets belong in environment variables or a secrets store, never in git.
//
//	go run ./cmd/secretscan
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var patterns = []struct {
	name string
	re   *regexp.Regexp
}{
	{"BrightCare live API key", regexp.MustCompile(`bc_live_[A-Za-z0-9]{16,}`)},
	{"private key", regexp.MustCompile("-----BEGIN " + "[A-Z ]*PRIVATE KEY-----")},
	{"hard-coded password", regexp.MustCompile(`(?i)password\s*[:=]\s*["'][^"'\s]{6,}["']`)},
}

var skipDirs = map[string]bool{"dist": true, ".environments": true, ".git": true}
var textExt = map[string]bool{".go": true, ".mod": true, ".json": true, ".yml": true, ".yaml": true, ".md": true, ".env": true, ".txt": true, ".sh": true}

func main() {
	found, scanned := 0, 0
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if skipDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if !textExt[filepath.Ext(path)] && !strings.HasPrefix(d.Name(), ".env") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		scanned++
		for i, line := range strings.Split(string(data), "\n") {
			for _, p := range patterns {
				if m := p.re.FindString(line); m != "" {
					found++
					fmt.Fprintf(os.Stderr, "%s:%d  %s  %s…\n", filepath.ToSlash(path), i+1, p.name, m[:min(12, len(m))])
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "✗", err)
		os.Exit(1)
	}
	if found > 0 {
		fmt.Fprintf(os.Stderr, "\n✗ %d possible secret(s) found in %d files.\n", found, scanned)
		fmt.Fprintln(os.Stderr, "  Read it from an environment variable instead (os.Getenv(\"NAME\")).")
		fmt.Fprintln(os.Stderr, "  Anything that was ever committed is in git history, so it has to be rotated too.")
		os.Exit(1)
	}
	fmt.Printf("✓ no secrets found in %d files\n", scanned)
}
