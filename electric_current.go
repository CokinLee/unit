package unit

// ElectricCurrent represents a SI unit of electric current (in ampere, A)
type ElectricCurrent float64

// Float64 implement the Unit interface
func (e ElectricCurrent) Float64() float64 {
	return float64(e)
}

// ...
const (
	// SI
	Yoctoampere                 = Ampere * 1e-24
	Zeptoampere                 = Ampere * 1e-21
	Attoampere                  = Ampere * 1e-18
	Femtoampere                 = Ampere * 1e-15
	Picoampere                  = Ampere * 1e-12
	Nanoampere                  = Ampere * 1e-9
	Microampere                 = Ampere * 1e-6
	Milliampere                 = Ampere * 1e-3
	Deciampere                  = Ampere * 1e-2
	Centiampere                 = Ampere * 1e-1
	Ampere      ElectricCurrent = 1e0
	Decaampere                  = Ampere * 1e1
	Hectoampere                 = Ampere * 1e2
	Kiloampere                  = Ampere * 1e3
	Megaampere                  = Ampere * 1e6
	Gigaampere                  = Ampere * 1e9
	Teraampere                  = Ampere * 1e12
	Petaampere                  = Ampere * 1e15
	Exaampere                   = Ampere * 1e18
	Zettaampere                 = Ampere * 1e21
	Yottaampere                 = Ampere * 1e24
)

// Yoctoamperes returns the electric current in yA
func (e ElectricCurrent) Yoctoamperes() float64 {
	return float64(e / Yoctoampere)
}

// Zeptoamperes returns the electric current in zA
func (e ElectricCurrent) Zeptoamperes() float64 {
	return float64(e / Zeptoampere)
}

// Attoamperes returns the electric current in aA
func (e ElectricCurrent) Attoamperes() float64 {
	return float64(e / Attoampere)
}

// Femtoamperes returns the electric current in fA
func (e ElectricCurrent) Femtoamperes() float64 {
	return float64(e / Femtoampere)
}

// Picoamperes returns the electric current in pA
func (e ElectricCurrent) Picoamperes() float64 {
	return float64(e / Picoampere)
}

// Nanoamperes returns the electric current in nA
func (e ElectricCurrent) Nanoamperes() float64 {
	return float64(e / Nanoampere)
}

// Microamperes returns the electric current in µA
func (e ElectricCurrent) Microamperes() float64 {
	return float64(e / Microampere)
}

// Milliamperes returns the electric current in mA
func (e ElectricCurrent) Milliamperes() float64 {
	return float64(e / Milliampere)
}

// Deciamperes returns the electric current in dA
func (e ElectricCurrent) Deciamperes() float64 {
	return float64(e / Deciampere)
}

// Centiamperes returns the electric current in cA
func (e ElectricCurrent) Centiamperes() float64 {
	return float64(e / Centiampere)
}

// Amperes returns the electric current in A
func (e ElectricCurrent) Amperes() float64 {
	return float64(e / Ampere)
}

// Decaamperes returns the electric current in daA
func (e ElectricCurrent) Decaamperes() float64 {
	return float64(e / Decaampere)
}

// Hectoamperes returns the electric current in hA
func (e ElectricCurrent) Hectoamperes() float64 {
	return float64(e / Hectoampere)
}

// Kiloamperes returns the electric current in kA
func (e ElectricCurrent) Kiloamperes() float64 {
	return float64(e / Kiloampere)
}

// Megaamperes returns the electric current in MA
func (e ElectricCurrent) Megaamperes() float64 {
	return float64(e / Megaampere)
}

// Gigaamperes returns the electric current in GA
func (e ElectricCurrent) Gigaamperes() float64 {
	return float64(e / Gigaampere)
}

// Teraamperes returns the electric current in TA
func (e ElectricCurrent) Teraamperes() float64 {
	return float64(e / Teraampere)
}

// Petaamperes returns the electric current in PA
func (e ElectricCurrent) Petaamperes() float64 {
	return float64(e / Petaampere)
}

// Exaamperes returns the electric current in EA
func (e ElectricCurrent) Exaamperes() float64 {
	return float64(e / Exaampere)
}

// Zettaamperes returns the electric current in ZA
func (e ElectricCurrent) Zettaamperes() float64 {
	return float64(e / Zettaampere)
}

// Yottaamperes returns the electric current in YA
func (e ElectricCurrent) Yottaamperes() float64 {
	return float64(e / Yottaampere)
}
