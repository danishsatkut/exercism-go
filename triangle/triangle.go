// Package triangle provides functionality related to triangle geometric shape.
package triangle

import "math"

// Kind represents kind of triangle
type Kind int

// Constants to represent types of triangles
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the type of triangle based on the
// length of the sides provided.
func KindFromSides(a, b, c float64) Kind {
	return newTriangle(a, b, c).kind()
}

type triangle struct {
	a, b, c side
}

func newTriangle(a, b, c float64) triangle {
	return triangle{side(a), side(b), side(c)}
}

func (t triangle) isInEqual() bool {
	return t.a.isInEqual(t.b, t.c) && t.b.isInEqual(t.a, t.c) && t.c.isInEqual(t.a, t.b)
}

func (t triangle) isValid() bool {
	return t.a.isValid() && t.b.isValid() && t.c.isValid() && t.isInEqual()
}

func (t triangle) kind() Kind {
	if !t.isValid() {
		return NaT
	}

	if t.isEquilateral() {
		return Equ
	}

	if t.isIsosceles() {
		return Iso
	}

	return Sca
}

func (t triangle) isEquilateral() bool {
	return t.a.isEqual(t.b) && t.b.isEqual(t.c)
}

func (t triangle) isIsosceles() bool {
	return t.a.isEqual(t.b) || t.b.isEqual(t.c) || t.c.isEqual(t.a)
}

type side float64

func (s side) isValid() bool {
	return !(s <= 0 || math.IsNaN(float64(s)) || math.IsInf(float64(s), 0))
}

func (s side) isEqual(other side) bool {
	return s == other
}

func (s side) isInEqual(x, y side) bool {
	return s <= x+y
}
