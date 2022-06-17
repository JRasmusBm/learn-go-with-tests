package timeutils

import (
	"fmt"
	"time"
)

type ClockTime struct {
	H int
	M int
	S int
}

func (c ClockTime) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", c.H, c.M, c.S)
}

func (ct ClockTime) ToTime() time.Time {
	return time.Date(312, time.October, 28, ct.H, ct.M, ct.S, 0, time.UTC)
}
