package hashutil

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
)

type Md5 []byte

var badMd5Length = fmt.Errorf("Md5 must have %d bytes length", md5.Size)

// convert type Md5 to string representation (upper case)
func (hash Md5) String() string {
	return hex.EncodeToString(hash)
}

// StringToMd5 valid and convert string to Md5 type
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

// BytesToMd5 valid and convert []bytes to Md5 type
func BytesToMd5(bytes []byte) (Md5, error) {
	if len(bytes) != md5.Size {
		return nil, badMd5Length
	}

	for _, b := range bytes {
		if b > 0xff {
			return nil, errors.New("Md5 must contains only 0-FF bytes")
		}
	}

	return bytes, nil
}

// EmptyMd5 return type Md5 of empty string
func EmptyMd5() Md5 {
	return md5.New().Sum(nil)
}
