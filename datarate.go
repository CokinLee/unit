package unit

// Datarate represents a unit of data rate (in bits per second, bit/s)
type Datarate float64

// Float64 implement the Unit interface
func (d Datarate) Float64() float64 {
	return float64(d)
}

// ...
const (
	// base 10 (SI prefixes)
	BitPerSecond      Datarate = 1e0
	KilobitPerSecond           = BitPerSecond * 1e3
	MegabitPerSecond           = BitPerSecond * 1e6
	GigabitPerSecond           = BitPerSecond * 1e9
	TerabitPerSecond           = BitPerSecond * 1e12
	PetabitPerSecond           = BitPerSecond * 1e15
	ExabitPerSecond            = BitPerSecond * 1e18
	ZettabitPerSecond          = BitPerSecond * 1e21
	YottabitPerSecond          = BitPerSecond * 1e24

	BytePerSecond      = BitPerSecond * 8
	KilobytePerSecond  = BytePerSecond * 1e3
	MegabytePerSecond  = BytePerSecond * 1e6
	GigabytePerSecond  = BytePerSecond * 1e9
	TerabytePerSecond  = BytePerSecond * 1e12
	PetabytePerSecond  = BytePerSecond * 1e15
	ExabytePerSecond   = BytePerSecond * 1e18
	ZettabytePerSecond = BytePerSecond * 1e21
	YottabytePerSecond = BytePerSecond * 1e24

	// base 2 (IEC prefixes)
	KibibitPerSecond = BitPerSecond * 1024
	MebibitPerSecond = KibibitPerSecond * 1024
	GibibitPerSecond = MebibitPerSecond * 1024
	TebibitPerSecond = GibibitPerSecond * 1024
	PebibitPerSecond = TebibitPerSecond * 1024
	ExbibitPerSecond = PebibitPerSecond * 1024
	ZebibitPerSecond = ExbibitPerSecond * 1024
	YobibitPerSecond = ZebibitPerSecond * 1024

	KibibytePerSecond = BytePerSecond * 1024
	MebibytePerSecond = KibibytePerSecond * 1024
	GibibytePerSecond = MebibytePerSecond * 1024
	TebibytePerSecond = GibibytePerSecond * 1024
	PebibytePerSecond = TebibytePerSecond * 1024
	ExbibytePerSecond = PebibytePerSecond * 1024
	ZebibytePerSecond = ExbibytePerSecond * 1024
	YobibytePerSecond = ZebibytePerSecond * 1024
)

// BitsPerSecond returns the data rate in bit/s
func (d Datarate) BitsPerSecond() float64 {
	return float64(d)
}

// KilobitsPerSecond returns the data rate in kbit/s
func (d Datarate) KilobitsPerSecond() float64 {
	return float64(d / KilobitPerSecond)
}

// MegabitsPerSecond returns the data rate in Mbit/s
func (d Datarate) MegabitsPerSecond() float64 {
	return float64(d / MegabitPerSecond)
}

// GigabitsPerSecond returns the data rate in Gbit/s
func (d Datarate) GigabitsPerSecond() float64 {
	return float64(d / GigabitPerSecond)
}

// TerabitsPerSecond returns the data rate in Tbit/s
func (d Datarate) TerabitsPerSecond() float64 {
	return float64(d / TerabitPerSecond)
}

// PetabitsPerSecond returns the data rate in Pbit/s
func (d Datarate) PetabitsPerSecond() float64 {
	return float64(d / PetabitPerSecond)
}

// ExabitsPerSecond returns the data rate in Ebit/s
func (d Datarate) ExabitsPerSecond() float64 {
	return float64(d / ExabitPerSecond)
}

// ZettabitsPerSecond returns the data rate in Zbit/s
func (d Datarate) ZettabitsPerSecond() float64 {
	return float64(d / ZettabitPerSecond)
}

// YottabitsPerSecond returns the data rate in Ybit/s
func (d Datarate) YottabitsPerSecond() float64 {
	return float64(d / YottabitPerSecond)
}

// BytesPerSecond returns the data rate in B/s
func (d Datarate) BytesPerSecond() float64 {
	return float64(d / BytePerSecond)
}

