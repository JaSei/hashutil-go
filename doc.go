/*
Hashutil is collection of common hash (crypto) types.
It also contains some helper functions for hash manipulation.

SYNOPSIS

	shaA, _ := hashtype.StringToSha256("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	shaB, _ := hashtype.BytesToSha256(sha256.New().Sum(nil))

	if shaA == shaB {
		log.Println(shaA.String, "are equal")
	}

*/

package hashutil
