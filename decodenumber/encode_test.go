package decodenumber

import (
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name     string
		encoded  string
		expected string
	}{
		{
			name:     "case 1",
			encoded:  "LLRR=",
			expected: "210122",
		},
		{
			name:     "case 2 ==RLL",
			encoded:  "==RLL",
			expected: "000210",
		},
		{
			name:     "case 3",
			encoded:  "=LLRR",
			expected: "221012",
		},
		{
			name:     "case 4",
			encoded:  "RRL=R",
			expected: "012001",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := Execute(tc.encoded)
			if actual != tc.expected {
				t.Errorf("Test %s failed. Expected %s but got %s", tc.name, tc.expected, actual)
			}
		})
	}
}
