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

func (i Ints) Discard(func(int) bool) Ints {
	return Ints{}
}

type Strings []string

func (s Strings) Keep(func(string) bool) Strings {
	return Strings{}
}

func (s Strings) Discard(func(string) bool) Strings {
	return Strings{}
}

type Lists []Ints

func (l Lists) Keep(func([]int) bool) Lists {
	return Lists{}
}

func (l Lists) Discard(func([]int) bool) Lists {
	return Lists{}
}
