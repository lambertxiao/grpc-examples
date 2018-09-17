package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address = "localhost:8080"
	defaultName = "world"
)

func main() {
	// 发起与服务端的连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}

	defer conn.Close()

	// 生成GRPC的client端
	client := pb.NewGreeterClient(conn)
	name := defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 客户端调用service方法
	resp, err := client.SayHello(ctx, &pb.HelloRequest{ Name: name })

	if err != nil {
		log.Fatalf("cannot greet: %v", err)
	}

	log.Printf("greeting: %s", resp.Message)
}