// KilobytesPerSecond returns the data rate in kB/s
func (d Datarate) KilobytesPerSecond() float64 {
	return float64(d / KilobytePerSecond)
}

// MegabytesPerSecond returns the data rate in MB/s
func (d Datarate) MegabytesPerSecond() float64 {
	return float64(d / MegabytePerSecond)
}

// GigabytesPerSecond returns the data rate in GB/s
func (d Datarate) GigabytesPerSecond() float64 {
	return float64(d / GigabytePerSecond)
}

// TerabytesPerSecond returns the data rate in TB/s
func (d Datarate) TerabytesPerSecond() float64 {
	return float64(d / TerabytePerSecond)
}

// PetabytesPerSecond returns the data rate in PB/s
func (d Datarate) PetabytesPerSecond() float64 {
	return float64(d / PetabytePerSecond)
}

// ExabytesPerSecond returns the data rate in EB/s
func (d Datarate) ExabytesPerSecond() float64 {
	return float64(d / ExabytePerSecond)
}

// ZettabytesPerSecond returns the data rate in ZB/s
func (d Datarate) ZettabytesPerSecond() float64 {
	return float64(d / ZettabytePerSecond)
}

// YottabytesPerSecond returns the data rate in YB/s
func (d Datarate) YottabytesPerSecond() float64 {
	return float64(d / YottabytePerSecond)
}

// KibibitsPerSecond returns the data rate in Kibit/s
func (d Datarate) KibibitsPerSecond() float64 {
	return float64(d / KibibitPerSecond)
}

// MebibitsPerSecond returns the data rate in Mibit/s
func (d Datarate) MebibitsPerSecond() float64 {
	return float64(d / MebibitPerSecond)
}

// GibibitsPerSecond returns the data rate in Gibit/s
func (d Datarate) GibibitsPerSecond() float64 {
	return float64(d / GibibitPerSecond)
}

// TebibitsPerSecond returns the data rate in Tibit/s
func (d Datarate) TebibitsPerSecond() float64 {
	return float64(d / TebibitPerSecond)
}

// PebibitsPerSecond returns the data rate in Pibit/s
func (d Datarate) PebibitsPerSecond() float64 {
	return float64(d / PebibitPerSecond)
}

// ExbibitsPerSecond returns the data rate in Eibit/s
func (d Datarate) ExbibitsPerSecond() float64 {
	return float64(d / ExbibitPerSecond)
}

// ZebibitsPerSecond returns the data rate in Zibit/s
func (d Datarate) ZebibitsPerSecond() float64 {
	return float64(d / ZebibitPerSecond)
}

// YobibitsPerSecond returns the data rate in Yibit/s
func (d Datarate) YobibitsPerSecond() float64 {
	return float64(d / YobibitPerSecond)
}

// KibibytesPerSecond returns the data rate in KiB/s
func (d Datarate) KibibytesPerSecond() float64 {
	return float64(d / KibibytePerSecond)
}

// MebibytesPerSecond returns the data rate in MiB/s
func (d Datarate) MebibytesPerSecond() float64 {
	return float64(d / MebibytePerSecond)
}

// GibibytesPerSecond returns the data rate in GiB/s
func (d Datarate) GibibytesPerSecond() float64 {
	return float64(d / GibibytePerSecond)
}

// TebibytesPerSecond returns the data rate in TiB/s
func (d Datarate) TebibytesPerSecond() float64 {
	return float64(d / TebibytePerSecond)
}

// PebibytesPerSecond returns the data rate in PiB/s
func (d Datarate) PebibytesPerSecond() float64 {
	return float64(d / PebibytePerSecond)
}

// ExbibytesPerSecond returns the data rate in EiB/s
func (d Datarate) ExbibytesPerSecond() float64 {
	return float64(d / ExbibytePerSecond)
}

// ZebibytesPerSecond returns the data rate in ZiB/s
func (d Datarate) ZebibytesPerSecond() float64 {
	return float64(d / ZebibytePerSecond)
}

// YobibytesPerSecond returns the data rate in YiB/s
func (d Datarate) YobibytesPerSecond() float64 {
	return float64(d / YobibytePerSecond)
}
