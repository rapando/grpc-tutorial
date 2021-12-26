## Go gRPC

inspiration from 

- [talentica](https://talentica.com/a-quick-overview-of-grpc-in-golang/)

- [grpc.io](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)

---


To generate the proto code from the proto file,

```sh
protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    machine/machine.proto
```

---

To run:

```sh
go run server/main.go
go run client/main.go