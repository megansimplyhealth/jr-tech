package slots

import "fmt"

// Label is what members see in the booking app, e.g. "09:00–09:30 at Leeds".
func (s Slot) Label() string {
	return fmt.Sprintf("%s–%s at %s", s.Start.Format("15:04"), s.End.Format("15:04"), s.Clinic)
}
