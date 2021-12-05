package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateAppName(t *testing.T) {
	for _, tt := range []struct {
		name string
		ok   bool
	}{
		{"aladdin", true},
		{"al23addin", true},
		{"common-qae", true},
		{"23aladdin", false},
		{"common_qae", false},
		{"q", true},
		{"", false},
		{"-common-qae", false},
	} {
		err := ValidateAppName(tt.name)
		assert.Equal(t, tt.ok, err == nil)
	}
}
