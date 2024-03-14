package main

import (
	"testing"
)

var tests = []struct {
	input  string
	output string
	err    bool
}{
	{"a4bc2d5e", "aaaabccddddde", false},
	{"abcd", "abcd", false},
	{"45", "", true},
	{"", "", false},
}

func TestUnpack(t *testing.T) {
	for _, tt := range tests {
		res, err := Unpack(tt.input)
		if (res != tt.output) || ((err != nil) != tt.err) {
			t.Errorf("Unpack(%s) = %s, expected %s, ERROR: %t , expected %t",
				tt.input, res, tt.output, err != nil, tt.err)
		}
	}
}
