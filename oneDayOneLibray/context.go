package oneDayOneLibray

import (
	context "context"
	"fmt"
	"time"
)

// 参考 https://blog.csdn.net/shenziheng1/article/details/113924703

func CtcHttpTest() {

}
func CtxTimeOutTest() {
	//创建了一个1s超时的content
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func CtxTest() {

	ctx := context.Background()
	fmt.Println(ctx)
	ctx = context.WithValue(ctx, "key", "value")
	fmt.Println(ctx)
	val, ok := ctx.Value("key").(string)
	if ok {
		fmt.Println(val)
	}

}

// ContextTest 主协程发消息通知子协程任务的是否中断
func ContextTest() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)
	//consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for t := range ticker.C {
			fmt.Println("ticker.C------->", t) //golang时间戳
			select {
			case <-done:
				fmt.Println("childprocessinterrupt...")
				return
			default:
				fmt.Printf("sendmessage:%d\n", <-messages)
			}
		}
	}()

	//producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	//发送任务中断操作
	done <- true
	time.Sleep(1 * time.Second)
	close(done)

	fmt.Println("mainprocessexit!")
}
