package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input  string
	output string
	fl     flags
}{
	{"abc\tfcd\tfds\ndff\tdff\tfdfi", "abc" + "\n" + "dff", flags{fields: 1, delimiter: "\t"}},
	{"abc-fcd-fds\ndff-dff-fdfi", "abc" + "\n" + "dff", flags{fields: 1, delimiter: "-"}},
	{"abcfcdfds\ndff-dff-fdfi", "dff", flags{fields: 1, delimiter: "-", separated: true}},
}

func TestCut(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, Cut(tt.input, tt.fl), tt.output)
	}
}
