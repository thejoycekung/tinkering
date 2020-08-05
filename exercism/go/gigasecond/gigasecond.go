// Gigasecond helps calculate time after a gigasecond.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes a Time and adds a gigasecond to it, returning a Time.
func AddGigasecond(t time.Time) time.Time {
	var giga_str = "1000000000s"
	var giga, _ = time.ParseDuration(giga_str)
	t = t.Add(giga)
	return t
}
