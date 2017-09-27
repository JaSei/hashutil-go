# hashutil

[![Release](https://img.shields.io/github/release/avast/hashutil-go.svg?style=flat-square)](https://github.com/avast/hashutil-go/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/avast/hashutil-go.svg?style=flat-square)](https://travis-ci.org/avast/hashutil-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/avast/hashutil-go?style=flat-square)](https://goreportcard.com/report/github.com/avast/hashutil-go)
[![GoDoc](https://godoc.org/github.com/avast/hashutil-go?status.svg&style=flat-square)](http://godoc.org/github.com/avast/hashutil-go)
[![codecov.io](https://codecov.io/github/avast/hashutil-go/coverage.svg?branch=master)](https://codecov.io/github/avast/hashutil-go?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/avast/hashutil-go/-/badge.svg)](https://sourcegraph.com/github.com/avast/hashutil-go?badge)

Hashutil is collection of common hash (crypto) types. It also contains some
helper functions for hash manipulation.

### SYNOPSIS

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

## Usage

#### type Hash

```go
type Hash struct {
}
```

hashutil.Hash represents hash in native form and kind of hash

#### func  BytesToHash

```go
func BytesToHash(hash hash.Hash, bytes []byte) (Hash, error)
```
BytesToHash validate and convert []byte of hash function implement hash.Hash
interface to hashutil.Hash representation

#### func  EmptyHash

```go
func EmptyHash(hash hash.Hash) Hash
```
EmptyHash return hashutil.Hash of empty (zero-length) hash function implement
hash.Hash

#### func  StringToHash

```go
func StringToHash(hash hash.Hash, hexString string) (Hash, error)
```
StringToHash validate and convert hexString of hash function implement hash.Hash
to hashutil.Hash representation

#### func (Hash) Equal

```go
func (hash Hash) Equal(otherHash Hash) bool
```
Equal two hashutil.Hash, return true if hashes are same

#### func (Hash) IsEmpty

```go
func (hash Hash) IsEmpty() bool
```
IsEmpty return true, if hashutil.Hash represents empty (zero-length) hash (of
given hash.Hash)

#### func (Hash) String

```go
func (hash Hash) String() string
```
String return hexString representation in lower case of hashutil.Hash

#### func (Hash) ToBytes

```go
func (hash Hash) ToBytes() []byte
```
ToBytes return []byte of hashutil.Hash

## Contributing

Contributions are very much welcome.

### Makefile

Makefile provides several handy rules, like README.md `generator` , `setup` for prepare build/dev environment, `test`, `cover`, etc...

Try `make help` for more information.
