package timesince

type formatSelection int

const (
	SECONDS formatSelection = iota
	MINUTES
	HOURS
	DAYS
)

const (
	FMT_SECONDS       = "%.1f seconds"
	FMT_SECONDS_SHORT = "%.1fs"
	FMT_MINUTES       = "%.1f minutes"
	FMT_MINUTES_SHORT = "%.1fm"
	FMT_HOURS         = "%.1f hours"
	FMT_HOURS_SHORT   = "%.1fh"
	FMT_DAYS          = "%.1f days"
	FMT_DAYS_SHORT    = "%.1fd"
)

var (
	FORMAT_MAP = map[formatSelection]string{
		SECONDS: FMT_SECONDS,
		MINUTES: FMT_MINUTES,
		HOURS:   FMT_HOURS,
		DAYS:    FMT_DAYS,
	}
	SHORT_FORMAT_MAP = map[formatSelection]string{
		SECONDS: FMT_SECONDS_SHORT,
		MINUTES: FMT_MINUTES_SHORT,
		HOURS:   FMT_HOURS_SHORT,
		DAYS:    FMT_DAYS_SHORT,
	}
)

const (
	FMT_AGO = "%s ago"
	FMT_IN  = "in %s"
)
