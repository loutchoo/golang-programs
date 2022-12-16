package main

import (
	b64 "encoding/base64"
	"fmt"
)

func script() {
	var number int = 11
	var secondnumber int = 13
	fmt.Println(number, "+", secondnumber, "=", number+secondnumber)

	i := 1
	for i < 18 {
		fmt.Println("le nombre i est égal à :", i)
		i = i + 1
	}

	for nombre := 1; nombre <= 10; nombre++ {
		fmt.Println("Nombre :", nombre)
	}

	if num := number; num < 10 {
		fmt.Println("Le nombre est inférieur à 10 ! :D")
	} else {
		fmt.Println("Le nombre est supérieur à 10 ! :DDD")
	}

	test := 12
	fmt.Println("Write", test, "as")
	switch test {
	case 1:
		fmt.Println(test, "oeoeoeoe")
	case 2:
		fmt.Println("oeoeooee", test)
	case 3:
		fmt.Println("oe", test, "oe")
	}

	data := "oui zin c'est la base 64"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
}
