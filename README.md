# datadome

## implementation
the current structure is below and the main idea is : 

- create a entrypoint to handle the /cart on main.go
- call the module.go to analyse the http request
- module.go will parse the request and send the relevant headers to the api_server.go
- the api_server.go would return a value from AWS, for instance (TODO: mock)


tree ./datadome

├── README.md

├── api_server

├──── api_server.go

├── customer

├──── customer.go

├── go.mod

├── main.go

└── module

├──── module.go


## build / run
```
#you should have go installed in your machine

cd $HOME/go/src #or to your configured $GOPATH
git clone https://github.com/laurogr/datadome.git
cd datadome
go run main.go
#open your browser and open : localhost:8081/cart
#some logs will be displayed with the steps that were executed, for instance :

[MAIN] /cart
[MODULE] AnalyzeHttpRequest
[API_SERVER] AnalyzeHttpRequest : Calls remote server : {"method":"GET","host":"localhost:8081","requestURI":"/cart","ip":"[::1]:60516","referer":"","useragent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Safari/605.1.15"}
[CUSTOMER] ACESS TO CART GRANTED, CONTINUE WORKFLOW WITH ORIGINAL HTTP REQUEST
```

