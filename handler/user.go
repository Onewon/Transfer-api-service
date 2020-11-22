package handler

import (
	"fmt"
	db "transfer-api-service/db/mysql"
)

func IsAuthValid(uid string, auth string) bool {
	stmt, err := db.DBConn().Prepare("select user_auth from tbl_user_auth " +
		"where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	var auth_tmp string
	err = stmt.QueryRow(uid).Scan(&auth_tmp)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if auth != auth_tmp {
		fmt.Println("Authorization is not allowed.")
		return false
	}
	return true
}
