// Package lengthconv converts units of distance measurement.
package lengthconv

import "fmt"

// Feet is 12 inches!
type Feet float64

// Meters is 100 centimeters!
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
