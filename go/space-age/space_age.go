// Package space provides a method for calculating the age on the planet
package space

// Planet represents a planet as a string
type Planet string

// Age calculates how old someone would be on a given planet
func Age(s float64, p Planet) float64 {
	const earthPeriod = 31557600.0

	planetsPeriod := map[Planet]float64{
		"Mercury": 0.2408467 * earthPeriod,
		"Venus":   0.61519726 * earthPeriod,
		"Earth":   earthPeriod,
		"Mars":    1.8808158 * earthPeriod,
		"Jupiter": 11.862615 * earthPeriod,
		"Saturn":  29.447498 * earthPeriod,
		"Uranus":  84.016846 * earthPeriod,
		"Neptune": 164.79132 * earthPeriod,
	}

	return s / planetsPeriod[p]
}
