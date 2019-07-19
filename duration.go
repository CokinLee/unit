package unit

// Duration represents a SI unit of time (in seconds, s)
type Duration float64

// Float64 implement the Unit interface
func (d Duration) Float64() float64 {
	return float64(d)
}

// ...
const (
	// SI
	Yoctosecond          = Second * 1e-24
	Zeptosecond          = Second * 1e-21
	Attosecond           = Second * 1e-18
	Femtosecond          = Second * 1e-15
	Picosecond           = Second * 1e-12
	Nanosecond           = Second * 1e-9
	Microsecond          = Second * 1e-6
	Millisecond          = Second * 1e-3
	Centisecond          = Second * 1e-2
	Decisecond           = Second * 1e-1
	Second      Duration = 1e0
	Decasecond           = Second * 1e1
	Hectosecond          = Second * 1e2
	Kilosecond           = Second * 1e3
	Megasecond           = Second * 1e6
	Gigasecond           = Second * 1e9
	Terasecond           = Second * 1e12
	Petasecond           = Second * 1e15
	Exasecond            = Second * 1e18
	Zettasecond          = Second * 1e21
	Yottasecond          = Second * 1e24

	// non-SI
	Minute         = Second * 60
	Hour           = Minute * 60
	Day            = Hour * 24
	Week           = Day * 7
	ThirtyDayMonth = Day * 30
	JulianYear     = Day * 365.25
	Year360        = Day * 30 * 12
	Century        = Year360 * 100
)

// Yoctoseconds returns the time in ys
func (d Duration) Yoctoseconds() float64 {
	return float64(d / Yoctosecond)
}

// Zeptoseconds returns the time in zs
func (d Duration) Zeptoseconds() float64 {
	return float64(d / Zeptosecond)
}

// Attoseconds returns the time in as
func (d Duration) Attoseconds() float64 {
	return float64(d / Attosecond)
}

// Femtoseconds returns the time in fs
func (d Duration) Femtoseconds() float64 {
	return float64(d / Femtosecond)
}

// Picoseconds returns the time in ps
func (d Duration) Picoseconds() float64 {
	return float64(d / Picosecond)
}

// Nanoseconds returns the time in ns
func (d Duration) Nanoseconds() float64 {
	return float64(d / Nanosecond)
}

// Microseconds returns the time in µs
func (d Duration) Microseconds() float64 {
	return float64(d / Microsecond)
}

// Milliseconds returns the time in ms
func (d Duration) Milliseconds() float64 {
	return float64(d / Millisecond)
}

// Centiseconds returns the time in cs
func (d Duration) Centiseconds() float64 {
	return float64(d / Centisecond)
}

// Deciseconds returns the time in ds
func (d Duration) Deciseconds() float64 {
	return float64(d / Decisecond)
}

// Seconds returns the time in s
func (d Duration) Seconds() float64 {
	return float64(d / Second)
}

// Decaseconds returns the time in das
func (d Duration) Decaseconds() float64 {
	return float64(d / Decasecond)
}

// Hectoseconds returns the time in hs
func (d Duration) Hectoseconds() float64 {
	return float64(d / Hectosecond)
}

// Kiloseconds returns the time in ks
func (d Duration) Kiloseconds() float64 {
	return float64(d / Kilosecond)
}

// Megaseconds returns the time in Ms
func (d Duration) Megaseconds() float64 {
	return float64(d / Megasecond)
}

// Gigaseconds returns the time in Gs
func (d Duration) Gigaseconds() float64 {
	return float64(d / Gigasecond)
}

// Teraseconds returns the time in Ts
func (d Duration) Teraseconds() float64 {
	return float64(d / Terasecond)
}

// Petaseconds returns the time in Ps
func (d Duration) Petaseconds() float64 {
	return float64(d / Petasecond)
}

// Exaseconds returns the time in volt
func (d Duration) Exaseconds() float64 {
	return float64(d / Exasecond)
}

// Zettaseconds returns the time in Zs
func (d Duration) Zettaseconds() float64 {
	return float64(d / Zettasecond)
}

// Yottaseconds returns the time in Ys
func (d Duration) Yottaseconds() float64 {
	return float64(d / Yottasecond)
}

// Minutes returns the time in m
func (d Duration) Minutes() float64 {
	return float64(d / Minute)
}

// Hours returns the time in h
func (d Duration) Hours() float64 {
	return float64(d / Hour)
}

// Days returns the time in d
func (d Duration) Days() float64 {
	return float64(d / Day)
}

// Weeks returns the time in w
func (d Duration) Weeks() float64 {
	return float64(d / Week)
}

// ThirtyDayMonths returns the time in M
func (d Duration) ThirtyDayMonths() float64 {
	return float64(d / ThirtyDayMonth)
}

// JulianYears returns the time in Y
func (d Duration) JulianYears() float64 {
	return float64(d / JulianYear)
}
