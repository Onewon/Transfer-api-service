# Funds Transfer Service API

## Description

An online API for funds transfer service in local machine. User can retrieve balance and transactions, and send money to other user. Once the transfer processing interrupt, the transaction will be rollback.



## Deployment

1.To run shell scripts to setup environment.
``` bash
# Start mysql docker container and redis docker container, to inject SQL file into MySQL。
cd Scripts/
sudo bash deployment.sh
```
2.Run main file to run service
``` bash
go run main.go
```
## API Reference
- [Retrieve Transactions](#retrieve-transactions)
  - [Description](#description)
  - [Request URL](#request-url)
  - [HTTP Request Method](#http-request-method)
  - [Parameters](#parameters)
  - [Response Representations](#response-representations)
- [Retrieve Balance](#retrieve-balance)
    - [Description](#description-1)
    - [Request URL](#request-url-1)
    - [HTTP Request Method](#http-request-method-1)
    - [Parameters](#parameters-1)
    - [Response Representations](#response-representations-1)
- [Funds Transfer](#funds-transfer)
    - [Description](#description-2)
    - [Request URL](#request-url-2)
    - [HTTP Request Method](#http-request-method-2)
    - [Parameters](#parameters-2)
    - [Response Representations](#response-representations-2)


## Retrieve Transactions

##### Description

- User Retrieve Transactions Interface

##### Request URL
- ` http://localhost:8080/api/user/transaction`

##### HTTP Request Method
- Get


##### Parameters

| Parameter name | Type   | Description       | Optional |
| :------------- | :----- | :---------------- | -------- |
| uid            | string | user name         | Not Null |
| auth           | string | **Authorization** | Not Null |
| bankCode       | string | Bank Code         | Not Null |
| accountNo      | string | Account Number    | Not Null |
| from           | string | IP address        | Not Null |
| env            | string | Device info(OS)   | Optional |
| page           | int    | Transaction Page  | Optional |

##### Response Representations

``` 
  {
    "error_code": 0,
    "data": {
      "uid": <User name>,
	  "requestTime": <Request Time>,
	  "transactions":[
	   {
	   "accountNo": <Account Number>,
	   "targetAccountNo": <Target Account Number>,
	   "amount": <Amount>,
	   "time" : <TransactionsTime>,
	   "status": <Transactions Status>
	   },
	   {
	   "accountNo": <Account Number>,
	   "targetAccountNo": <Target Account Number>,
	   "amount": <Amount>,
	   "time" : <TransactionsTime>,
	   "status": <Transactions Status>
	   },
	   ...
	  ]
	  "transactionPage": 1
    }
  }
```


## Retrieve Balance

##### Description

- User Retrieve Balance Interface

##### Request URL
- ` http://localhost:8080/api/user/balance`

##### HTTP Request Method
- Get

##### Parameters

| Parameter name | Type   | Description       | Optional |
| :------------- | :----- | :---------------- | -------- |
| uid            | string | user name         | Not Null |
| auth           | string | **Authorization** | Not Null |
| bankCode       | string | Bank Code         | Not Null |
| accountNo      | string | Account Number    | Not Null |
| from           | string | IP address        | Not Null |
| env            | string | Device info(OS)   | Optional |

##### Response Representations

``` 
  {
    "error_code": 0,
    "data": {
      "uid": <User name>,
	  "accountNo": <Account Number>,
      "balance": <Balance>,
	  "requestTime": <Request Time>,
    }
  }
```


| Property name | Type    | Description     |
| :------------ | :------ | :-------------- |
| error_code    | int     | Error code      |
| data          | -       | Response data   |
| uid           | string  | User name       |
| accountNo     | string  | Account number  |
| balance       | decimal | Account balance |
| requestTime   | date    | Request Time    |

## Funds Transfer

##### Description

- Funds Transfer Interface

##### Request URL
- ` http://localhost:8080/api/user/transfer `
  
##### HTTP Request Method
- Post

##### Parameters

| Parameter name  | Type    | Description           | Optional |
| :-------------- | :------ | :-------------------- | -------- |
| uid             | string  | user id               | Not Null |
| auth            | string  | **Authorization**     | Not Null |
| srcCode         | string  | Source bank Code      | Not Null |
| accountNo       | string  | Source account Number | Not Null |
| targetCode      | string  | Target bank Code      | Not Null |
| targetAccountNo | string  | Target account Number | Not Null |
| amount          | decimal | Transaction money     | Not Null |
| from            | string  | IP address            | Not Null |
| env             | string  | Device info(OS)       | Optional |
##### Response Representations

``` 
  {
    "error_code": 0,
    "data": {
      "responseTime": 100,
      "reqStatus": "success",
	  "src": {
		"serialNo":  <Transaction ID>,
		"srcCode": "BA001",
		"accountNo": <Source Card Number>,
		"targetCode": "BA002",
		"targetAccountNo": <Target Card Number>,
		"amount": <Amount>,
		"uid": <Userid>,
		"from": <Source IP>,
		"env": <Operating System Info>,
		"reqTime": <Request Time>,
	  }
    }
  }
```

| Property name | Type   | Description               |
| :------------ | :----- | :------------------------ |
| error_code    | int    | Error code                |
| data          | -      | Response data             |
| responseTime  | int    | Service response time(ms) |
| requestStatus | stirng | Request status            |
| src           | -      | Request data              |