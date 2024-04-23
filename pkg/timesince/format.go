package timesince

import (
	"fmt"
	"math"
	"time"
)

// FormatTimeDelta returns a human-readable difference between two
// times. ref is the reference time, and comp is the comparison
// time. This has little effect other than to change whether the
// output is formatted as "... ago" (if ref > comp) or "In ..." (if
// ref < comp)
func FormatTimeDelta(ref, comp time.Time, short bool) string {
	delta := TimeDeltaSeconds(ref, comp)
	useAgo := delta >= 0
	delta = math.Abs(delta)

	var format formatSelection
	var formatDelta float64
	if delta < 60 {
		format = SECONDS
		formatDelta = delta
	} else if delta < 60*60 {
		format = MINUTES
		formatDelta = delta / 60
	} else if delta < 60*60*24 {
		format = HOURS
		formatDelta = delta / 60 / 60
	} else {
		format = DAYS
		formatDelta = delta / 60 / 60 / 24
	}

	var formatString string
	if short {
		formatString = SHORT_FORMAT_MAP[format]
	} else {
		formatString = FORMAT_MAP[format]
	}

	formatted := fmt.Sprintf(formatString, formatDelta)
	// adjust long formats to make them not plural if the delta ~= 1
	if !short && fmt.Sprintf("%.1f", formatDelta) == "1.0" {
		formatted = formatted[:len(formatted)-1]
	}

	if useAgo {
		return fmt.Sprintf(FMT_AGO, formatted)
	}
	return fmt.Sprintf(FMT_IN, formatted)
}
