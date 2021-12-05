# Go-GRPC playground
---
This Project is a simple demo for grpc streaming communication between grpc server and client.
### Generate grpc code via protoc
```
protoc -I. --go-grpc_out=.\ --go_out=.\ .\protos\PropleInfo.proto
```

### Start grpc server
```
go run main.go
```

### Sending data into grpc server
```
cd client/
go run client.go
```
