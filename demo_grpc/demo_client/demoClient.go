package demo_client

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	demopb "v0/demo_grpc/demo_proto/gen"
)

func DemoClient(addr string) (*grpc.ClientConn, demopb.GreeterClient, error) {
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	dc := demopb.NewGreeterClient(conn)
	return conn, dc, nil
}
