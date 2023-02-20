package main

import (
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	demopb "v0/demo_grpc/demo_proto/gen"
	"v0/demo_grpc/demo_server"
	streampb "v0/demo_stream/stream_proto/stream_gen"
	"v0/demo_stream/stream_server"
)

func main() {
	// simpleRpcServer 简单一元模式服务端
	simpleRpcServer()
	// 流模式
	//streamServer()
}

// streamServer 流模式
func streamServer() {
	lis, _ := net.Listen("tcp", ":5052")
	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
	streampb.RegisterStreamDemoServerServer(s, &stream_server.StreamDemoServer{})
	go func() {
		fmt.Println("stream server : 5052")
		_ = s.Serve(lis)
	}()
	// 接收终止终端信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("stop")
}

// simpleRpcServer 简单一元模式服务端
func simpleRpcServer() {
	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
	demopb.RegisterGreeterServer(s, &demo_server.DemoService{})
	lis, _ := net.Listen("tcp", ":5051")
	go func() {
		fmt.Println("demo server : 5051")
		_ = s.Serve(lis)
	}()
	// 接收终止终端信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("stop")
}
