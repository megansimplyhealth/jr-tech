// Package payout sends approved claim reimbursements to members' bank
// accounts via the PayFast payments provider.
package payout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Payout is an approved reimbursement waiting to be paid.
type Payout struct {
	ClaimID     string `json:"claimId"`
	MemberID    string `json:"memberId"`
	AmountPence int    `json:"amountPence"`
}

// Receipt is what PayFast returns for a successful payment.
type Receipt struct {
	PaymentID string `json:"paymentId"`
	Attempts  int    `json:"-"`
}

const (
	// ProviderTimeout was 10s. Reduced in #530 so workers aren't tied up when
	// PayFast is slow (see INC-2250). Staging p95 is ~180ms so 2s is plenty.
	ProviderTimeout = 2 * time.Second

	// MaxAttempts - added in #498 (hotfix during INC-2250 provider outage).
	MaxAttempts = 3

	retryBackoff = 200 * time.Millisecond
)

// Client talks to PayFast.
type Client struct {
	BaseURL     string
	HTTP        *http.Client
	MaxAttempts int
	Backoff     time.Duration
	Log         *log.Logger
}

// NewClient returns a client configured the way production is.
func NewClient(baseURL string, logger *log.Logger) *Client {
	return &Client{
		BaseURL:     baseURL,
		HTTP:        &http.Client{Timeout: ProviderTimeout},
		MaxAttempts: MaxAttempts,
		Backoff:     retryBackoff,
		Log:         logger,
	}
}

// Pay sends one payout to PayFast, retrying if the call fails.
func (c *Client) Pay(p Payout) (Receipt, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return Receipt{}, err
	}

	var lastErr error
	for attempt := 1; attempt <= c.MaxAttempts; attempt++ {
		req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/v1/payments", bytes.NewReader(body))
		if err != nil {
			return Receipt{}, err
		}
		req.Header.Set("Content-Type", "application/json")
		// TODO(PAY-771): send an Idempotency-Key header

		start := time.Now()
		resp, err := c.HTTP.Do(req)
		if err != nil {
			lastErr = err
			c.Log.Printf("level=warn msg=\"payfast call failed\" claim=%s attempt=%d/%d took=%s err=%q",
				p.ClaimID, attempt, c.MaxAttempts, time.Since(start).Round(time.Millisecond), err)
			time.Sleep(c.Backoff)
			continue
		}

		receipt, retry, err := c.handle(resp)
		if err == nil {
			receipt.Attempts = attempt
			c.Log.Printf("level=info msg=\"payout sent\" claim=%s member=%s pence=%d payment=%s attempt=%d took=%s",
				p.ClaimID, p.MemberID, p.AmountPence, receipt.PaymentID, attempt, time.Since(start).Round(time.Millisecond))
			return receipt, nil
		}
		lastErr = err
		c.Log.Printf("level=warn msg=\"payfast error\" claim=%s attempt=%d/%d err=%q", p.ClaimID, attempt, c.MaxAttempts, err)
		if !retry {
			break
		}
		time.Sleep(c.Backoff)
	}

	c.Log.Printf("level=error msg=\"payout failed\" claim=%s member=%s err=%q action=\"queued for next batch\"",
		p.ClaimID, p.MemberID, lastErr)
	return Receipt{}, fmt.Errorf("payout %s failed after %d attempts: %w", p.ClaimID, c.MaxAttempts, lastErr)
}

// handle reads a PayFast response. retry reports whether the error is worth retrying.
func (c *Client) handle(resp *http.Response) (r Receipt, retry bool, err error) {
	defer resp.Body.Close()
	switch {
	case resp.StatusCode >= 500:
		msg, _ := io.ReadAll(resp.Body)
		return r, true, fmt.Errorf("payfast %d: %s", resp.StatusCode, bytes.TrimSpace(msg))
	case resp.StatusCode >= 400:
		msg, _ := io.ReadAll(resp.Body)
		return r, false, fmt.Errorf("payfast %d: %s", resp.StatusCode, bytes.TrimSpace(msg))
	}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return r, true, fmt.Errorf("decoding payfast response: %w", err)
	}
	return r, false, nil
}
