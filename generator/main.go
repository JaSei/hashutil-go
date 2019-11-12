package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"log"
	"strings"
	"text/template"

	"github.com/JaSei/pathutil-go"
)

type cryptoTest struct {
	//all imports needed for test
	Imports []string
	//New contains new constructor for tested {{.Hash}}
	New string
	//NewOther contains new constructor for test of HashTo{{.Hash}}
	NewOther string
}

type hashDef struct {
	Hash            string
	Size            int
	EmptyHash       string
	EmptyBase64Hash string
	CryptoTest      cryptoTest
}

func main() {
	//for calculate base64 representation use https://cryptii.com/pipes/md5-hash (`view(empty)` -> `hash function` -> `base64` -> `text`)
	hashes := []hashDef{
		hashDef{
			"Md5",
			md5.Size,
			"d41d8cd98f00b204e9800998ecf8427e",
			"1B2M2Y8AsgTpgAmY7PhCfg==",
			cryptoTest{[]string{"crypto/md5", "crypto/sha1"}, "md5.New", "sha1.New"},
		},
		hashDef{
			"Sha1",
			sha1.Size,
			"da39a3ee5e6b4b0d3255bfef95601890afd80709",
			"2jmj7l5rSw0yVb/vlWAYkK/YBwk=",
			cryptoTest{[]string{"crypto/sha1", "crypto/sha256"}, "sha1.New", "sha256.New"},
		},
		hashDef{
			"Sha256",
			sha256.Size,
			"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
			cryptoTest{[]string{"crypto/sha256", "crypto/sha1"}, "sha256.New", "sha1.New"},
		},
		hashDef{
			"Sha384",
			sha512.Size384,
			"38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b",
			"OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLlb",
			cryptoTest{[]string{"crypto/sha512", "crypto/sha1"}, "sha512.New384", "sha1.New"},
		},
		{
			"Sha512",
			sha512.Size,
			"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
			"z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==",
			cryptoTest{[]string{"crypto/sha512", "crypto/sha1"}, "sha512.New", "sha1.New"},
		},
	}

	tmpl, err := template.ParseFiles("generator/code.tmpl", "generator/test.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	for _, hd := range hashes {
		generate(tmpl, "CODE", hd)
		generate(tmpl, "TEST", hd)
	}
}

func generate(tmpl *template.Template, tmplName string, hd hashDef) {
	postfix := "_generated.go"
	if tmplName == "TEST" {
		postfix = "_generated_test.go"
	}

	path, _ := pathutil.New(strings.ToLower(hd.Hash) + postfix)
	w, err := path.OpenWriter()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Generate %s", path)

	err = tmpl.ExecuteTemplate(w, tmplName, hd)
	if err != nil {
		log.Fatal(err)
	}

	w.Close()
}
