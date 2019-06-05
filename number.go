package unit

// Number represents percent and Chinese number unit
type Number float64

// Float64 implement the Unit interface
func (n Number) Float64() float64 {
	return float64(n)
}

// ...
const (
	// SI
	Permille        = Base * 1e-3
	Percent         = Base * 1e-2
	Base     Number = 1e0

	// CN
	Shi     = Base * 1e1
	Bai     = Base * 1e2
	Qian    = Base * 1e3
	Wan     = Base * 1e4
	ShiWan  = Base * 1e5
	BaiWan  = Base * 1e6
	QianWan = Base * 1e7
	Yi      = Base * 1e8
	ShiYi   = Base * 1e9
	BaiYi   = Base * 1e10
	QianYi  = Base * 1e11
	WanYi   = Base * 1e12
)

// Base returns the number in %
func (n Number) Base() float64 {
	return float64(n / Base)
}

// Percent returns the number in %
func (n Number) Percent() float64 {
	return float64(n / Percent)
}

// Permille returns the number in ‰
func (n Number) Permille() float64 {
	return float64(n / Permille)
}

// Shi returns the number in 十
func (n Number) Shi() float64 {
	return float64(n / Shi)
}

// Bai returns the number in 百
func (n Number) Bai() float64 {
	return float64(n / Bai)
}

// Qian returns the number in 千
func (n Number) Qian() float64 {
	return float64(n / Qian)
}

// Wan returns the number in 万
func (n Number) Wan() float64 {
	return float64(n / Wan)
}

// ShiWan returns the number in 十万
func (n Number) ShiWan() float64 {
	return float64(n / ShiWan)
}

// BaiWan returns the number in 百万
func (n Number) BaiWan() float64 {
	return float64(n / BaiWan)
}

// QianWan returns the number in 千万
func (n Number) QianWan() float64 {
	return float64(n / QianWan)
}

// Yi returns the number in 亿
func (n Number) Yi() float64 {
	return float64(n / Yi)
}

// ShiYi returns the number in 十亿
func (n Number) ShiYi() float64 {
	return float64(n / ShiYi)
}

// BaiYi returns the number in 百亿
func (n Number) BaiYi() float64 {
	return float64(n / BaiYi)
}

// QianYi returns the number in 千亿
func (n Number) QianYi() float64 {
	return float64(n / QianYi)
}

// WanYi returns the number in 万亿
func (n Number) WanYi() float64 {
	return float64(n / WanYi)
}
