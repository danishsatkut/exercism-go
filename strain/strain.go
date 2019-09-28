package strain

type Ints []int

func (i Ints) Keep(predicate func(int) bool) Ints {
	if i == nil {
		return nil
	}

	results := make(Ints, 0, len(i))

	for _, n := range i {
		if predicate(n) {
			results = append(results, n)
		}
	}

	return results
}

func (i Ints) Discard(predicate func(int) bool) Ints {
	if i == nil {
		return nil
	}

	results := make(Ints, 0, len(i))

	for _, n := range i {
		if !predicate(n) {
			results = append(results, n)
		}
	}

	return results
}

type Strings []string

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

func (s Strings) Discard(func(string) bool) Strings {
	return Strings{}
}

type Lists []Ints

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

func (l Lists) Discard(func([]int) bool) Lists {
	return Lists{}
}
