package main

//compatible_(https://md5decrypt.net/en/Sha256/)

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("++++++ Encode SHA-256 ++++++")
	data := ""
	fmt.Print("Insert: ")
	fmt.Scanln(&data)
	fmt.Println("-----------------------------")
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)
	fmt.Print("Hash: ")
	fmt.Printf("%x\n", hash)
}
