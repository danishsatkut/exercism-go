package strain

// Ints is a collection of int
type Ints []int

// Keep returns a new Ints containing those int where the predicate is true
func (i Ints) Keep(predicate func(int) bool) Ints {
	return i.eachWith(func(results Ints, n int) Ints {
		if predicate(n) {
			return append(results, n)
		}

		return results
	})
}

// Discard returns a new Ints containing those int where the predicate is false
func (i Ints) Discard(predicate func(int) bool) Ints {
	return i.eachWith(func(results Ints, n int) Ints {
		if !predicate(n) {
			return append(results, n)
		}

		return results
	})
}

func (i Ints) eachWith(operation func(Ints, int) Ints) Ints {
	if i == nil {
		return nil
	}

	results := make(Ints, 0, len(i))

	for _, n := range i {
		results = operation(results, n)
	}

	return results
}

// Strings is a collection of string
type Strings []string

// Keep returns a new Strings containing those string where the predicate is true
func (s Strings) Keep(predicate func(string) bool) Strings {
	if s == nil {
		return nil
	}

	results := make(Strings, 0, len(s))

	for _, el := range s {
		if predicate(el) {
			results = append(results, el)
		}
	}

	return results
}

// Lists is a collection of Ints
type Lists []Ints

// Keep returns a new List containing those Ints where the predicate is true
func (l Lists) Keep(predicate func([]int) bool) Lists {
	if l == nil {
		return nil
	}

	results := make(Lists, 0, len(l))

	for _, list := range l {
		if predicate(list) {
			results = append(results, list)
		}
	}

	return results
}
