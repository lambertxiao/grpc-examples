package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

// server结构体用来实现 helloworld.GreeterServer 接口的具体操作
type server struct {}

// 实现 GreeterServer 的 SayHello 方法的具体逻辑
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello" + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// 需要两个参数, grpc service 和具体的实现
	pb.RegisterGreeterServer(s, &server{})
	// 在给定的 grpc service 上 注册反射服务
	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
