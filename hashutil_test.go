package hashutil

import (
	"crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

//const testEmptyMd5String = "d41d8cd98f00b204e9800998ecf8427e"

func TestStringToHashAndEqual(t *testing.T) {
	_, err := StringToHash(crypto.MD5, "")
	assert.Error(t, err, "Empty string isn't valid hash")

	_, err = StringToHash(crypto.MD5, "x41d8cd98f00b204e9800998ecf8427e")
	assert.Error(t, err, "x isn't isn't valid char in hash")

	hash, err := StringToHash(crypto.MD5, testEmptyMd5String)
	assert.NoError(t, err, "Convert string to Md5 without errors")

	hash2, err := StringToHash(crypto.MD5, "D41D8CD98F00B204E9800998ECF8427E")
	assert.NoError(t, err, "Convert string to Md5 without errors")

	assert.Equal(t, hash, hash2, "Upper and lower of same Md5 are equal")

	assert.True(t, hash.Equal(hash2) && hash2.Equal(hash))
}

func TestBytesToHash(t *testing.T) {
	_, err := BytesToHash(crypto.MD5, []byte{})
	assert.Error(t, err)

	_, err = BytesToHash(crypto.MD5, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	assert.NoError(t, err)
}
