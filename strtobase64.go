package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Bienvenue dans le programme string vers base 64 :)")

	data := "petit message pour les anciens"

	encode := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encode)

	decode, _ := b64.StdEncoding.DecodeString(encode)
	fmt.Println(string(decode))
}
