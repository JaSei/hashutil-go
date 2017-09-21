package hashutil

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

const testEmptyMd5String = "d41d8cd98f00b204e9800998ecf8427e"

func TestStringToHashAndEqual(t *testing.T) {
	_, err := StringToHash(md5.New(), "")
	assert.Error(t, err, "Empty string isn't valid hash")
	assert.Equal(t, err.Error(), "Hash function represent by '*md5.digest' must have a length of 16 bytes (actual have 0)")

	_, err = StringToHash(md5.New(), "x41d8cd98f00b204e9800998ecf8427e")
	assert.Error(t, err, "x isn't isn't valid char in hash")

	hash, err := StringToHash(md5.New(), testEmptyMd5String)
	assert.NoError(t, err, "Convert string to Md5 without errors")

	hash2, err := StringToHash(md5.New(), "D41D8CD98F00B204E9800998ECF8427E")
	assert.NoError(t, err, "Convert string to Md5 without errors")

	assert.Equal(t, hash, hash2, "Upper and lower of same Md5 are equal")

	assert.True(t, hash.Equal(hash2) && hash2.Equal(hash))
}

func TestBytesToHash(t *testing.T) {
	_, err := BytesToHash(md5.New(), []byte{})
	assert.Error(t, err)

	_, err = BytesToHash(md5.New(), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	assert.NoError(t, err)
}

func TestEmptyHash(t *testing.T) {
	hash, err := StringToHash(md5.New(), testEmptyMd5String)
	assert.NoError(t, err)
	assert.Equal(t, hash, EmptyHash(md5.New()))
}

func TestHashToString(t *testing.T) {
	hash, err := StringToHash(md5.New(), testEmptyMd5String)
	assert.NoError(t, err)
	assert.Equal(t, hash.String(), testEmptyMd5String)
}

func TestIsEmpty(t *testing.T) {
	hash, err := StringToHash(md5.New(), testEmptyMd5String)
	assert.NoError(t, err)
	assert.True(t, hash.IsEmpty())
}

func TestMixOtherKindOfHash(t *testing.T) {
	hash1, err := StringToHash(md5.New(), testEmptyMd5String)
	assert.NoError(t, err)

	hash2 := EmptyHash(sha256.New())

	assert.False(t, hash1.Equal(hash2))
}

func TestMixOtherKindOfHashButWithSameLength(t *testing.T) {
	goVersion := runtime.Version()
	t.Log(goVersion)
	// e.g. go1.2.2 linux/amd64
	if goVersion[2] == 1 && goVersion[4] < 5 {
		t.Skip("Go 1.4 and less not implement New512_256 function")
	}

	hash1 := EmptyHash(sha512.New512_256())
	hash2 := EmptyHash(sha256.New())

	assert.NotEqual(t, hash1.String(), hash2.String(), "Sha512_256 is other hash then Sha256")
	assert.False(t, hash1.Equal(hash2))

	hashFake, err := BytesToHash(sha256.New(), hash1.ToBytes())

	assert.NoError(t, err)
	assert.False(t, hashFake.Equal(hash1), "Hash with same content, but other hash function are not equal")
}
