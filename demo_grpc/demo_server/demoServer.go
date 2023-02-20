package demo_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	demopb "v0/demo_grpc/demo_proto/gen"
)

type DemoService struct {
	demopb.UnimplementedGreeterServer
}

func (ds *DemoService) SayHello(ctx context.Context, in *demopb.HelloRequest) (*demopb.HelloReply, error) {
	// 去除metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	for i, v := range md {
		fmt.Printf("metadata:k-%s,v-%s\n", i, v)
	}
	return &demopb.HelloReply{Message: "你好" + in.Name + "这里是golang demo"}, nil
}
