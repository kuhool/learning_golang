package basicLevel

import "fmt"

func FuncTest() {
	//调用函数作为参数的函数
	num := test(func() int { return 100 })
	fmt.Println(num)

	num1 := test(intFunc)
	fmt.Println(num1)

	//注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，
	//不过，值传递是值的拷贝。引用传递是地址的拷贝，
	//一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
	//注意2：map、slice、chan、指针、interface默认以引用的方式传递。

	//使用 slice 对象做变参时，必须展开。（slice...）
	//s := []int{1, 2, 3}
	//res := test("sum: %d", s...) // slice... 展开slice
	//println(res)

}

// 申明一个函数
// 函数的参数也是一个函数，参数为函数和返回值
func test(fn func() int) int {
	return fn()
}

// TestArgs1 不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）
// 参数不固定，类型固定
// 注意：其中args是一个slice，我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数.
func TestArgs1(args ...int) {

}

// TestArgs2 任意类型的不定参数： 就是函数的参数和每个参数的类型都不是固定的。
func TestArgs2(args ...interface{}) {

}

// TestArgs3 注意：变量后面需要加...和类型
// 比如：n ...int
func TestArgs3(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}

	return fmt.Sprintf(s, x)
}
func intFunc() int {
	return 200
}
