package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetOnesFromFirstOne64(t *testing.T) {
	var tests = []struct {
		in       uint64
		expected uint64
	}{
		{
			in:       0x8000000000000000,
			expected: 0xffffffffffffffff,
		},
		{
			in:       0x0000300000000000,
			expected: 0x00003fffffffffff,
		},
	}
	for _, test := range tests {
		out := SetAllToOnesFromFirstOne64(test.in)
		out1 := SetAllToOnesFromFirstOne64_1(test.in)
		assert.Equal(t, test.expected, out)
		assert.Equal(t, test.expected, out1)
	}
}

func TestSetAllToOnesFromFirstOne64(t *testing.T) {
	var tests = []struct {
		in       uint64
		expected int
	}{
		{
			in:       0x8000000000000000,
			expected: 64,
		},
	}
	for _, test := range tests {
		out := GetMinNumBits64(test.in)
		out1 := getMinNumBits64_1(test.in)
		assert.Equal(t, test.expected, out)
		assert.Equal(t, test.expected, out1)
	}
}
