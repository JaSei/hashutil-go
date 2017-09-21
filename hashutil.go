package hashutil

import (
	"bytes"
	"crypto"
	"encoding/hex"
	"fmt"
)

type Hash []byte

func StringToHash(hashType crypto.Hash, str string) (Hash, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	if err := checkLen(hashType, bytes); err != nil {
		return nil, err
	}

	return bytes, nil

}

func checkLen(hashType crypto.Hash, bytes []byte) error {
	actualLength := len(bytes)
	if actualLength != hashType.Size() {
		return fmt.Errorf("Hash function #%d must have a length of %d bytes (actual have %d)", hashType, hashType.Size(), actualLength)
	}

	return nil
}

func BytesToHash(hashType crypto.Hash, bytes []byte) (Hash, error) {
	if err := checkLen(hashType, bytes); err != nil {
		return nil, err
	}

	return bytes, nil
}

func (hash Hash) Equal(otherHash Hash) bool {
	return bytes.Compare(hash, otherHash) == 0
}

func (hash Hash) String() string {
	return hex.EncodeToString(hash)
}
