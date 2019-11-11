package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"log"
	"strings"
	"text/template"

	"github.com/JaSei/pathutil-go"
)

type hashDef struct {
	Hash      string
	Size      int
	EmptyHash string
}

const tmpl = `// Code generated by github.com/jasei/hashutil/generator DO NOT EDIT
package hashutil

import (
    "bytes"
	"encoding/hex"
	"encoding/base64"
    "fmt"
	"hash"
    "strings"
)

// {{.Hash}} type represents {{.Hash}} checksum
type {{.Hash}} []byte

// StringTo{{.Hash}} return a new {{.Hash}} checksum from string (hex) representation
func StringTo{{.Hash}}(hexString string) ({{.Hash}}, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return {{.Hash}}{}, err
	}

    return BytesTo{{.Hash}}(bytes)
}

// BytesTo{{.Hash}} return a new {{.Hash}} checksum from bytes (binary) representation
func BytesTo{{.Hash}}(bytes []byte) ({{.Hash}}, error) {
	if len(bytes) != {{.Size}} {
        return {{.Hash}}{}, fmt.Errorf("Hash function {{.Hash}} must have a length of {{.Size}} bytes (actual have %d)", len(bytes))
	}

	return {{.Hash}}(bytes), nil

}

//HashTo{{.Hash}} return a new {{.Hash}} checksum from hash.Hash representation
// HashTo{{.Hash}} convert hashutil.Hash to {{.Hash}}
func HashTo{{.Hash}}(h hash.Hash) ({{.Hash}}, error) {
	return BytesTo{{.Hash}}(h.Sum(nil))
}

// Empty{{.Hash}} return {{.Hash}} of empty file
func Empty{{.Hash}}() {{.Hash}} {
	h,_ := StringTo{{.Hash}}("{{.EmptyHash}}")
	return h
}

// Equal return true if is {{.Hash}}s equal
func (h {{.Hash}}) Equal(o {{.Hash}}) bool {
	return bytes.Equal(h, o)
}

// String return (hex) string representation of {{.Hash}}
func (h {{.Hash}}) String() string {
	return hex.EncodeToString(h)
}

// UpperString return (hex) string representation in upper case of {{.Hash}}
func (h {{.Hash}}) UpperString() string {
	return strings.ToUpper(hex.EncodeToString(h))
}

// ToBytes return []byte of hashutil.{{.Hash}}
func (h {{.Hash}}) ToBytes() []byte {
	return h
}

// ToBase64 return base64 representation of {{.Hash}}
func (h {{.Hash}}) ToBase64() string {
	return base64.StdEncoding.EncodeToString(([]byte)(h))
}

// IsEmpty return true if is {{.Hash}} 'empty hash'
func (h {{.Hash}}) IsEmpty() bool {
	return Empty{{.Hash}}().Equal(h)
}

`

func main() {
	hashes := []hashDef{
		newHashDef("Md5", md5.Size, md5.New()),
		newHashDef("Sha1", sha1.Size, sha1.New()),
		newHashDef("Sha256", sha256.Size, sha256.New()),
		newHashDef("Sha384", sha512.Size384, sha512.New384()),
		newHashDef("Sha512", sha512.Size, sha512.New()),
	}

	tmpl, err := template.New("hash").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	for _, hd := range hashes {
		path, _ := pathutil.New(strings.ToLower(hd.Hash) + "_generated.go")
		w, err := path.OpenWriter()

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Generate %s", path)

		err = tmpl.Execute(w, hd)
		if err != nil {
			log.Fatal(err)
		}

		w.Close()
	}
}

func newHashDef(name string, count int, h hash.Hash) hashDef {
	return hashDef{name, count, hex.EncodeToString(h.Sum(nil))}
}