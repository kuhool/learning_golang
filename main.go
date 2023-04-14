package main

import (
	"fmt"
	"learning_golang/basicLevel"
	"learning_golang/todo/atomic"
	"learning_golang/todo/grt"
	"learning_golang/todo/log"
	"learning_golang/todo/selectT"
	"learning_golang/todo/timer"
	"os"
	"runtime/trace"
	"strings"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var engine *xorm.Engine

func main() {
	basicLevel.Test()
	//test_log()
	// test_trace()
	// test_atomic()
	// test_select()
	// test_timer()
	// test_grt()
	// test_xrom()
	// stru.Test()
	// sli.Test()
	// pointer.Test()
	// arr.Test()
	// test_strings()
	// test_slice_string()
	// test_slice()
	// test_time() //练习使用time包
	// test_iota()
	fmt.Println("-------finished!!!------- ")

}

func test_log() {
	log.Test()
}
func test_trace() {
	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
}
func test_atomic() {
	atomic.Test()
}

func test_select() {
	selectT.Test()
}

func test_timer() {
	timer.Test()
}

func odd(block chan struct{}) {
	for i := 1; i <= 100; i++ {
		<-block
		if i%2 == 1 {
			fmt.Println("奇数：", i)
		}
	}
}

func even(block chan struct{}) {
	for i := 1; i <= 100; i++ {
		block <- struct{}{}
		if i%2 == 0 {
			fmt.Println("偶数：", i)
		}
	}
}
func test_grt() {
	grt.Test()
}
func test_xrom() {
	engine, err := xorm.NewEngine("mysql", "lvweb1:lavion#2013@tcp(polar-dev.rwlb.rds.aliyuncs.com:3306)/db_saas?charset=utf8")
	fmt.Println(engine, err)
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	fmt.Println(engine.Ping())

}

func test_strings() {
	s := "hello world hello world"
	str := "wl"

	//返回字符串str中的任何一个字符在字符串s中第一次出现的位置。
	//如果找不到或str为空则返回-1
	index := strings.IndexAny(s, str)
	fmt.Println(index)

	a := "qwert"
	fmt.Println(strings.Split(a, ""))
}

func test_slice_string() {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:] //6为索引位子
	fmt.Println(s2)

	str_1 := "hello world"
	s := []rune(str_1) //字符串转换为byte类型,如果是中文使用[]rune(str_1)
	fmt.Println(str_1, s)
	s[6] = 'G'
	s_1 := s[:8]
	fmt.Println(string(s), string(s_1))

	// a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x [重点关注]

	// 数组or切片转字符串：
	array_or_slice := [...]int{1, 2, 3}
	str_9 := strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
	fmt.Println(fmt.Sprint(array_or_slice), str_9)

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
	s6 = arr[1:4] //前包后不包（前开后闭）
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

	s_1 := []int{1, 2, 3, 4, 5}
	s_2 := make([]int, 10)
	s_2[0] = 1000
	fmt.Println(s_1, s_2)
	copy(s_2, s_1) //copy是两个slice直接数据拷贝，索引区域覆盖对应的值
	fmt.Println(s_1, s_2)

	s_3 := []int{1, 2, 3}
	s_3 = append(s_3, s_2...) //append 合并两个切片，返回新的 (合并后) 切片对象。
	fmt.Println(s_3)
	s_3 = append(s_3, 4, 5, 6) //append讲多个元素合并到切片中
	fmt.Println(s_3)

	s_5 := []int{5, 6, 7, 8, 9}
	fmt.Println(s_5)
	s_6 := s_5[1:2]
	fmt.Println(s_6)
	s_7 := s_6[1:4] //前包后不包（前开后闭）
	fmt.Println(s_7)

	//字符串切片，string底层就是一个byte的数组，因此，也可以进行切片操作。

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
