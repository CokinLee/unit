package unit

// ElectricalConductance represents a SI unit of electrical conductance (in siemens, S)
type ElectricalConductance float64

// Float64 implement the Unit interface
func (e ElectricalConductance) Float64() float64 {
	return float64(e)
}

// ...
const (
	Siemens ElectricalConductance = 1e0 // SI
)

// Siemens returns the amount of substance in S
func (e ElectricalConductance) Siemens() float64 {
	return float64(e)
}
