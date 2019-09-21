// Package accumulate provides functionality to accumulate results of operations
package accumulate

// Accumulate performs an operation on items of list and returns
// a list of results.
func Accumulate(list []string, op func(string) string) []string {
	result := make([]string, 0, len(list))

	for _, s := range list {
		result = append(result, op(s))
	}

	return result
}
