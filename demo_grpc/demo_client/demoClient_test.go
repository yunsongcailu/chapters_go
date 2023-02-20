package demo_client

import (
	"context"
	"google.golang.org/grpc/metadata"
	"testing"
	demopb "v0/demo_grpc/demo_proto/gen"
)

func TestDemoClient(t *testing.T) {
	pythonConn, pythonDemoClient, _ := DemoClient("0.0.0.0:5051")
	resp, _ := pythonDemoClient.SayHello(context.Background(), &demopb.HelloRequest{Name: "tom"})
	t.Log(resp)
	pythonConn.Close()
	goConn, goDemoClient, _ := DemoClient("0.0.0.0:5051")
	// grpc metadata
	// md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	md := metadata.New(map[string]string{
		"name":  "tom",
		"token": "123",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, _ = goDemoClient.SayHello(ctx, &demopb.HelloRequest{Name: "tom"})
	t.Log(resp)
	goConn.Close()
}
