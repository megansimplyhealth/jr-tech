// Replays the 12:05 payout batch from Tue 22 Sep 2026 against a PayFast
// simulator that reproduces the provider's real latency that lunchtime.
//
//	go run ./cmd/replay                  # results table + provider settlement
//	go run ./cmd/replay -log             # raw payout-service logs
//	go run ./cmd/replay -member M-2044   # one member only
//
// Takes ~8 seconds - it uses the real production timeout.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"brightcare/payouts/internal/payout"
	"brightcare/payouts/internal/provider"
)

// PayFast response times observed for each claim in the batch, per request
// received (from PayFast's own request log - see evidence/payfast-latency.md).
var observedLatency = map[string][]time.Duration{
	"CLM-91001": {ms(310)},
	"CLM-91002": {ms(290)},
	"CLM-91003": {ms(2410), ms(940)},
	"CLM-91004": {ms(350)},
	"CLM-91005": {ms(2380), ms(2210), ms(820)},
	"CLM-91006": {ms(2450), ms(1100)},
	"CLM-91007": {ms(2520), ms(2470), ms(2390)},
	"CLM-91008": {ms(330)},
	"CLM-91009": {ms(2300), ms(760)},
	"CLM-91010": {ms(1850)},
	"CLM-91011": {ms(2610), ms(2440), ms(2350)},
	"CLM-91012": {ms(300)},
}

var batch = []payout.Payout{
	{ClaimID: "CLM-91001", MemberID: "M-2001", AmountPence: 4500},
	{ClaimID: "CLM-91002", MemberID: "M-2017", AmountPence: 12000},
	{ClaimID: "CLM-91003", MemberID: "M-2044", AmountPence: 6000},
	{ClaimID: "CLM-91004", MemberID: "M-2051", AmountPence: 3550},
	{ClaimID: "CLM-91005", MemberID: "M-2063", AmountPence: 15000},
	{ClaimID: "CLM-91006", MemberID: "M-2078", AmountPence: 8000},
	{ClaimID: "CLM-91007", MemberID: "M-2090", AmountPence: 20000},
	{ClaimID: "CLM-91008", MemberID: "M-2102", AmountPence: 2500},
	{ClaimID: "CLM-91009", MemberID: "M-2115", AmountPence: 9500},
	{ClaimID: "CLM-91010", MemberID: "M-2130", AmountPence: 7000},
	{ClaimID: "CLM-91011", MemberID: "M-2144", AmountPence: 11000},
	{ClaimID: "CLM-91012", MemberID: "M-2159", AmountPence: 4000},
}

func ms(n int) time.Duration { return time.Duration(n) * time.Millisecond }

// clockWriter stamps log lines as if it were 12:05 on the incident day.
type clockWriter struct {
	mu    sync.Mutex
	start time.Time
	buf   bytes.Buffer
}

func (w *clockWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	ts := time.Date(2026, 9, 22, 12, 5, 0, 0, time.UTC).Add(time.Since(w.start))
	fmt.Fprintf(&w.buf, "%s %s", ts.Format("2006-01-02T15:04:05.000Z"), p)
	return len(p), nil
}

func main() {
	showLog := flag.Bool("log", false, "print raw payout-service logs")
	member := flag.String("member", "", "only show this member")
	flag.Parse()

	fake := provider.NewFake(func(claimID string, n int) time.Duration {
		l := observedLatency[claimID]
		if n <= len(l) {
			return l[n-1]
		}
		return ms(300)
	})
	srv := httptest.NewServer(fake)
	defer srv.Close()

	out := &clockWriter{start: time.Now()}
	client := payout.NewClient(srv.URL, log.New(out, "", 0))

	fmt.Fprintln(os.Stderr, "Replaying 12:05 payout batch (12 payouts, real timeouts - ~8s)…")

	type result struct {
		p   payout.Payout
		r   payout.Receipt
		err error
	}
	results := make([]result, len(batch))
	var wg sync.WaitGroup
	for i, p := range batch {
		wg.Add(1)
		go func(i int, p payout.Payout) {
			defer wg.Done()
			r, err := client.Pay(p)
			results[i] = result{p, r, err}
		}(i, p)
	}
	wg.Wait()
	time.Sleep(3 * time.Second) // let PayFast finish requests the client gave up on

	if *showLog {
		logs := strings.ReplaceAll(out.buf.String(), srv.URL, "https://api.payfast.example")
		for _, line := range strings.Split(strings.TrimSpace(logs), "\n") {
			if *member == "" || strings.Contains(line, *member) || strings.Contains(line, claimFor(*member)) {
				fmt.Println(line)
			}
		}
		return
	}

	ledger := fake.Ledger()
	paidCount := map[string]int{}
	for _, p := range ledger {
		paidCount[p.ClaimID]++
	}

	fmt.Println("\nPayout service view - batch 12:05, Tue 22 Sep 2026")
	fmt.Println()
	fmt.Printf("%-11s %-8s %-9s %-10s %s\n", "claim", "member", "amount", "status", "attempts")
	fmt.Println(strings.Repeat("-", 52))
	for _, r := range results {
		if *member != "" && r.p.MemberID != *member {
			continue
		}
		status, attempts := "PAID", fmt.Sprint(r.r.Attempts)
		if r.err != nil {
			status, attempts = "FAILED", fmt.Sprint(payout.MaxAttempts)
		}
		fmt.Printf("%-11s %-8s %-9s %-10s %s\n", r.p.ClaimID, r.p.MemberID, pounds(r.p.AmountPence), status, attempts)
	}

	fmt.Println("\nPayFast settlement report - money that actually left the bank")
	fmt.Println()
	fmt.Printf("%-11s %-8s %-9s %-9s %s\n", "claim", "member", "amount", "payments", "total sent")
	fmt.Println(strings.Repeat("-", 56))
	sort.Slice(results, func(i, j int) bool { return results[i].p.ClaimID < results[j].p.ClaimID })
	var expected, sent int
	for _, r := range results {
		n := paidCount[r.p.ClaimID]
		expected += r.p.AmountPence
		sent += n * r.p.AmountPence
		if *member != "" && r.p.MemberID != *member {
			continue
		}
		flag := ""
		if n > 1 {
			flag = "  <-- DUPLICATE"
		}
		fmt.Printf("%-11s %-8s %-9s %-9d %s%s\n", r.p.ClaimID, r.p.MemberID, pounds(r.p.AmountPence), n, pounds(n*r.p.AmountPence), flag)
	}
	if *member == "" {
		fmt.Printf("\nExpected to pay %s. Actually paid %s. Over-paid %s.\n\n", pounds(expected), pounds(sent), pounds(sent-expected))
	}
}

func claimFor(member string) string {
	for _, p := range batch {
		if p.MemberID == member {
			return p.ClaimID
		}
	}
	return "\x00"
}

func pounds(pence int) string { return fmt.Sprintf("£%d.%02d", pence/100, pence%100) }
