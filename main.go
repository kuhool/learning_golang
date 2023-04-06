package main

import (
	"fmt"
	"time"
)

func main() {

	var arr []int = []int{9: 10, 8: 10}
	fmt.Println(arr) // [9 10]

	t := time.Now() //Now()格式不能直接使用需要转换
	fmt.Println(t)
	t1 := t.Unix() //获取时间戳
	fmt.Println(t1)
	t = t.Add(time.Hour * 24) //时间相加一天
	fmt.Println(t)
	t2 := t.Format("2006-01-02 15:04:05") //格式化Y-m-d H:i:s
	fmt.Println(t2)

	//指定时间戳转换成时间
	timeStamp := 1680602716
	t3 := time.Unix(int64(timeStamp), 0)
	fmt.Println(t3)

	timeStamp = 1569826535             // 2019-09-30 14:55:35
	t = time.Unix(int64(timeStamp), 0) // time.Time
	year, month, day := t.Date()

	hour, minute, second := t.Clock()

	fmt.Println(year, month, day, hour, minute, second)

	// 同样可以根据时间格式提供的方法
	// 可以获取星期几，还可以获取当前是一年的多少天
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Weekday(), t.YearDay())

	fmt.Println(time.Now().Format("1"))

	// 判断某个时间是不是在某个时间之前/之后/相等，传入的参数必须是时间格式
	t4 := t.After(time.Now())
	t5 := t.Before(time.Now())
	t6 := t.Equal(t)
	// 返回的是bool值
	fmt.Println(t4, t5, t6) // false true true
	fmt.Println("Hello World!")

}
