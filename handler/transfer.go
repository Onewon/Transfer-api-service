package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"

	db "transfer-api-service/db/mysql"
	redisPool "transfer-api-service/db/redis"
	util "transfer-api-service/util"
)

const (
	TRANSACTION_PENDING_STATUS = 0
	TRANSACTION_SUCCESS_STATUS = 1
	TRANSACTION_FAIL_STATUS    = 2
	TRANSACTION_SET            = "transaction_collection"
	TRANSACTION_SECRET         = "#RXSZGT"
)

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	type TransferRequest struct {
		TransactionID   string
		UserId          string
		AccountNo       string
		TargetAccountNo string
		Amount          float64
		IP              string
		Environment     string
		RequestTime     string
	}
	type TransferResponse struct {
		ResponseTime      string
		TransactionStatus string
		Data              TransferRequest
	}

	if r.Method == http.MethodGet {
		w.Write([]byte("<html><body><h1>Use Post Method to Link.</h1></body></html>"))
		return
	}

	// Post
	if r.Method == http.MethodPost {
		r.ParseForm()
		user_id := r.Form.Get("uid")
		account_no := r.Form.Get("accountNo")
		target_account_no := r.Form.Get("targetAccountNo")
		ip_source := r.Form.Get("from")
		env := r.Form.Get("env")
		transfer_amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)
		//1. check amount available
		if transfer_amount <= 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("transfering amount is unavailable.")
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}

		//2. load the params
		var reques_time string
		reques_time = time.Now().Format("2006-01-02 03:04:05")
		transfer_request := TransferRequest{"", user_id, account_no, target_account_no, transfer_amount, ip_source, env, reques_time}
		//3. Check Balance Vaild
		stmt, err := db.DBConn().Prepare("select account_balance from tbl_user_savings " +
			"where account_number=? limit 1")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer stmt.Close()
		var balance float64
		err = stmt.QueryRow(transfer_request.AccountNo).Scan(&balance)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		if balance < transfer_request.Amount {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Account balance not enough.")
			return
		}
		//4 Update the balance
		stmt, err = db.DBConn().Prepare("update tbl_user_savings " +
			"set account_balance=account_balance-? where account_number=? limit 1 ")
		if err != nil {
			fmt.Println("Prepared sql fail, err:" + err.Error())
			return
		}
		defer stmt.Close()
		ret, err := stmt.Exec(transfer_request.Amount, transfer_request.AccountNo)
		if err != nil {
			fmt.Println("Sql cannot execuse, err:" + err.Error())
			return
		}
		if rf, err := ret.RowsAffected(); nil == err {
			if rf <= 0 {
				fmt.Println("Fail to update balance.")
			}
			fmt.Println("Update balance successfully.")
		}

		//5.Write Transaction Record table
		stmt, err = db.DBConn().Prepare(
			"insert ignore into tbl_user_transaction (`transaction_ID`,`user_name`,`account_number`, `target_account_number`,`amount`,`status`)" +
				"values (?,?,?,?,?,?)")
		if err != nil {
			fmt.Println("Failed to prepare statement, err:" + err.Error())
		}
		defer stmt.Close()

		// generate tsc ID logic
		// "tsc"+hash(timestamp[6] + secret)
		var tID string
		timestamp := fmt.Sprintf("%x", time.Now().Unix())
		tID = "tsc" + util.MD5([]byte(util.Reverse(timestamp)[:6]+TRANSACTION_SECRET))

		transfer_request.TransactionID = tID
		ret, err = stmt.Exec(transfer_request.TransactionID, transfer_request.UserId, transfer_request.AccountNo, transfer_request.TargetAccountNo,
			transfer_request.Amount, TRANSACTION_PENDING_STATUS)
		if err != nil {
			fmt.Println(err.Error())
		}
		if rf, err := ret.RowsAffected(); nil == err {
			if rf <= 0 {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("Fail to record transaction %s into table.\n", transfer_request.TransactionID)
			}
		}
		defer rollbackTransaction(transfer_request.TransactionID)

		// 6. Request target account, use redis filter tsc
		rdsConn := redisPool.RedisPool().Get()
		defer rdsConn.Close()
		value, err := redis.Int64(rdsConn.Do("sismember", TRANSACTION_SET,
			transfer_request.TransactionID))
		if err != nil {
			fmt.Println("Fail to execute redis command.", err.Error())
		} else if value == 1 {
			fmt.Printf("Transaction %s has exist.\n", transfer_request.TransactionID)
			return
		}
		rdsConn.Do("SADD", TRANSACTION_SET, transfer_request.TransactionID)

		// 7. Check Target Account Vaild, inactive then rollback
		stmt, err = db.DBConn().Prepare("select account_status from tbl_user_savings " +
			"where account_number=? limit 1")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer stmt.Close()

		var tg_account_status int
		err = stmt.QueryRow(transfer_request.TargetAccountNo).Scan(&tg_account_status)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Target Account isn't exist.")
				return
			}
			fmt.Println(err.Error())
			return
		}
		if tg_account_status == 1 {
			fmt.Println("Target Account is inactive.")
			return
		}

		stmt, err = db.DBConn().Prepare("update tbl_user_savings " +
			"set account_balance=account_balance+? where account_number=? limit 1 ")
		if err != nil {
			fmt.Println("Prepared sql fail, err:" + err.Error())
			return
		}
		defer stmt.Close()
		ret, err = stmt.Exec(transfer_request.Amount, transfer_request.TargetAccountNo)
		if err != nil {
			fmt.Println("Sql cannot execuse, err:" + err.Error())
			return
		}
		if rf, err := ret.RowsAffected(); nil == err {
			if rf <= 0 {
				fmt.Println("Fail to finish transaction.")
				return
			}
			fmt.Println("Update target balance successfully.")
		}
		// 8. Update Transaction status in table
		stmt, err = db.DBConn().Prepare(
			fmt.Sprintf("update tbl_user_transaction set status= %d"+
				" where transaction_ID=? limit 1 ", TRANSACTION_SUCCESS_STATUS))
		if err != nil {
			fmt.Println("Failed to prepare statement, err:" + err.Error())
		}
		defer stmt.Close()

		ret, err = stmt.Exec(transfer_request.TransactionID)
		if err != nil {
			fmt.Println(err.Error())
		}
		if rf, err := ret.RowsAffected(); nil == err {
			if rf <= 0 {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("Fail to update status of transaction %s into table.\n", transfer_request.TransactionID)
			}
		}

		// 9 Response
		var response_time string
		var transaction_status string
		response_time = "10ms"
		transaction_status = "success"
		retResp := TransferResponse{
			response_time,
			transaction_status,
			transfer_request}

		ret_data, err := json.Marshal(retResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ret_data)
		return
	}
}

