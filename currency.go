package unit

// Currency represents a unit of Chinese currency (元、角、分)
type Currency float64

// Float64 implement the Unit interface
func (c Currency) Float64() float64 {
	return float64(c)
}

// ...
const (
	// CN
	Yuan Currency = 1e0
	Jiao          = Yuan * 1e-1
	Fen           = Yuan * 1e-2
)

// Yuan returns the cunrrency in 元
func (c Currency) Yuan() float64 {
	return float64(c / Yuan)
}

// Jiao returns the cunrrency in 角
func (c Currency) Jiao() float64 {
	return float64(c / Jiao)
}

// Fen returns the cunrrency in 分
func (c Currency) Fen() float64 {
	return float64(c / Fen)
}
