package pkg_context

import (
	"context"
	"fmt"
	"time"
)

func DemoContext() {
	fmt.Println("context上下文, Backgroud,WithCancel,WithDeadline,WithTimeout,WithValue")
	// 空上下文
	fmt.Println("创建空上下文c := context.Backgroud()")
	c := context.Background()
	n := ""
	for {
		quit := false
		fmt.Printf("WithCancel[1],WithDeadline[2],withTimeout[3],withValue[4],QUIT:[q]\n")
		fmt.Scanln(&n)
		switch n {
		case "q":
			quit = true
			break
		case "1":
			cancelCtx(c)
		case "2":
			withDeadline(c)
		case "3":
			withTimeout(c)
		case "4":
			withValue(c)
		}
		if quit {
			break
		}
	}
	fmt.Println("end")
}

// 可取消的
func cancelCtx(c context.Context) {
	fmt.Println("在c的背景下,创建一个可取消的context:ctx,cancel := context.WithCancel(c)")
	ctx, cancel := context.WithCancel(c)
	defer cancel()
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done(): // 当 n == 5 时break,主进程退出 触发defer cancel(), 触发ctx.Done()优雅退出 协程
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	dst := gen(ctx)
	for n := range dst {
		fmt.Printf("dst chan:%d\n", n)
		if n == 5 {
			fmt.Printf("当 n == %d 时break,主进程退出 触发defer cancel(), 触发ctx.Done()优雅退出 协程\n", n)
			break
		}
	}
}

// 截止日期的
func withDeadline(c context.Context) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(c, d)
	// 尽管有最后期限,最好还是defer 主动关闭
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("time.After 1 second overslept")
	case <-ctx.Done():
		fmt.Printf("time 50 millisecond ctx error:%v\n", ctx.Err())
	}
	return
}

// 超时的
func withTimeout(c context.Context) {
	ctx, cancel := context.WithTimeout(c, 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(time.Second):
		fmt.Printf("second time overslept")
	case <-ctx.Done():
		fmt.Printf("withTimeout 50 millisecond,ctx error:%v\n", ctx.Err())
	}
	return
}

// 赋值
func withValue(c context.Context) {
	fmt.Printf("withValue key 最好是自定义数据类型\n")
	type cKey string
	k := cKey("lang")
	ctx := context.WithValue(c, k, "Go")
	f := func(ctx context.Context, k cKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
		}
		fmt.Println("key not found")
	}

	f(ctx, k)
	f(ctx, cKey("color"))
}
