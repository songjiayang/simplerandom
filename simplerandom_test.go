package simplerandom

import (
	"testing"

	"github.com/golib/assert"
)

func Test_Random_Number(t *testing.T) {
	l, _ := Number.GenerateString(6)
	r, _ := Number.GenerateString(6)

	assert.NotEqual(t, l, r)
}

func Test_Random_Literal(t *testing.T) {
	l, _ := Literal.GenerateString(6)
	r, _ := Literal.GenerateString(6)

	assert.NotEqual(t, l, r)
}

func Test_Random_URL(t *testing.T) {
	l, _ := URL.GenerateString(6)
	r, _ := URL.GenerateString(6)

	assert.NotEqual(t, l, r)
}
