// Package tempconv performs Celsius, Fahrenheit and Kelvin conversions.
package tempconv

import "fmt"

// Celsius temp scale type
type Celsius float64

// Fahrenheit temp scale type
type Fahrenheit float64

// Kelvin temp scale type
type Kelvin float64

const (
	// AbsoluteZeroC is pretty cold!
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is when ice freezes
	FreezingC Celsius = 0
	// BoilingC is when water boils
	BoilingC Celsius = 100
	// AbsoluteZeroK is when all movement stops
	AbsoluteZeroK Kelvin = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
