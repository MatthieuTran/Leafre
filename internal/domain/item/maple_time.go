package item

import "time"

const (
	MAX_TIME  uint64 = 150842304000000000
	ZERO_TIME uint64 = 94354848000000000
	PERM_TIME uint64 = 150841440000000000
)

type MapleDateTime uint64

// From https://stackoverflow.com/a/59905670
func (dt MapleDateTime) ToTime() time.Time {
	const unixTimeBaseAsWin = 11644473600000000000 // The unix base time (January 1, 1970 UTC) as ns since Win32 epoch (1601-01-01)
	const nsToSecFactor = 1000000000

	timeToConvert := uint64(dt)

	unixsec := int64(timeToConvert-unixTimeBaseAsWin) / nsToSecFactor
	unixns := int64(timeToConvert % nsToSecFactor)

	time := time.Unix(unixsec, unixns)

	return time.Local()
}
