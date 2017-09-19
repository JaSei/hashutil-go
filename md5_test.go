package hashutil

import (
	"crypto/md5"
	"testing"

	"github.com/stretchr/testify/assert"
)

const emptyMd5 = "d41d8cd98f00b204e9800998ecf8427e"

func TestMd5(t *testing.T) {
	s1, err := StringToMd5(emptyMd5)
	assert.NoError(t, err, "Conversion string to Md5 without errors")

	s2, err := StringToMd5("D41D8CD98F00B204E9800998ECF8427E")
	assert.NoError(t, err, "Conversion string to Md5 without errors")

	assert.Equal(t, s1, s2, "Upper and lower of same Md5 are equal")

	assert.Equal(t, s1, (Md5)(md5.New().Sum(nil)), "Empty hash calculated by crypto library are equal to hash from string")

	assert.Equal(t, s1.String(), emptyMd5, "Converse Md5 to strings")

	_, err = StringToMd5("")
	assert.Error(t, err, "Empty string isn't valid Md5")

	_, err = StringToMd5("X41D8CD98F00B204E9800998ECF8427E")
	assert.Error(t, err, "X isn't valid char in Md5")

	assert.Equal(t, EmptyMd5().String(), emptyMd5, "EmptyMd5 function return Md5 of nothing")
}
