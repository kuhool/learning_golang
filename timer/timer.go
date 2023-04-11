package timer

import (
	"fmt"
)

func Test() {
	//声明一个timer,timer是一次性的
	// timer1 := time.NewTimer(1 * time.Second) //定义了一个timer定时器
	// fmt.Println(time.Now())
	// t1 := <-timer1.C //等待1秒，然后打印当前时间
	// fmt.Println(t1)

	//验证timer只打印一次
	// timer2 := time.NewTimer(2 * time.Second)
	// fmt.Println(time.Now())
	// for {
	// 	t2 := <-timer2.C //只打印一次后，阻塞
	// 	fmt.Println(t2)
	// }

	//timer的延时功能(Sleep ,time.NewTimer,time.After)
	// fmt.Println(time.Now())
	// time.Sleep(time.Second)
	// fmt.Println(time.Now())
	// timer3 := time.NewTimer(1 * time.Second)
	// t3 := <-timer3.C
	// fmt.Println(t3)
	// <-time.After(1 * time.Second)
	// fmt.Println(time.Now())

	//停止定时器
	// timer4 := time.NewTimer(1 * time.Second)
	// go func() {
	// 	<-timer4.C
	// 	fmt.Println("定时器执行了")
	// }()
	// time.Sleep(2 * time.Second)
	// b := timer4.Stop()
	// if b {
	// 	fmt.Println("timer4已经关闭")
	// }

	// 5.重置定时器
	// timer5 := time.NewTimer(3 * time.Second)
	// timer5.Reset(1 * time.Second)
	// fmt.Println(time.Now())
	// fmt.Println(<-timer5.C)

	// 6.使用ticker,多次执行Ticker
	// ticker1 := time.NewTicker(2 * time.Second)
	// i := 0
	// go func() {
	// 	for {
	// 		i++
	// 		fmt.Println(<-ticker1.C)
	// 		if i == 10 {
	// 			ticker1.Stop()
	// 			return
	// 		}
	// 	}
	// }()

	// for {

	// }

	// ticker2 := time.NewTicker(2 * time.Second)
	// fmt.Println(<-ticker2.C)

	fmt.Println("timer.Test")
}
