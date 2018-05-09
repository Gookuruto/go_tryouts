package main

import (
	"io/ioutil"
	"crypto/aes"
	"io"
	"crypto/rand"
	"crypto/cipher"
	"os"
	"bytes"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"crypto/x509"
	"encoding/pem"

)

func read_keys()(interface{},[]byte) {
b, err:= ioutil.ReadFile("keystore1.p12")
if err!=nil {
fmt.Println(err)
return nil,nil
}
password := "2401940saa"
privk, pKey , err := pkcs12.Decode(b,password)
	derPubKey,err :=x509.MarshalPKIXPublicKey(pKey)
	if err != nil {
		fmt.Println(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPubKey,
	})
	fmt.Println(derPubKey)
	fmt.Println(len(derPubKey))

	ioutil.WriteFile("key.pub", pubBytes, 0644)
if err!=nil {
fmt.Println(err)
}
return privk,derPubKey
}

func main() {
	// read content from your file
	plaintext, err := ioutil.ReadFile("new.txt")
	if err != nil {
		panic(err.Error())
	}

	// this is a key
	_,key2 := read_keys()
	block, err := aes.NewCipher(key2)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewOFB(block,iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// create a new file for saving the encrypted data.
	f, err := os.Create("kupa_szyfr.txt")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}

	// done
}
