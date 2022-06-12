package simplerandom

import (
	"testing"
)

func Test_Random_Number(t *testing.T) {
	l, _ := Number.GenerateString(6)
	r, _ := Number.GenerateString(6)
	if l == r {
		t.Fatalf("two random numbers should be different")
	}
}

func Test_Random_Literal(t *testing.T) {
	l, _ := Literal.GenerateString(6)
	r, _ := Literal.GenerateString(6)
	if l == r {
		t.Fatalf("two random literals should be different")
	}
}

func Test_Random_URL(t *testing.T) {
	l, _ := URL.GenerateString(6)
	r, _ := URL.GenerateString(6)
	if l == r {
		t.Fatalf("two random urls should be different")
	}
}
