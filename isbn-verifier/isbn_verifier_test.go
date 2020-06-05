package isbn

import (
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	t.Run("for ISBN-10", func(t *testing.T) {
		for _, test := range testCases {
			observed := IsValidISBN(test.isbn)
			if observed == test.expected {
				t.Logf("PASS: %s", test.description)
			} else {
				t.Errorf("FAIL: %s\nIsValidISBN(%q)\nExpected: %t, Actual: %t",
					test.description, test.isbn, test.expected, observed)
			}
		}
	})

	t.Run("for ISBN-13", func(t *testing.T) {
		for _, test := range testCases {
			observed := IsValidISBN(test.isbn13)
			if observed == test.expected {
				t.Logf("PASS: %s", test.description)
			} else {
				t.Errorf("FAIL: %s\nIsValidISBN(%q)\nExpected: %t, Actual: %t",
					test.description, test.isbn13, test.expected, observed)
			}
		}
	})
}

func BenchmarkIsValidISBN(b *testing.B) {
	b.Run("for ISBN-10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, n := range testCases {
				IsValidISBN(n.isbn)
			}
		}
	})

	b.Run("for ISBN-13", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, n := range testCases {
				IsValidISBN(n.isbn13)
			}
		}
	})
}