func rollbackTransaction(tsc string) {
	stmt, err := db.DBConn().Prepare("select user_name,account_number,amount,status from tbl_user_transaction " +
		"where transaction_ID=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	var uid string
	var account_no string
	var amount float64
	var tsc_status int
	err = stmt.QueryRow(tsc).Scan(&uid, &account_no, &amount, &tsc_status)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Transaction isn't exist.")
			return
		}
		fmt.Println(err.Error())
		return
	}
	// transaction no response
	if tsc_status == TRANSACTION_PENDING_STATUS {
		// rollback balance
		stmt, err = db.DBConn().Prepare("update tbl_user_savings " +
			"set account_balance=account_balance+? where account_number=? limit 1 ")
		if err != nil {
			fmt.Println("Prepared sql fail, err:" + err.Error())
			return
		}
		defer stmt.Close()
		ret, err := stmt.Exec(amount, account_no)
		if err != nil {
			fmt.Println("rollback Sql cannot execuse, err:" + err.Error())
			return
		}
		if rf, err := ret.RowsAffected(); nil == err {
			if rf <= 0 {
				fmt.Println("Fail to rollback balance.")
			}
			fmt.Println("Rollback balance successfully.")
		}
		// update transaction table
		stmt, err = db.DBConn().Prepare("update tbl_user_transaction " +
			"set status=? where transaction_ID=? limit 1 ")
		if err != nil {
			fmt.Println("Prepared sql fail, err:" + err.Error())
			return
		}
		defer stmt.Close()
		ret, err = stmt.Exec(TRANSACTION_FAIL_STATUS, tsc)
		if err != nil {
			fmt.Println("update transaction Sql cannot execuse, err:" + err.Error())
			return
		}
		if rf, err := ret.RowsAffected(); nil == err {
			if rf <= 0 {
				fmt.Println("Fail to update transaction table.")
			}
			fmt.Println("Update transaction table successfully.")
		}
	}
}
