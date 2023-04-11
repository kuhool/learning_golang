package selectT

import (
	"fmt"
	"time"
)

func Test() {

	// //向管道中输出数据
	// // 2个管道
	// output1 := make(chan string)
	// output2 := make(chan string)
	// // 跑2个子协程，写数据
	// go test1(output1)
	// go test2(output2)
	// // 用select监控
	// //select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句
	// select { //只执行一个case
	// case s1 := <-output1:
	// 	fmt.Println("s1=", s1)
	// case s2 := <-output2:
	// 	fmt.Println("s2=", s2)
	// }

	// time.Sleep(time.Second * 10)

	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("selectT.Test")
}

func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("add write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
func test1(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "test2"
}
