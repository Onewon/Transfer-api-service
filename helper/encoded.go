package main

import (
	"fmt"
	"transfer-api-service/util"
)

const (
	salt      = "*#FAFB"
	auth_salt = "*#FF"
)

func main() {
	password := "userpassword"
	encodedPsw := util.Sha1([]byte(salt + password))
	fmt.Println("Source: " + password)
	fmt.Println("Encoded: " + encodedPsw)

	auth := util.MD5([]byte(auth_salt + encodedPsw))
	fmt.Println("Auth: " + auth)
	//
	password = "userpassword2"
	encodedPsw = util.Sha1([]byte(salt + password))
	fmt.Println("Source: " + password)
	fmt.Println("Encoded: " + encodedPsw)

	auth = util.MD5([]byte(auth_salt + encodedPsw))
	fmt.Println("Auth: " + auth)
}

/*
Source: userpassword
Encoded: e36fef69cd47c627ef16830f0e424c15c56a8222
Auth: 0fa6a7b00d7322b757be311df22b5da3

Source: userpassword2
Encoded: 01e6d17645d155d5b5fbcafea8fceb59be9850ca
Auth: e3e35aa9ca3036f18c107fd30f37b9fe
*/
