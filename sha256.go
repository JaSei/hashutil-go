package hashutil

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Sha256 []byte

var badSha256Length = fmt.Errorf("Sha256 must have a length of %d bytes", sha256.Size)

// convert type Sha256 to string representation (lower case)
func (hash Sha256) String() string {
	return hex.EncodeToString(hash)
}

// StringToSha256 validate and convert a string to Sha256 type
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

// BytesToSha256 validate and convert []bytes to Sha256 type
func BytesToSha256(bytes []byte) (Sha256, error) {
	if len(bytes) != sha256.Size {
		return nil, badSha256Length
	}

	return bytes, nil
}

// EmptySha256 return type Sha256 of empty (zero-length) string
func EmptySha256() Sha256 {
	return sha256.New().Sum(nil)
}
