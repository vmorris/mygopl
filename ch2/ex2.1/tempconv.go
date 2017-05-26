// Package lengthconv converts distance measurement units
package lengthconv

import "fmt"

// Feet is a feet :P
type Feet float64

// Meters is a meter :P
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
