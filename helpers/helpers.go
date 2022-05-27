package helpers

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"time"
)

// DurationToString converts provided duration to human readable string presentation.
func DurationToString(d time.Duration) string {
	ns := time.Duration(d.Nanoseconds())
	ms := float64(ns) / 1000000.0
	var unit string
	var amount string
	if ns < time.Microsecond {
		unit = "ns"
		amount, unit = humanize.CommafWithDigits(float64(ns), 0), "ns"
	} else if ns < time.Millisecond {
		unit = "μs"
		amount, unit = humanize.CommafWithDigits(ms*1000.0, 3), "μs"
	} else if ns < time.Second {
		unit = "ms"
		amount, unit = humanize.CommafWithDigits(ms, 3), "ms"
	} else if ns < time.Minute {
		unit = "s"
		amount, unit = humanize.CommafWithDigits(ms/1000.0, 3), "s"
	} else if ns < time.Hour {
		unit = "m"
		amount, unit = humanize.CommafWithDigits(ms/60000.0, 3), "m"
	} else if ns < 24*time.Hour {
		unit = "h"
		amount, unit = humanize.CommafWithDigits(ms/3600000.0, 3), "h"
	} else {
		days := ms / 86400000.0
		unit = "day"
		if days > 1 {
			unit = "days"
		}
		amount = humanize.CommafWithDigits(days, 3)
	}
	return fmt.Sprintf("%s %s", amount, unit)
}
