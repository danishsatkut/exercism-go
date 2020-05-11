package pythagorean

import (
	"math"
	"sort"
)

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var uniqTriplets = make(map[Triplet]struct{})

	for _, triplet := range primitiveTriplets(1, max) {
		for i := min / triplet.a; i <= max/triplet.c; i++ {
			if t := triplet.multiply(i); t.isWithinRange(min, max) && t.isValid() {
				uniqTriplets[t] = struct{}{}
			}
		}
	}

	var triplets = make([]Triplet, 0, len(uniqTriplets))

	for t := range uniqTriplets {
		triplets = append(triplets, t)
	}

	return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(sum int) []Triplet {
	triplets := make([]Triplet, 0, 5)

	for _, triplet := range Range(1, sum/2) {
		if triplet.perimeter() == sum {
			triplets = append(triplets, triplet)
		}
	}

	sort.Slice(triplets, func(i, j int) bool {
		return triplets[i].lessThan(triplets[j])
	})

	return triplets
}

// Triplet represents a Pythagorean triplet
type Triplet struct {
	a, b, c int
}

func (t Triplet) isValid() bool {
	if t.a >= t.c || t.b >= t.c || t.a >= t.b {
		return false
	}

	return true
}

func (t Triplet) isWithinRange(min, max int) bool {
	return t.a >= min && t.b >= min && t.c <= max
}

func (t Triplet) lessThan(other Triplet) bool {
	if t.a != other.a {
		return t.a < other.a
	}

	if t.b != other.b {
		return t.b < other.b
	}

	return t.c < other.c
}

func (t Triplet) multiply(factor int) Triplet {
	return Triplet{
		a: t.a * factor,
		b: t.b * factor,
		c: t.c * factor,
	}
}

func (t Triplet) perimeter() int {
	return t.a + t.b + t.c
}

func primitiveTriplets(min, max int) []Triplet {
	var primitiveTriplets = make([]Triplet, 0, max)

	for i := min; i <= max; i++ {
		if t := newTriplet(i); t.isValid() {
			primitiveTriplets = append(primitiveTriplets, t)
		}
	}

	return primitiveTriplets
}

func newTriplet(n int) Triplet {
	t := Triplet{a: n}

	if n%2 == 0 {
		sq := math.Pow(float64(n)/2, 2)

		t.b = int(sq - 1)
		t.c = int(sq + 1)
	} else {
		sq := math.Pow(float64(n), 2) / 2

		t.b = int(sq - 0.5)
		t.c = int(sq + 0.5)
	}

	return t
}
