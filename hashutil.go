/*
Package hashutil is collection of common hash (crypto) types.
It also contains some helper functions for hash manipulation.



SYNOPSIS

	import (
		"github.com/avast/hashutil-go",
		"crypto/md5"
	)

	hash, _ := hashutil.StringToMd5("d41d8cd98f00b204e9800998ecf8427e")
	otherHash, _ := hashutil.BytesToMd5(md5.New().Sum(nil))
	// or
	// otherhash := hashutil.EmptyMd5()

	if hash.Equal(otherHash) {
		// do something
	}
*/
package hashutil

//go:generate go run generator/main.go
