package main

import (
	"fmt"
	"net/http"
	cfg "transfer-api-service/config"
	"transfer-api-service/handler"
)

func main() {
	// static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// interface
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/api/user/balance", handler.AuthInterceptor(handler.BalanceHandler))
	http.HandleFunc("/api/user/transaction", handler.AuthInterceptor(handler.TransactionHandler))
	http.HandleFunc("/api/user/transfer", handler.AuthInterceptor(handler.TransferHandler))

	fmt.Printf("Service start，listening [%s]...\n", cfg.ServiceHost)
	// Start Service
	err := http.ListenAndServe(cfg.ServiceHost, nil)
	if err != nil {
		fmt.Printf("Failed to start service, err:%s", err.Error())
	}
}
