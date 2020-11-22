package main

/*
Desp:   Transfer Interface
Method: Post
command: go run transfer_money.go getAuthFromDb.go
*/

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	APIHost = "localhost:8989"
)

func main() {
	username := "user_B"
	auth := getAuth(username)
	// 1.normal tsc A-B
	// params := url.Values{"uid": {"user_B"},
	// 	"auth":            {"e3e35aa9ca3036f18c107fd30f37b9fe"},
	// 	"accountNo":       {"101120223032"},
	// 	"targetAccountNo": {"101120223031"},
	// 	"amount":          {"5"},
	// 	"from":            {"192.168.0.101"},
	// 	"env":             {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36"}}

	// var expect string
	// expect = ``

	// resp, err := http.PostForm(fmt.Sprintf("http://%s/api/user/transfer", APIHost), params)
	// if err != nil {
	// 	println(err.Error())
	// }
	// defer resp.Body.Close()
	// ret, _ := ioutil.ReadAll(resp.Body)

	// actual := string(ret)
	// if actual != string(expect) {
	// 	fmt.Printf("Val:%s\n", params)
	// 	fmt.Printf("Actual:%s\n", actual)
	// 	fmt.Printf("Expect:%s\n", expect)
	// } else {
	// 	fmt.Println("OK")
	// }

	// 2. normal tsc B-A
	/*
		params := url.Values{"uid": {"user_A"},
			"auth":            {"0fa6a7b00d7322b757be311df22b5da3"},
			"accountNo":       {"101120223031"},
			"targetAccountNo": {"101120223032"},
			"amount":          {"5"},
			"from":            {"192.168.0.101"},
			"env":             {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36"}}

		var expect string
		expect = ``

		resp, err := http.PostForm(fmt.Sprintf("http://%s/api/user/transfer", APIHost), params)
		if err != nil {
			println(err.Error())
		}
		defer resp.Body.Close()
		ret, _ := ioutil.ReadAll(resp.Body)

		actual := string(ret)
		if actual != string(expect) {
			fmt.Printf("Val:%s\n", params)
			fmt.Printf("Actual:%s\n", actual)
			fmt.Printf("Expect:%s\n", expect)
		} else {
			fmt.Println("OK")
		}*/

	// 3.my account wrong or balance not enough
	/*
		params := url.Values{"uid": {"user_B"},
			"auth":            {"e3e35aa9ca3036f18c107fd30f37b9fe"},
			"accountNo":       {"101120223033"},
			"targetAccountNo": {"101120223031"},
			"amount":          {"20"},
			"from":            {"192.168.0.101"},
			"env":             {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36"}}

		var expect string
		expect = ``

		resp, err := http.PostForm(fmt.Sprintf("http://%s/api/user/transfer", APIHost), params)
		if err != nil {
			println(err.Error())
		}
		defer resp.Body.Close()
		ret, _ := ioutil.ReadAll(resp.Body)

		actual := string(ret)
		if actual != string(expect) {
			fmt.Printf("Val:%s\n", params)
			fmt.Printf("Actual:%s\n", actual)
			fmt.Printf("Expect:%s\n", expect)
		} else {
			fmt.Println("OK")
		}
	*/

	//4. target account wrong and rollback
	params := url.Values{
		"uid":             {username},
		"auth":            {auth},
		"accountNo":       {"101120223032"},
		"targetAccountNo": {"101120223031"}, //available target account
		// "targetAccountNo": {"101120223035"}, //wrong target account
		"amount": {"10"},
		"from":   {"192.168.0.101"},
		"env":    {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36"}}

	var expect string
	expect = ``

	resp, err := http.PostForm(fmt.Sprintf("http://%s/api/user/transfer", APIHost), params)
	if err != nil {
		println(err.Error())
	}
	defer resp.Body.Close()
	ret, _ := ioutil.ReadAll(resp.Body)

	actual := string(ret)
	if actual != string(expect) {
		fmt.Printf("Val:%s\n", params)
		fmt.Printf("Actual:%s\n", actual)
		fmt.Printf("Expect:%s\n", expect)
	} else {
		fmt.Println("OK")
	}
}
