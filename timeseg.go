package costwhere

import "time"

type TimeSeg struct {
	Start time.Time
	End   time.Time
}

func newTimeSeg() TimeSeg {
	now := time.Now()
	return TimeSeg{
		Start: now,
		End:   now,
	}
}
func (t *TimeSeg) Cost() (d time.Duration) {
	d = t.End.Sub(t.Start)
	return
}
