package unit

// Datasize represents a unit of data size (in bits, bit)
type Datasize float64

// Float64 implement the Unit interface
func (d Datasize) Float64() float64 {
	return float64(d)
}

// ...
const (
	// base 10 (SI prefixes)
	Bit      Datasize = 1e0
	Kilobit           = Bit * 1e3
	Megabit           = Bit * 1e6
	Gigabit           = Bit * 1e9
	Terabit           = Bit * 1e12
	Petabit           = Bit * 1e15
	Exabit            = Bit * 1e18
	Zettabit          = Bit * 1e21
	Yottabit          = Bit * 1e24

	Byte      = Bit * 8
	Kilobyte  = Byte * 1e3
	Megabyte  = Byte * 1e6
	Gigabyte  = Byte * 1e9
	Terabyte  = Byte * 1e12
	Petabyte  = Byte * 1e15
	Exabyte   = Byte * 1e18
	Zettabyte = Byte * 1e21
	Yottabyte = Byte * 1e24

	// base 2 (IEC prefixes)
	Kibibit = Bit * 1024
	Mebibit = Kibibit * 1024
	Gibibit = Mebibit * 1024
	Tebibit = Gibibit * 1024
	Pebibit = Tebibit * 1024
	Exbibit = Pebibit * 1024
	Zebibit = Exbibit * 1024
	Yobibit = Zebibit * 1024

	Kibibyte = Byte * 1024
	Mebibyte = Kibibyte * 1024
	Gibibyte = Mebibyte * 1024
	Tebibyte = Gibibyte * 1024
	Pebibyte = Tebibyte * 1024
	Exbibyte = Pebibyte * 1024
	Zebibyte = Exbibyte * 1024
	Yobibyte = Zebibyte * 1024
)

// Bits returns the datasize in bit
func (d Datasize) Bits() float64 {
	return float64(d)
}

// Kilobits returns the datasize in kbit
func (d Datasize) Kilobits() float64 {
	return float64(d / Kilobit)
}

// Megabits returns the datasize in Mbit
func (d Datasize) Megabits() float64 {
	return float64(d / Megabit)
}

// Gigabits returns the datasize in Gbit
func (d Datasize) Gigabits() float64 {
	return float64(d / Gigabit)
}

// Terabits returns the datasize in Tbit
func (d Datasize) Terabits() float64 {
	return float64(d / Terabit)
}

// Petabits returns the datasize in Pbit
func (d Datasize) Petabits() float64 {
	return float64(d / Petabit)
}

// Exabits returns the datasize in Ebit
func (d Datasize) Exabits() float64 {
	return float64(d / Exabit)
}

// Zettabits returns the datasize in Zbit
func (d Datasize) Zettabits() float64 {
	return float64(d / Zettabit)
}

// Yottabits returns the datasize in Ybit
func (d Datasize) Yottabits() float64 {
	return float64(d / Yottabit)
}

// Bytes returns the datasize in B
func (d Datasize) Bytes() float64 {
	return float64(d / Byte)
}

// Kilobytes returns the datasize in kB
func (d Datasize) Kilobytes() float64 {
	return float64(d / Kilobyte)
}

// Megabytes returns the datasize in MB
func (d Datasize) Megabytes() float64 {
	return float64(d / Megabyte)
}

// Gigabytes returns the datasize in GB
func (d Datasize) Gigabytes() float64 {
	return float64(d / Gigabyte)
}

// Terabytes returns the datasize in TB
func (d Datasize) Terabytes() float64 {
	return float64(d / Terabyte)
}

// Petabytes returns the datasize in PB
func (d Datasize) Petabytes() float64 {
	return float64(d / Petabyte)
}

// Exabytes returns the datasize in EB
func (d Datasize) Exabytes() float64 {
	return float64(d / Exabyte)
}

// Zettabytes returns the datasize in ZB
func (d Datasize) Zettabytes() float64 {
	return float64(d / Zettabyte)
}

// Yottabytes returns the datasize in YB
func (d Datasize) Yottabytes() float64 {
	return float64(d / Yottabyte)
}

// Kibibits returns the datasize in Kibit
func (d Datasize) Kibibits() float64 {
	return float64(d / Kibibit)
}

// Mebibits returns the datasize in Mibit
func (d Datasize) Mebibits() float64 {
	return float64(d / Mebibit)
}

// Gibibits returns the datasize in Gibit
func (d Datasize) Gibibits() float64 {
	return float64(d / Gibibit)
}

// Tebibits returns the datasize in Tibit
func (d Datasize) Tebibits() float64 {
	return float64(d / Tebibit)
}

// Pebibits returns the datasize in Pibit
func (d Datasize) Pebibits() float64 {
	return float64(d / Pebibit)
}

// Exbibits returns the datasize in Eibit
func (d Datasize) Exbibits() float64 {
	return float64(d / Exbibit)
}

// Zebibits returns the datasize in Zibit
func (d Datasize) Zebibits() float64 {
	return float64(d / Zebibit)
}

// Yobibits returns the datasize in Yibit
func (d Datasize) Yobibits() float64 {
	return float64(d / Yobibit)
}

// Kibibytes returns the datasize in KiB
func (d Datasize) Kibibytes() float64 {
	return float64(d / Kibibyte)
}

// Mebibytes returns the datasize in MiB
func (d Datasize) Mebibytes() float64 {
	return float64(d / Mebibyte)
}

// Gibibytes returns the datasize in GiB
func (d Datasize) Gibibytes() float64 {
	return float64(d / Gibibyte)
}

// Tebibytes returns the datasize in TiB
func (d Datasize) Tebibytes() float64 {
	return float64(d / Tebibyte)
}

// Pebibytes returns the datasize in PiB
func (d Datasize) Pebibytes() float64 {
	return float64(d / Pebibyte)
}

// Exbibytes returns the datasize in EiB
func (d Datasize) Exbibytes() float64 {
	return float64(d / Exbibyte)
}

// Zebibytes returns the datasize in ZiB
func (d Datasize) Zebibytes() float64 {
	return float64(d / Zebibyte)
}

// Yobibytes returns the datasize in YiB
func (d Datasize) Yobibytes() float64 {
	return float64(d / Yobibyte)
}
