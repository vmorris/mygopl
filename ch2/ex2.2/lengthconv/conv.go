package lengthconv

// FToM converts Feet to Meters
func FToM(f Feet) Meters { return Meters(f * .3048) }

// MToF converts Meters to Feet
func MToF(m Meters) Feet { return Feet(m * 3.28084) }
