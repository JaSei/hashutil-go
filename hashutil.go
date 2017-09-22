/*
Hashutil is collection of common hash (crypto) types.
It also contains some helper functions for hash manipulation.



SYNOPSIS

	import (
		"github.com/avast/hashutil-go",
		"crypto/md5"
	)

	hash, _ := hashutil.StringToHash(md5.New(), "d41d8cd98f00b204e9800998ecf8427e")
	otherHash, _ := hashutil.BytesToHash(md5.New(), md5.New().Sum(nil))
	// or
	// otherhash := hashutil.EmptyHash(md5.New())

	if hash.Equal(otherHash) {
		// do something
	}
*/
package hashutil

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"hash"
	"reflect"
)

// hashutil.Hash represents hash in native form and kind of hash
type Hash struct {
	hash []byte
	kind hash.Hash
}

// StringToHash validate and convert hexString of hash function implement hash.Hash to hashutil.Hash representation
func StringToHash(hash hash.Hash, hexString string) (Hash, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return Hash{}, err
	}

	if err := checkLen(hash, bytes); err != nil {
		return Hash{}, err
	}

	return Hash{bytes, hash}, nil

}

// BytesToHash validate and convert []byte of hash function implement hash.Hash interface to hashutil.Hash representation
func BytesToHash(hash hash.Hash, bytes []byte) (Hash, error) {
	if err := checkLen(hash, bytes); err != nil {
		return Hash{}, err
	}

	return Hash{bytes, hash}, nil
}

func checkLen(hash hash.Hash, bytes []byte) error {
	actualLength := len(bytes)
	if actualLength != hash.Size() {
		return fmt.Errorf(
			"Hash function represent by '%s' must have a length of %d bytes (actual have %d)",
			reflect.TypeOf(hash),
			hash.Size(),
			actualLength,
		)
	}

	return nil
}

// EmptyHash return hashutil.Hash of empty (zero-length) hash function implement hash.Hash
func EmptyHash(hash hash.Hash) Hash {
	return Hash{hash.Sum(nil), hash}
}

// Equal two hashutil.Hash, return true if hashes are same
func (hash Hash) Equal(otherHash Hash) bool {
	return reflect.TypeOf(hash.kind) == reflect.TypeOf(otherHash.kind) &&
		bytes.Equal(hash.hash, otherHash.hash)
}

// String return hexString representation in lower case of hashutil.Hash
func (hash Hash) String() string {
	return hex.EncodeToString(hash.hash)
}

// ToBytes return []byte of hashutil.Hash
func (hash Hash) ToBytes() []byte {
	return hash.hash
}

// IsEmpty return true, if hashutil.Hash represents empty (zero-length) hash (of given hash.Hash)
func (hash Hash) IsEmpty() bool {
	return hash.Equal(EmptyHash(hash.kind))
}
