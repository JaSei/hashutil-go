/*
hashutil is collection of common hash (crypto) types
and implement some simple function and method for hash manipulating

hash types are comparable ([]bytes)

SYNOPSIS

	shaA, _ := hashtype.StringToSha256("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	shaB, _ := hashtype.BytesToSha256(sha256.New().Sum(nil))

	if shaA == shaB {
		log.Println(shaA.String, "are equal")
	}

*/
package hashutil

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type Sha256 []byte

var badSha256Length = errors.New("Sha256 must have 32 bytes length")

// convert type Sha256 to string representation (upper case)
func (hash Sha256) String() string {
	return hex.EncodeToString(hash)
}

// StringToSha256 valid and convert string to Sha256 type
func StringToSha256(str string) (Sha256, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	if len(bytes) != sha256.Size {
		return nil, badSha256Length
	}

	return bytes, nil
}

// BytesToSha256 valid and convert []bytes to Sha256 type
func BytesToSha256(bytes []byte) (Sha256, error) {
	if len(bytes) != sha256.Size {
		return nil, badSha256Length
	}

	for _, b := range bytes {
		if b > 0xff {
			return nil, errors.New("Sha256 must contains only 0-FF bytes")
		}
	}

	return bytes, nil
}

// EmptySha256 return type Sha256 of empty string
func EmptySha256() Sha256 {
	return sha256.New().Sum(nil)
}
