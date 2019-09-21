package accumulate

func Accumulate(list []string, op func(string) string) []string {
	result := make([]string, 0, len(list))

	for _, s := range list {
		result = append(result, op(s))
	}

	return result
}
