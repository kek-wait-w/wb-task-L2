package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input  []string
	output []string
	fl     flags
}{
	{[]string{"bb cc dd", "aa bb cc"}, []string{"aa bb cc", "bb cc dd"}, flags{-1, false, false, false, ""}},
	{[]string{"aa bb cc", "bb cc dd"}, []string{"bb cc dd", "aa bb cc"}, flags{-1, false, true, false, ""}},
	{[]string{"aa bb cc", "bb cc dd", "bb cc dd"}, []string{"bb cc dd", "aa bb cc"}, flags{-1, false, true, true, ""}},
	{[]string{"4 5 6", "1 2 3"}, []string{"1 2 3", "4 5 6"}, flags{-1, true, false, false, ""}},
	{[]string{"4 5 7", "1 2 6", "3 2 5"}, []string{"3 2 5", "1 2 6", "4 5 7"}, flags{3, true, false, false, ""}},
}

func TestSort(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, Sort(tt.input, tt.fl), tt.output)
	}
}
