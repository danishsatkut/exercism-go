package isbn

import "testing"

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
