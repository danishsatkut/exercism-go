package pythagorean

type Triplet struct {
	a, b, c int
}

func (t Triplet) isValid() bool {
	if t.a >= t.c || t.b >= t.c {
		return false
	}

	return true
}

func (t Triplet) isWithinRange(min, max int) bool {
	return t.a >= min && t.b >= min && t.c <= max
}

func (t Triplet) multiply(factor int) Triplet {
	return Triplet{
		a: t.a * factor,
		b: t.b * factor,
		c: t.c * factor,
	}
}

var Default = Triplet{3, 4, 5}

func Range(min, max int) []Triplet {
	var triplets = make([]Triplet, 0, max)

	for i := min / Default.a; i <= max/Default.c; i++ {
		if t := Default.multiply(i); t.isWithinRange(min, max) && t.isValid() {
			triplets = append(triplets, t)
		}
	}

	return triplets
}

func Sum(sum int) []Triplet {
	return nil
}
