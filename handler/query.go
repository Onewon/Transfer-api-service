package handler

/*
Query for balance and transaction
*/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "transfer-api-service/db/mysql"
)

func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	type BalanceQueryResponse struct {
		UserId         string
		AccountNo      string
		AccountBalance float64
	}
	if r.Method == http.MethodGet {
		//get parameters
		vars := r.URL.Query()
		user_id := vars.Get("uid")
		acc_no := vars.Get("accountNo")

		query := BalanceQueryResponse{}

		stmt, err := db.DBConn().Prepare(
			"select account_balance from tbl_user_savings " +
				"where user_name=? and account_number =? limit 1")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer stmt.Close()
		query.UserId = user_id
		query.AccountNo = acc_no
		err = stmt.QueryRow(user_id, acc_no).Scan(&query.AccountBalance)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}

		ret_data, err := json.Marshal(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ret_data)
		return
	}
}

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	type TransactionRecord struct {
		TransactionID   string
		AccountNo       string
		TargetAccountNo string
		Amount          float64
		Time            string
		Status          string
	}
	type TransactionQueryResponse struct {
		UserId           string
		TransactionsList []TransactionRecord
		TransactionPage  int64
		IP               string
		Environment      string
		RequestTime      string
	}

	if r.Method == http.MethodGet {
		//1. get parameters
		vars := r.URL.Query()
		user_id := vars.Get("uid")
		acc_no := vars.Get("accountNo")
		ip_source := vars.Get("from")
		env := vars.Get("env")
		st_date := vars.Get("sdate")
		end_date := vars.Get("edate")
		page := vars.Get("page")
		var tsc_page int64
		tsc_page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		if len(page) == 0 {
			tsc_page = 1
		}

		reques_time := time.Now().Format("2006-01-02 03:04:05")
		transaction_query := TransactionQueryResponse{}

		stmt, err := db.DBConn().Prepare("select transaction_ID,account_number," +
			"target_account_number,amount,status,last_update from tbl_user_transaction " +
			"where user_name=? and last_update between ? and ? and account_number=? limit ?,?")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer stmt.Close()

		//2. get all transactions
		var st_index int64
		if tsc_page == 1 {
			st_index = 0
		} else {
			st_index = 5 * (tsc_page - 1)
		}
		offset := 5
		rows, err := stmt.Query(user_id, st_date, end_date, acc_no, st_index, offset)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()
		if !rows.Next() {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("No result")
			return
		}
		transactionRecordRows := []TransactionRecord{}
		for rows.Next() {
			var tmp TransactionRecord
			err = rows.Scan(&tmp.TransactionID, &tmp.AccountNo,
				&tmp.TargetAccountNo, &tmp.Amount, &tmp.Status, &tmp.Time)
			if err != nil {
				fmt.Println(err.Error())
			}
			transactionRecordRows = append(transactionRecordRows, tmp)
		}
		err = rows.Err()
		if err != nil {
			fmt.Println(err.Error())
		}

		//3. load params
		transaction_query.UserId = user_id
		transaction_query.IP = ip_source
		transaction_query.Environment = env
		transaction_query.RequestTime = reques_time
		transaction_query.TransactionPage = tsc_page
		transaction_query.TransactionsList = transactionRecordRows

		//4. response
		ret_data, err := json.Marshal(transaction_query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ret_data)
		return
	}
}
