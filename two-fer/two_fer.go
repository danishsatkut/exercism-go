// Package twofer provides functionality to share with a provided name.
package twofer

import "fmt"

// ShareWith returns a two-fer statement for a given name.
// If no name is provided is replaces person's name with you.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
