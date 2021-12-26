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
go mod tidy # run this only once.

go run server/main.go
go run client/main.go
```

Sample Server logs

```log
2021/12/26 20:27:31 initializing grpc on port 9000
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:1
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:2
2021/12/26 20:27:37 instruction :  operator:"ADD"
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:3
2021/12/26 20:27:37 instruction :  operator:"DIV"
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:4
2021/12/26 20:27:37 instruction :  operator:"MUL"
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:5
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:6
2021/12/26 20:27:37 instruction :  operator:"SUB"
```

Sample Client logs

```log
2021/12/26 20:27:31 initializing grpc on port 9000
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:1
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:2
2021/12/26 20:27:37 instruction :  operator:"ADD"
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:3
2021/12/26 20:27:37 instruction :  operator:"DIV"
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:4
2021/12/26 20:27:37 instruction :  operator:"MUL"
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:5
2021/12/26 20:27:37 instruction :  operator:"PUSH" operand:6
2021/12/26 20:27:37 instruction :  operator:"SUB"
```

---

> [_rapando](https://twitter.com/rapando)