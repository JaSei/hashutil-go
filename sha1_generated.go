// Code generated by github.com/jasei/hashutil/generator DO NOT EDIT
package hashutil

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

// Sha1 type represents Sha1 checksum
type Sha1 []byte

// StringToSha1 return a new Sha1 checksum from string (hex) representation
func StringToSha1(hexString string) (Sha1, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return Sha1{}, err
	}

	return BytesToSha1(bytes)
}

// BytesToSha1 return a new Sha1 checksum from bytes (binary) representation
func BytesToSha1(bytes []byte) (Sha1, error) {
	if len(bytes) != 20 {
		return Sha1{}, fmt.Errorf("Hash function Sha1 must have a length of 20 bytes (actual have %d)", len(bytes))
	}

	return Sha1(bytes), nil

}

//HashToSha1 return a new Sha1 checksum from hash.Hash representation
// HashToSha1 convert hashutil.Hash to Sha1
func HashToSha1(h hash.Hash) (Sha1, error) {
	return BytesToSha1(h.Sum(nil))
}

// EmptySha1 return Sha1 of empty file
func EmptySha1() Sha1 {
	h, _ := StringToSha1("da39a3ee5e6b4b0d3255bfef95601890afd80709")
	return h
}

// Equal return true if is Sha1s equal
func (h Sha1) Equal(o Sha1) bool {
	return bytes.Equal(h, o)
}

// String return (hex) string representation of Sha1
func (h Sha1) String() string {
	return hex.EncodeToString(h)
}

// UpperString return (hex) string representation in upper case of Sha1
func (h Sha1) UpperString() string {
	return strings.ToUpper(hex.EncodeToString(h))
}

// ToBytes return []byte of hashutil.Sha1
func (h Sha1) ToBytes() []byte {
	return h
}

// ToBase64 return base64 representation of Sha1
func (h Sha1) ToBase64() string {
	return base64.StdEncoding.EncodeToString(([]byte)(h))
}

// IsEmpty return true if is Sha1 'empty hash'
func (h Sha1) IsEmpty() bool {
	return EmptySha1().Equal(h)
}
