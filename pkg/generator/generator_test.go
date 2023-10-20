package generator_test

import (
	"testing"

	. "github.com/elangreza14/test-coverage/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func TestGenerateOddOrEven(t *testing.T) {
	gen := GenerateOddOrEven(1)
	assert.Equal(t, gen(), int64(1))
	assert.Equal(t, gen(), int64(3))
}
