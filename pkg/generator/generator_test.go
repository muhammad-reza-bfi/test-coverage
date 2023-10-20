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

func TestFanIn(t *testing.T) {
	t.Parallel()

	genOdd := GenerateOddOrEvenRaw(1)
	genEven := GenerateOddOrEvenRaw(2)
	gen := FanIn(genOdd, genEven)
	assert.NotEqual(t, <-gen, <-gen)
}
