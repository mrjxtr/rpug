package layout

import (
	"strconv"
	"time"
)

// copyrightStartYear is the year rpug was first published. Leave it alone, it's a
// historical fact, not something that should track the clock.
const copyrightStartYear = 2025

// copyrightYears builds the year text for the footer notice. Returns just the start
// year while we're still in it, otherwise a "2025-2026" range so the notice covers
// every year the project has been revised. Uses server local time, which is UTC on
// Fly.io, so the rollover happens at UTC midnight.
func copyrightYears() string {
	start := strconv.Itoa(copyrightStartYear)

	current := time.Now().Year()
	if current <= copyrightStartYear {
		return start
	}

	return start + "-" + strconv.Itoa(current)
}
