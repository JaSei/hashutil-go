package hashutil

import (
	"crypto/md5"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testEmptyMd5String = "d41d8cd98f00b204e9800998ecf8427e"

func TestMd5(t *testing.T) {
	s1, err := StringToMd5(testEmptyMd5String)
	assert.NoError(t, err, "Convert string to Md5 without errors")

	s2, err := StringToMd5("D41D8CD98F00B204E9800998ECF8427E")
	assert.NoError(t, err, "Convert string to Md5 without errors")

	assert.Equal(t, s1, s2, "Upper and lower of same Md5 are equal")

	assert.Equal(t, s1, (Md5)(md5.New().Sum(nil)), "An empty hash calculated by crypto library is equal to hash from string")

	assert.Equal(t, s1.String(), testEmptyMd5String, "Convert Md5 to strings")

	_, err = StringToMd5("")
	assert.Error(t, err, "Empty string isn't valid Md5")

	_, err = StringToMd5("X41D8CD98F00B204E9800998ECF8427E")
	assert.Error(t, err, "X isn't valid char in Md5")

	assert.Equal(t, EmptyMd5().String(), testEmptyMd5String, "EmptyMd5 function return Md5 of nothing")

	_, err = BytesToMd5([]byte{})
	assert.Error(t, err)

	_, err = BytesToMd5([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	assert.NoError(t, err)

}
