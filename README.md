# Funds Transfer Service API

## Learned by Project
1 多个docker脚本部署 -> 之后可以再学一下dockerfile
2 既然写了shell脚本，-> 学习写下python脚本
3 学会了直接使用echo追加字符到文本，不再依赖于文件了。
4 思路涉及到了分布式的部署，-> 可以学下 Docker Compose
5 Golang学会使用defer来取消某种操作
6 知道了linux 直接交互命令和后台交互的区别，直接交互的命令之后是无法自动化的
7 知道了golang 还有专门testing的写法
8 知道了golang 还需要logging -> 学习下怎么正规的logging
9 知道了golang还有各种语法限制：强制struct大写命名，某些小写命令无法通过。
实例化的对象大写命名好像也有问题。
10 具体实践了数据库的悲观锁(走索引的行锁,不走索引的表锁)和乐观锁(逻辑上尽量避免加锁)
2.1 对于系统操作的业务，我倾向于用shell awk sed 等系统自带的工具，
像要跑业务，跑任务，那可以考虑用python ，例如多线程任务，excel ，pdf 报表任务等


## Want To Do
- [x] Auth拦截分层 并把handler都分离了
- [ ] URLEncoding 标准库有坑
- [ ] dockerfile
- [ ] 修改为Python 部署脚本
- [ ] Docker Compose 使用
- [ ] handler中还差个logging
- [ ] 学会写 testing
- [ ] 域名重定向到localhost
重构
[请求处理层] 鉴权 auth
[业务逻辑层] query and transfer logic
[通用处理层] Manager -> helper ->限定于特定逻辑
（如JSON response struct可以继承）
[数据持久层] 对接数据源

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