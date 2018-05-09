package main

import (
	"fmt"
	"github.com/square/certigo/jceks"
)

/*
func readKeyStore(filename string, password []byte) keystore.KeyStore {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	keyStore, err := keystore.Decode(f, password)
	if err != nil {
		log.Fatal(err)
	}
	return keyStore
}*/
func main() {
	//b, err:= ioutil.ReadFile("keystore.jceks")
	//if err!=nil {
	//	fmt.Println(err)
	//	return
	//}
	password := "2401940saa"
	keystore,err:=jceks.LoadFromFile("keystore.jceks",[]byte(password))
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(keystore)
	/*if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(pKey)*/
}
