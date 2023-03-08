package main

import (
	"testing"
	"unicode/utf8"
)

// Unit test
func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"", ""},
		{"!12345", "54321!"},
	}
	for _, tc := range testCases {
		rev, err := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Error: %q, Reverse: %q, want %q", err, rev, tc.want)
		}
	}
}

// Fuzz test : The unit test has limitations, namely that each input must be added to the test by the developer.
// One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test cases you came up with did not reach.
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
