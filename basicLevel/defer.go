package basicLevel

import (
	"fmt"
	"sync"
	"time"
)

// defer f.Close
type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func Close(t Test) {
	t.Close()
}
func testD(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")
}
func testE() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制
	//打印结果：defer: 10 120

	x += 10
	y += 100
	println("x =", x, "y =", y)
}

var lock sync.Mutex

func testFF() {
	lock.Lock()
	lock.Unlock()
}

func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}
func ForTest() {
	t1 := time.Now()

	for i := 0; i < 10000; i++ {
		testFF()
	}
	elapsed := time.Since(t1)
	fmt.Println("test elapsed: ", elapsed)
}
func ForDeferTest() {
	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i++ {
			testdefer()
		}
		elapsed := time.Since(t1)
		fmt.Println("testdefer elapsed: ", elapsed)
	}()
}

func testF() {
	ForTest()
	ForDeferTest()
}

func Foo() (i int) {

	i = 0
	defer func() {
		fmt.Println(i)
	}()

	return 2
}
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func DeferTest() {
	defer_call()
}

// 滥用 defer 可能会导致性能问题，尤其是在一个 “大循环” 里。
func DeferTest8() {
	Foo()
}

// 滥用 defer 可能会导致性能问题，尤其是在一个 “大循环” 里。
func DeferTest7() {
	testF()
}

// DeferTest 延迟调用参数在注册时求值或复制，可用指针或闭包 “延迟” 读取。
func DeferTest6() {
	testE()
}

// DeferTest 多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。
// 规定：哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。
func DeferTest5() {
	testD(0)
}

func DeferTest4() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer Close(t) //函数做了值拷贝，输出的变量就是bca ,如果换成引用传值就是输出cc
	}
}

func DeferTest3() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close() //输出结果是ccc,原因每次调用defer的时候都保存相同t的地址
	}
}

// DeferTest defer 碰上闭包
func DeferTest2() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }() //输出结果为：4444
	}
}

// DeferTest  1. defer 是先进后出
func DeferTest1() {
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println(i) //先进后出原则，进入01234，后出43210
	}
}
