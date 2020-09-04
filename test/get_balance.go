package main

/*
Desp:   Retrieve Balance
Method: Get
command: go run get_balance.go getAuthFromDb.go
*/

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "testing"
)

const (
	APIHost = "localhost:8080"
)

func main() {
	username := "user_A"
	auth := getAuth(username) //"0fa6a7b00d7322b757be311df22b5da3"
	account := "101120223031"
	url := fmt.Sprintf("http://%s/api/user/balance?uid=%s&auth=%s&accountNo=%s",
		APIHost, username, auth, account)

	var expect string
	expect = `{"UserId":"user_A","AccountNo":"101120223031","AccountBalance":100}`

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, _ := ioutil.ReadAll(resp.Body)
	actual := string(s)

	if actual != string(expect) {
		fmt.Printf("Val:%s\n", url)
		fmt.Printf("Actual:%s\n", actual)
		fmt.Printf("Expect:%s\n", expect)
	} else {
		fmt.Println("OK")
	}

	username = "user_B"
	auth = "e3e35aa9ca3036f18c107fd30f37b9fe"
	account = "101120223032"
	url = fmt.Sprintf("http://%s/api/user/balance?uid=%s&auth=%s&accountNo=%s",
		APIHost, username, auth, account)

	expect = `{"UserId":"user_B","AccountNo":"101120223032","AccountBalance":100}`

	resp, err = http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, _ = ioutil.ReadAll(resp.Body)
	actual = string(s)

	if actual != string(expect) {
		fmt.Printf("Val:%s\n", url)
		fmt.Printf("Actual:%s\n", actual)
		fmt.Printf("Expect:%s\n", expect)
	} else {
		fmt.Println("OK")
	}
}
