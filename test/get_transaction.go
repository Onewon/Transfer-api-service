package main

/*
Desp:   Retrieve Transaction
Method: Get
command: go run get_transaction.go getAuthFromDb.go
*/

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	APIHost = "localhost:8080"
)

func main() {
	username := "user_B"
	auth := getAuth(username) //"e3e35aa9ca3036f18c107fd30f37b9fe"
	account := "101120223032"
	page := ""
	ip := "192.168.0.101"
	env := "Windows10_64bit"
	url := fmt.Sprintf("http://%s//api/user/transaction?uid=%s&auth=%s&accountNo=%s&page=%s&from=%s&env=%s",
		APIHost, username, auth, account, page, ip, env)
	fmt.Println(url)
	var expect string
	expect = `{"UserId":"user_B","TransactionsList":[{"TransactionID":"tscdec031b5a9dff74a83018b53a08e06de","AccountNo":"101120223032","TargetAccountNo":"101120223031","Amount":5,"Time":"2020-09-04 05:56:31","Status":"1"},{"TransactionID":"tsc1a4b58358e692a0512a609841e668474","AccountNo":"101120223032","TargetAccountNo":"101120223031","Amount":5,"Time":"2020-09-04 05:56:34","Status":"1"}],"TransactionPage":1,"IP":"192.168.0.101","Environment":"Windows10_64bit","RequestTime":"2020-09-04 02:04:21"}`

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	actual := string(s)

	if actual != string(expect) {
		fmt.Printf("Val:%s\n", url)
		fmt.Printf("Actual:%s\n", actual)
		fmt.Printf("Expect:%s\n", expect)
	} else {
		fmt.Println("OK")
	}
}
