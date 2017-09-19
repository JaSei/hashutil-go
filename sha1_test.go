package hashutil

import (
	"crypto/sha1"
	"testing"

	"github.com/stretchr/testify/assert"
)

const emptySha1 = "da39a3ee5e6b4b0d3255bfef95601890afd80709"

func TestSha1(t *testing.T) {
	s1, err := StringToSha1(emptySha1)
	assert.NoError(t, err, "Conversion string to Sha1 without errors")

	s2, err := StringToSha1("DA39A3EE5E6B4B0D3255BFEF95601890AFD80709")
	assert.NoError(t, err, "Conversion string to Sha1 without errors")

	assert.Equal(t, s1, s2, "Upper and lower of same Sha1 are equal")

	assert.Equal(t, s1, (Sha1)(sha1.New().Sum(nil)), "Empty hash calculated by crypto library are equal to hash from string")

	assert.Equal(t, s1.String(), emptySha1, "Converse Sha1 to strings")

	_, err = StringToSha1("")
	assert.Error(t, err, "Empty string isn't valid Sha1")

	_, err = StringToSha1("XA39A3EE5E6B4B0D3255BFEF95601890AFD80709")
	assert.Error(t, err, "X isn't valid char in Sha1")

	assert.Equal(t, EmptySha1().String(), emptySha1, "EmptySha1 function return Sha1 of nothing")

	_, err = BytesToSha1([]byte{})
	assert.Error(t, err)

	_, err = BytesToSha1([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	assert.NoError(t, err)
}
