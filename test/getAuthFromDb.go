package main

/*
Run sql script to get user certification.
*/
import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MySQLSource = "root:112233@tcp(192.168.80.128:3314)/maybankdb?charset=utf8"
)

func connDb() (DB *sql.DB) {
	DB, err := sql.Open("mysql", MySQLSource)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	return DB
}

func getAuth(s string) string {
	//Params
	//p1 := ""

	//Connect to MySQL
	db := connDb()

	q, err := ioutil.ReadFile("../sql/auth.sql")
	if err != nil {
		panic(err)
	}
	query := string(q[:])
	var auth string
	err = db.QueryRow(query, s).Scan(&auth)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Auth: " + auth)
	return auth
}
