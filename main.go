package main

import (
	"fmt"
	"time"
	// "learning_golang/array/array" //如何导入包
)

func main() {

	test_slice()
	// test_time() //练习使用time包
	// test_iota()
	fmt.Println("Hello World!")

}

func test_slice() {
	var s1 []int //定义空切片，对应的值为nil
	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}
	s2 := []int{} //初始化空切片，对应的值为空[]，而不是nil
	if s2 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}
	fmt.Println(s2)

	var s3 []int = make([]int, 0)
	fmt.Println(s3)

	//注意判断slice是否为空不能使用nil或[],只能使用len(s1)

	//初始化并赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	arr := [5]int{11, 22, 33, 44, 55}
	var s6 []int
	s6 = arr[1:4] //前包后不包
	fmt.Println(s6)

	// var slice00 []int = make([]int, 10)
	// var slice11 = make([]int, 10)
	// var slice22 = make([]int, 10, 10)
	// fmt.Printf("make全局slice0 ：%v\n", slice00)
	// fmt.Printf("make全局slice0 ：%v\n", slice11)
	// fmt.Printf("make全局slice0 ：%v\n", slice22)
	// fmt.Println(slice00)

	//切片的内存布局
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice_1 := data[1:6]
	fmt.Println(data)
	fmt.Println(slice_1)
	slice_1[1] = 100 //实际修改的为内存中数据
	fmt.Println(data)
	p := &data[1]
	fmt.Println(p)

	var d = [5]struct { //定义结构体数组
		x int
	}{}
	s := d[:] //结构体数据赋值给切片
	s[1].x = 100
	fmt.Println(d)
	fmt.Println(s)

	//append使用
	var a = []int{1, 2, 3}
	var b = []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println(c)
	cc := append(c, 7)
	fmt.Println(cc)

	ccc := append(cc, 8, 9, 10)
	fmt.Println(ccc)

	//append ：向 slice 尾部添加数据，返回新的 slice 对象。
	slice_2 := []int{1, 2, 3}
	slice_3 := append(slice_2, 4)
	// fmt.Println(&slice_2, &slice_3)
	fmt.Printf("%p\n", &slice_2)
	fmt.Printf("%p\n", &slice_3) //返回不一样的指针地址

	arr_1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(arr_1)
	slice_4 := data[:2:3]
	fmt.Println(slice_4, len(slice_4), cap(slice_4))
	slice_5_2 := []int{0, 1, 2, 3, 4}
	fmt.Println(slice_5_2)
	// slice中cap重新分配规律：2倍分配

}

func test_iota() {
	//巧妙使用iota,让数组索引自增
	const (
		name int = iota
		email
		age
	)

	var arr4 = [...]string{name: "lucy", email: "lucy@gmail.com", age: "10"}
	fmt.Println(name, email, age)
	fmt.Println(arr4)
}

func test_time() {
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
}
