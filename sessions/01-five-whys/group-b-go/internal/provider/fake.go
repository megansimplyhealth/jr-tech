// Package provider is a stand-in for PayFast, the third-party bank payments
// provider. It behaves like the real one: every request it receives is a
// payment, unless the request carries an Idempotency-Key it has already seen.
package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Payment is a payment the provider actually sent to a member's bank.
type Payment struct {
	PaymentID   string `json:"paymentId"`
	ClaimID     string `json:"claimId"`
	MemberID    string `json:"memberId"`
	AmountPence int    `json:"amountPence"`
}

// LatencyFunc decides how long the provider takes to answer the Nth request
// (1-based) it has received for a claim.
type LatencyFunc func(claimID string, n int) time.Duration

// Fake is an in-memory PayFast.
type Fake struct {
	Latency   LatencyFunc
	FailFirst int // respond 503 (without paying) to the first N requests per claim

	mu       sync.Mutex
	seen     map[string]int
	byKey    map[string]Payment
	ledger   []Payment
	sequence int
}

func NewFake(latency LatencyFunc) *Fake {
	if latency == nil {
		latency = func(string, int) time.Duration { return 50 * time.Millisecond }
	}
	return &Fake{Latency: latency, seen: map[string]int{}, byKey: map[string]Payment{}}
}

// Ledger returns every payment the provider has sent to a bank.
func (f *Fake) Ledger() []Payment {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]Payment(nil), f.ledger...)
}

func (f *Fake) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.URL.Path != "/v1/payments" {
		http.NotFound(w, r)
		return
	}
	var req struct {
		ClaimID     string `json:"claimId"`
		MemberID    string `json:"memberId"`
		AmountPence int    `json:"amountPence"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	f.mu.Lock()
	f.seen[req.ClaimID]++
	n := f.seen[req.ClaimID]
	f.mu.Unlock()

	// The provider does its work at its own pace. Note: it does NOT stop if
	// the caller hangs up - just like a real bank payment.
	time.Sleep(f.Latency(req.ClaimID, n))

	if n <= f.FailFirst {
		http.Error(w, "temporarily unavailable", http.StatusServiceUnavailable)
		return
	}

	f.mu.Lock()
	key := r.Header.Get("Idempotency-Key")
	p, dup := f.byKey[key]
	if key == "" || !dup {
		f.sequence++
		p = Payment{PaymentID: fmt.Sprintf("PF-%05d", 70000+f.sequence), ClaimID: req.ClaimID, MemberID: req.MemberID, AmountPence: req.AmountPence}
		f.ledger = append(f.ledger, p)
		if key != "" {
			f.byKey[key] = p
		}
	}
	f.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(p)
}
