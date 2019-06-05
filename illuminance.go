package unit

// Illuminance represents a SI unit for illuminance (in lux, lx)
type Illuminance float64

// Float64 implement the Unit interface
func (i Illuminance) Float64() float64 {
	return float64(i)
}

// constants
const (
	Lux Illuminance = 1e0 // SI
)

// Lux returns the illuminance in lx
func (i Illuminance) Lux() float64 {
	return float64(i)
}
