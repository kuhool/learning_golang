package basicLevel

import (
	"fmt"
)

func AnonFuncTest() {

	//getSqrt := func(x float64) float64 {
	//	return math.Sqrt(x)
	//}
	//fmt.Println(getSqrt(4))
	//fn := func(str string) { fmt.Println(str) }
	//fn("hello")

	//定义了一个函数类型的切片
	//fns := []func(x int) int{
	//	func(x int) int { return x + 1 },
	//}
	//println(fns[0](100))

	// --- function as field ---
	//d := struct {
	//	fn func() string
	//}{
	//	fn: func() string { return "Hello struct func 定义一个结构体匿名函数" },
	//}
	//println(d.fn())

	// --- channel of function --- 匿名函数在管道中传输
	//fc := make(chan func() string, 2)
	//fc <- func() string { return "Hello, World!" }
	//println((<-fc)())

	//外部引用函数参数局部变量
	//tmp1 := add(10)
	//fmt.Println(tmp1(1), tmp1(2))
	//// 此时tmp1和tmp2不是一个实体了
	//tmp2 := add(100)
	//fmt.Println(tmp2(1), tmp2(2))

	//f1, f2 := test01(10)
	//// base一直是没有消
	//fmt.Println(f1(1), f2(2))
	//// 此时base是9
	//fmt.Println(f1(3), f2(4))

	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("id：%d,fibonaci：%d\n", i, fibonaci(i))
	}
}

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}
