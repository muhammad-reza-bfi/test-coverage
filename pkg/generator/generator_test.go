package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOddOrEven(t *testing.T) {
	gen := GenerateOddOrEven(1)
	assert.Equal(t, gen(), int64(1))
	assert.Equal(t, gen(), int64(3))
}
