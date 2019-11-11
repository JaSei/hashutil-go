package hashutil

import (
	"crypto/md5"
	"crypto/sha1"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testEmptyMd5String = "d41d8cd98f00b204e9800998ecf8427e"

func TestStringToHashAndEqual(t *testing.T) {
	_, err := StringToMd5("")
	assert.Error(t, err, "Empty string isn't valid hash")
	assert.Equal(t, err.Error(), "Hash function Md5 must have a length of 16 bytes (actual have 0)")

	_, err = StringToMd5("x41d8cd98f00b204e9800998ecf8427e")
	assert.Error(t, err, "x isn't isn't valid char in hash")

	hash, err := StringToMd5(testEmptyMd5String)
	assert.NoError(t, err, "Convert string to Md5 without errors")

	hash2, err := StringToMd5("D41D8CD98F00B204E9800998ECF8427E")
	assert.NoError(t, err, "Convert string to Md5 without errors")

	assert.Equal(t, hash, hash2, "Upper and lower of same Md5 are equal")

	assert.True(t, hash.Equal(hash2) && hash2.Equal(hash))
}

func TestBytesToHash(t *testing.T) {
	_, err := BytesToMd5([]byte{})
	assert.Error(t, err)

	_, err = BytesToMd5([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	assert.NoError(t, err)
}

func TestEmptyHash(t *testing.T) {
	hash, err := StringToMd5(testEmptyMd5String)
	assert.NoError(t, err)
	assert.Equal(t, hash, EmptyMd5())
}

func TestHashToString(t *testing.T) {
	hash, err := StringToMd5(testEmptyMd5String)
	assert.NoError(t, err)
	assert.Equal(t, hash.String(), testEmptyMd5String)
	assert.Equal(t, hash.UpperString(), strings.ToUpper(testEmptyMd5String))
}

func TestIsEmpty(t *testing.T) {
	hash, err := StringToMd5(testEmptyMd5String)
	assert.NoError(t, err)
	assert.True(t, hash.IsEmpty())
}

func TestHashTo(t *testing.T) {
	hash, err := HashToMd5(md5.New())
	assert.NoError(t, err)
	assert.True(t, hash.IsEmpty())

	hash, err = HashToMd5(sha1.New())
	assert.Error(t, err)
}
