package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Lixxcn/demo/gRPC-demo/hello" // 替换为你的 proto 文件路径

	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewGreeterClient(conn)

	// 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用 SayHello 方法
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Lixx"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印响应
	log.Printf("Greeting: %s", r.GetMessage())
}
