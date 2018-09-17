三分钟上手 GRPC (Go)
======================

背景
-------------
在这个例子里,我们已经通过 [helloworld.proto] 生成了服务端和客户端的桩文件
(helloworld/helloworld/helloworld.proto).

前提
-------------

- Go 1.6 或更高
- 已经配置好GOPATH


INSTALL
-------

```
$ go get -u google.golang.org/grpc/examples/helloworld/greeter_client
$ go get -u google.golang.org/grpc/examples/helloworld/greeter_server
```

TRY IT!
-------

- Run the server

  ```
  $ greeter_server &
  ```

- Run the client

  ```
  $ greeter_client
  ```

OPTIONAL - Rebuilding the generated code
----------------------------------------

1. Install [protobuf compiler](https://github.com/google/protobuf/blob/master/README.md#protocol-compiler-installation)

1. Install the protoc Go plugin

   ```
   $ go get -u github.com/golang/protobuf/protoc-gen-go
   ```

1. Rebuild the generated Go code

   ```
   $ go generate google.golang.org/grpc/examples/helloworld/...
   ```

   Or run `protoc` command (with the grpc plugin)

   ```
   $ protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
   ```
