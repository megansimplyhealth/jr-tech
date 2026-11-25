// Package slots works out which appointment times a BrightCare clinic offers.
package slots

import (
	"time"
	_ "time/tzdata" // so the UK timezone works on every machine, Windows included
)

// Clinic times are always UK local time.
var London = mustLoad("Europe/London")

func mustLoad(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return loc
}

const (
	Length      = 30 * time.Minute
	openHour    = 9
	lunchHour   = 12
	reopenHour  = 13
	closingHour = 17
)

type Slot struct {
	Clinic string
	Start  time.Time
	End    time.Time
}

// ForDay lists the bookable slots for one clinic on one day: 09:00 to 12:00
// and 13:00 to 17:00 UK time, every 30 minutes.
func ForDay(clinic string, day time.Time) []Slot {
	midnight := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, London)
	var out []Slot
	for t := midnight.Add(openHour * time.Hour); t.Before(midnight.Add(closingHour * time.Hour)); t = t.Add(Length) {
		if !t.Before(midnight.Add(lunchHour*time.Hour)) && t.Before(midnight.Add(reopenHour*time.Hour)) {
			continue
		}
		out = append(out, Slot{Clinic: clinic, Start: t, End: t.Add(Length)})
	}
	return out
}
