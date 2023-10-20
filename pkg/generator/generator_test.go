package generator_test

import (
	"testing"

	. "github.com/muhammad-reza-bfi/test-coverage/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func TestGenerateOddOrEven(t *testing.T) {
	t.Parallel()

	gen := GenerateOddOrEven(1)
	assert.Equal(t, gen(), int64(1))
	assert.Equal(t, gen(), int64(3))
}
