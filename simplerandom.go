package simplerandom

import (
	"crypto/rand"
	mrand "math/rand"
)

var (
	Number  = New("0123456789")
	Literal = New("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	URL     = New("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-")
)

type SimpleRandom struct {
	chars string
}

func New(chars string) *SimpleRandom {
	return &SimpleRandom{
		chars: chars,
	}
}

// GenerateBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func (r *SimpleRandom) GenerateBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	// Note that err == nil only if we read len(b) bytes.
	_, err := rand.Read(b)
	if err != nil {
		for i := 0; i < n; i++ {
			b[i] = byte(mrand.Intn(255))
		}
	}

	return b, nil
}

// GenerateString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func (r *SimpleRandom) GenerateString(n int) (string, error) {
	bytes, err := r.GenerateBytes(n)
	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = r.chars[b%byte(len(r.chars))]
	}

	return string(bytes), nil
}
