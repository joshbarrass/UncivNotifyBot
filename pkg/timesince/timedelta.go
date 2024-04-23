package timesince

import "time"

func TimeDeltaSeconds(t1, t2 time.Time) float64 {
	return t1.Sub(t2).Seconds()
}

func AbsTimeDeltaSeconds(t1, t2 time.Time) float64 {
	return t1.Sub(t2).Abs().Seconds()
}
