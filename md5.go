package hashutil

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Md5 []byte

var badMd5Length = fmt.Errorf("Md5 must have a length of %d bytes", md5.Size)

// convert type Md5 to string representation (lower case)
func (hash Md5) String() string {
	return hex.EncodeToString(hash)
}

// compare two Md5 types
func (hash Md5) Compare(otherHash Md5) bool {
	return bytes.Compare(hash, otherHash) == 0
}

// StringToMd5 validate and convert a string to Md5 type
func StringToMd5(str string) (Md5, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	if len(bytes) != md5.Size {
		return nil, badMd5Length
	}

	return bytes, nil
}

// BytesToMd5 validate and convert []bytes to Md5 type
func BytesToMd5(bytes []byte) (Md5, error) {
	if len(bytes) != md5.Size {
		return nil, badMd5Length
	}

	return bytes, nil
}

// EmptyMd5 return type Md5 of empty (zero-length) string
func EmptyMd5() Md5 {
	return md5.New().Sum(nil)
}
