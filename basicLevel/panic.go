package basicLevel

import "fmt"

func PanicTest() {

	//panic 函数的两个参数
	//1、内置函数
	//2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
	//3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
	//4、直到goroutine整个退出，并报告错误

	//recover函数
	//1、内置函数
	//2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
	//3、一般的调用建议
	//a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
	//b). 可以获取通过panic传递的error
	//PanicTest1()
	//PanicTest2()
	//PanicTest3()
	//PanicTest4()
	//使用延迟匿名函数或封装后的函数使用recever都是可以的 。

	fmt.Println("PanicTest")
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func PanicTest1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			//println(err)
			//println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()
	panic("panic error!")
}

func PanicTest2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}

func PanicTest3() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
	panic("test panic1")
}

func PanicTest4() {
	defer func() {
		fmt.Println(recover()) //有效
	}()
	//defer recover() //无效！
	//defer fmt.Println(recover()) //无效！
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！
		}()
	}()

	panic("test panic")
}
