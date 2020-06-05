package isbn

import (
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	for _, test := range testCases {
		observed := IsValidISBN(test.isbn)
		if observed == test.expected {
			t.Logf("PASS: %s", test.description)
		} else {
			t.Errorf("FAIL: %s\nIsValidISBN(%q)\nExpected: %t, Actual: %t",
				test.description, test.isbn, test.expected, observed)
		}
	}
}

func TestConvertToISBN13(t *testing.T) {
	for _, test := range testCases {
		observed, err := ConvertToISBN13(test.isbn)

		if test.expected {
			if observed == test.isbn13 {
				t.Logf("PASS: %s", test.description)
			} else {
				t.Errorf("FAIL: %s\nConvertToISBN13(%q)\nExpected: %q, Actual: %q",
					test.description, test.isbn, test.isbn13, observed)
			}
		} else {
			if err == nil {
				t.Errorf("FAIL: Expected error for %q", test.isbn)
			} else {
				t.Logf("PASS: %s", test.description)
			}
		}
	}
}

func BenchmarkIsValidISBN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range testCases {
			IsValidISBN(n.isbn)
		}
	}
}
