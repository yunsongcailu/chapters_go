package stream_client

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
	streampb "v0/demo_stream/stream_proto/stream_gen"
)

func StreamClient(addr string) (*grpc.ClientConn, streampb.StreamDemoServerClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	sdc := streampb.NewStreamDemoServerClient(conn)
	return conn, sdc, err
}

// StreamGetClient 服务端流模式-客户端
func StreamGetClient() {
	conn, sdc, _ := StreamClient("0.0.0.0:5052")
	gc, err := sdc.GetStream(context.Background(), &streampb.StreamRequest{Name: "tom"})
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	var quit = make(chan int)
	defer cancel()
	go func(c context.Context) {
		defer func() {
			quit <- 1
			close(quit)
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println(c.Err())
				return
			default:
				resp, err := gc.Recv()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(resp.Data)
			}
		}
	}(ctx)
	<-quit
	conn.Close()
	fmt.Println("quit")
}

// StreamPutClient 客户端流模式-客户端
func StreamPutClient() {
	ctx, cancel := context.WithCancel(context.Background())
	var err error
	quit := make(chan int)
	conn, sdc, _ := StreamClient("0.0.0.0:5052")
	pc, err := sdc.PutStream(ctx)
	go func(c context.Context, e *error) {
		defer func() {
			quit <- 1
			close(quit)
		}()
		for {
			select {
			case <-c.Done():
				return
			default:
				data := fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02 15:04:05"), "putClient")
				resErr := pc.Send(&streampb.StreamRequest{Name: data})
				if resErr != nil {
					e = &resErr
					return
				}
			}
			time.Sleep(time.Second)
		}
	}(ctx, &err)
	time.Sleep(time.Second * 10)
	cancel()
	<-quit
	conn.Close()
	fmt.Println("put client exit")
}

// StreamAllClient 双向流模式-客户端
func StreamAllClient() {
	ctx, cancel := context.WithCancel(context.Background())
	conn, sdc, _ := StreamClient("0.0.0.0:5052")
	ac, err := sdc.AllStream(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, err := ac.Recv()
			if err != nil {
				fmt.Println("client recv error", err)
				return
			}
			fmt.Println("recv:", data.Data)
		}
	}()

	go func() {
		defer func() {
			wg.Done()
			cancel()
		}()
		i := 0
		for {
			if i >= 20 {
				return
			}
			data := fmt.Sprintf("%s%d", "all client name:", i)
			if err := ac.Send(&streampb.StreamRequest{Name: data}); err != nil {
				fmt.Println(err)
				return
			}
			i++
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

	conn.Close()
}
