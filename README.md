# hashutil-go

[![Release](https://img.shields.io/github/release/avast/hashutil-go.svg?style=flat-square)](https://github.com/avast/hashutil-go/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/avast/hashutil-go.svg?style=flat-square)](https://travis-ci.org/avast/hashutil-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/avast/hashutil-go?style=flat-square)](https://goreportcard.com/report/github.com/avast/hashutil-go)
[![GoDoc](https://godoc.org/github.com/avast/hashutil-go?status.svg&style=flat-square)](http://godoc.org/github.com/avast/hashutil-go)
[![codecov.io](https://codecov.io/github/avast/hashutil-go/coverage.svg?branch=master)](https://codecov.io/github/avast/hashutil-go?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/avast/hashutil-go/-/badge.svg)](https://sourcegraph.com/github.com/avast/hashutil-go?badge)

hashutil is collection of common hash (crypto) types
and implement some simple function and method for hash manipulating

hash types are comparable ([]bytes)

## SYNOPSIS

    shaA, _ := hashtype.StringToSha256("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
    shaB, _ := hashtype.BytesToSha256(sha256.New().Sum(nil))

    if shaA == shaB {
        log.Println(shaA.String, "are equal")
    }
