package slots

import (
	"testing"
	"time"
)

func day(s string) time.Time {
	d, err := time.ParseInLocation("2006-01-02", s, London)
	if err != nil {
		panic(err)
	}
	return d
}

func TestOrdinaryDay(t *testing.T) {
	got := ForDay("Leeds", day("2026-06-01"))
	if len(got) != 14 {
		t.Fatalf("want 14 slots (6 before lunch, 8 after), got %d", len(got))
	}
	if first := got[0].Start.Format("15:04"); first != "09:00" {
		t.Errorf("first slot at %s, want 09:00", first)
	}
	if last := got[len(got)-1].Start.Format("15:04"); last != "16:30" {
		t.Errorf("last slot at %s, want 16:30", last)
	}
}

func TestNoSlotsOverLunch(t *testing.T) {
	for _, s := range ForDay("Leeds", day("2026-06-01")) {
		if h := s.Start.Hour(); h == 12 {
			t.Errorf("slot at %s is in the lunch break", s.Start.Format("15:04"))
		}
	}
}

func TestLabel(t *testing.T) {
	s := ForDay("Leeds", day("2026-06-01"))[0]
	if got, want := s.Label(), "09:00–09:30 at Leeds"; got != want {
		t.Errorf("label %q, want %q", got, want)
	}
}

// The clocks go forward on 29 March 2026 and back on 25 October 2026.
// Weekend clinics are open on both days.
func TestClocksChangeDays(t *testing.T) {
	for _, d := range []string{"2026-03-29", "2026-10-25"} {
		got := ForDay("Leeds", day(d))
		if first := got[0].Start.Format("15:04"); first != "09:00" {
			t.Errorf("%s: first slot at %s UK time, want 09:00", d, first)
		}
		if len(got) != 14 {
			t.Errorf("%s: want 14 slots, got %d", d, len(got))
		}
	}
}
