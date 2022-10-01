# grpc_test
Test project for getting to know the gRPC in Go

<br />

# Get Started:
```
git clone https://github.com/Murchik/grpc_test.git grpc_test
cd grpc_test
```

## To start the server:
```
go run databus_server/main.go 8080 action
```
Avalibale actions: <br />
"mul" - multiplication <br />
"div" - division <br />
"sub" - subtraction <br />
"add" - addition

<br />

## To run the client:
```
go run databus_client/main.go 127.0.0.1:8080 param1 param2
```
Params are floating point numbers (float32)
