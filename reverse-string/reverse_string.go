package reverse

// Reverse performs string reversal
func Reverse(s string) string {
	r := []rune(s)
	count := len(r)

	output := make([]rune, count)

	for i := 0; i < count; i++ {
		output[count-i-1] = r[i]
	}

	return string(output)
}
