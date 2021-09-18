// Package lenconv performs Meters and Feet conversions.
package lenconv

import "fmt"

type Meters float64
type Feet float64

func (m Meters) String() string { return fmt.Sprintf("%g Meters", m) }
func (f Feet) String() string   { return fmt.Sprintf("%g Feet", f) }

// MToF converts Meters distance to Feet.
func MToF(m Meters) Feet { return Feet(m * 3.2808) }

// FToM converts Feet distance to Meters.
func FToM(f Feet) Meters { return Meters(f / 3.2808) }
