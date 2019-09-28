package strain

type Ints []int

func (i Ints) Keep(func(int) bool) Ints {
	return Ints{}
}

func (i Ints) Discard(func(int) bool) Ints {
	return Ints{}
}
