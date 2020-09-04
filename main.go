package main

import (
	"fmt"
	"net/http"
	cfg "transfer-api-service/config"
	"transfer-api-service/handler"
)

func main() {
	// 静态资源处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// 接口
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/api/user/balance", handler.BalanceHandler)
	http.HandleFunc("/api/user/transaction", handler.TransactionHandler)
	http.HandleFunc("/api/user/transfer", handler.TransferHandler)

	// 文件存取接口
	// http.HandleFunc("/file/upload", handler.HTTPInterceptor(handler.UploadHandler))

	fmt.Printf("Service start，listening [%s]...\n", cfg.ServiceHost)
	// Start Service
	err := http.ListenAndServe(cfg.ServiceHost, nil)
	if err != nil {
		fmt.Printf("Failed to start service, err:%s", err.Error())
	}
}
