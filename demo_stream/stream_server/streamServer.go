package stream_server

import (
	"context"
	"fmt"
	"sync"
	"time"
	streampb "v0/demo_stream/stream_proto/stream_gen"
)

type StreamDemoServer struct {
	streampb.UnimplementedStreamDemoServerServer
}

// GetStream 服务端流模式示例 关键词 stream
func (sds *StreamDemoServer) GetStream(req *streampb.StreamRequest, res streampb.StreamDemoServer_GetStreamServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	var err error
	quit := make(chan int)
	go func(c context.Context, e *error) {
		for {
			select {
			case <-c.Done():
				return
			default:
				data := fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02 15:04:05"), req.Name)
				resErr := res.Send(&streampb.StreamResponse{Data: data})
				if resErr != nil {
					e = &resErr
					quit <- 1
					close(quit)
				}
			}
			time.Sleep(time.Second)
		}
	}(ctx, &err)
	<-quit
	cancel()
	return err
}

// PutStream 客户端流模式
func (sds *StreamDemoServer) PutStream(req streampb.StreamDemoServer_PutStreamServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	var quit = make(chan int)
	var err error
	defer cancel()
	go func(c context.Context, e *error) {
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
				resp, err := req.Recv()
				if err != nil {
					e = &err
					return
				}
				fmt.Println(resp.Name)
			}
		}
	}(ctx, &err)
	<-quit
	return err
}

// AllStream 双向流模式
func (sds *StreamDemoServer) AllStream(req streampb.StreamDemoServer_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var recvChan = make(chan string, 10)
	go func() {
		defer func() {
			close(recvChan)
			wg.Done()
		}()
		for {
			data, err := req.Recv()
			if err != nil {
				fmt.Println("recv error", err)
				return
			}
			fmt.Println("recv:", data.Name)
			recvChan <- data.Name
		}
	}()
	go func() {
		defer wg.Done()
		for m := range recvChan {
			data := fmt.Sprintf("%s - %s", "server say:", m)
			err := req.Send(&streampb.StreamResponse{Data: data})
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()

	wg.Wait()
	return nil
}
