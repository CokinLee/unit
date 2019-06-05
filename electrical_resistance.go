package unit

// ElectricalResistance represents a SI derived unit of electrical resistance (in ohm, Ω)
type ElectricalResistance float64

// Float64 implement the Unit interface
func (e ElectricalResistance) Float64() float64 {
	return float64(e)
}

// ...
const (
	Ohm ElectricalResistance = 1e0 // SI
)

// Ohms returns the electrical resistance in Ω
func (e ElectricalResistance) Ohms() float64 {
	return float64(e)
}
