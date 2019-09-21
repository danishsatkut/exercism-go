// Package accumulate provides functionality to accumulate results of operations
package accumulate

// Mapper is function that maps a string to another string
type Mapper func(string) string

// Accumulate performs an operation on items of list and returns
// a list of results.
func Accumulate(list []string, mapper Mapper) []string {
	result := make([]string, 0, len(list))

	for _, s := range list {
		result = append(result, mapper(s))
	}

	return result
}
