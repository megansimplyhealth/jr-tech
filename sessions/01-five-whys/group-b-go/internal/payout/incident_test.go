package payout

import "testing"

// INC-2307 - write a test that reproduces the incident.
//
// A good reproduction test:
//   - fails today, for the same reason production failed
//   - uses the same client config production uses (NewClient)
//   - will pass once the root cause is fixed, and stops it coming back
//
// Hint: provider.NewFake takes a LatencyFunc - what did PayFast's latency look
// like at 12:05 on 22 Sep? (see evidence/)
func TestINC2307_MemberIsPaidExactlyOnce(t *testing.T) {
	t.Skip("INC-2307: write me - remove this line when you're ready")

	// fake := provider.NewFake(func(claimID string, n int) time.Duration { ... })
	// srv := httptest.NewServer(fake)
	// ...
	// assert len(fake.Ledger()) == 1
}
