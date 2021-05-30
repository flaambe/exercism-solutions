// Package triangle provides a method for determining the type of a triangle
package triangle

import "math"

// Kind represents a type of triangle.
type Kind int

const (
	// NaT represents Not a triangle.
	NaT Kind = iota
	// Equ represents Equilateral triangle.
	Equ
	// Iso represents Isosceles triangle.
	Iso
	// Sca represents Scalene triangle.
	Sca
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

func isTriangle(a, b, c float64) bool {
	allPositive := a > 0 && b > 0 && c > 0
	isValid := a+b >= c && a+c >= b && b+c >= a
	isNumber := !math.IsInf(a*b*c, 0)

	return allPositive && isValid && isNumber
}
