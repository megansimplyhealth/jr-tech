package payout

import (
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"brightcare/payouts/internal/provider"
)

func testClient(url string) *Client {
	c := NewClient(url, log.New(io.Discard, "", 0))
	c.Backoff = 0
	return c
}

func TestPay_Success(t *testing.T) {
	fake := provider.NewFake(nil) // staging-like: answers in 50ms
	srv := httptest.NewServer(fake)
	defer srv.Close()

	r, err := testClient(srv.URL).Pay(Payout{ClaimID: "CLM-1", MemberID: "M-1", AmountPence: 4500})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.PaymentID == "" {
		t.Fatal("expected a payment id")
	}
	if got := len(fake.Ledger()); got != 1 {
		t.Fatalf("expected 1 payment, got %d", got)
	}
}

func TestPay_RetriesWhenProviderUnavailable(t *testing.T) {
	fake := provider.NewFake(nil)
	fake.FailFirst = 1 // first call gets a 503
	srv := httptest.NewServer(fake)
	defer srv.Close()

	r, err := testClient(srv.URL).Pay(Payout{ClaimID: "CLM-2", MemberID: "M-2", AmountPence: 6000})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Attempts != 2 {
		t.Fatalf("expected success on attempt 2, got %d", r.Attempts)
	}
	if got := len(fake.Ledger()); got != 1 {
		t.Fatalf("expected 1 payment, got %d", got)
	}
}

func TestPay_GivesUpAfterMaxAttempts(t *testing.T) {
	fake := provider.NewFake(nil)
	fake.FailFirst = 99
	srv := httptest.NewServer(fake)
	defer srv.Close()

	if _, err := testClient(srv.URL).Pay(Payout{ClaimID: "CLM-3", MemberID: "M-3", AmountPence: 1000}); err == nil {
		t.Fatal("expected an error")
	}
	if got := len(fake.Ledger()); got != 0 {
		t.Fatalf("expected no payments, got %d", got)
	}
}
