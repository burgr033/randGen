package randomanimals

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateAdjective(t *testing.T) {
	adj := GenerateAdjective()
	require.Contains(t, adjectives, adj)
}

func TestGenerateName(t *testing.T) {
	name := GenerateName()
	require.Contains(t, animals, name)
}

func TestGeneratePair(t *testing.T) {
	tests := map[string]struct {
		sep      string
		expected string
	}{
		"space": {
			sep:      " ",
			expected: `\w+ \w+`,
		},
		"minus": {
			sep:      "-",
			expected: `\w+-\w+`,
		},
		"word": {
			sep:      "word",
			expected: `\w+word\w+`,
		},
		"none": {
			sep:      "",
			expected: `\w+`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := GeneratePair(tt.sep)
			require.Regexp(t, regexp.MustCompile(tt.expected), result)
		})
	}
}
