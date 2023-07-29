package basicLevel

import "fmt"

//参考文章
// https://mp.weixin.qq.com/s/aC5BZWuO7bRJdc_xUmMTTw
//小结：
//1. if条件语句，隐私代码段，
//2. if条件语句的作用域
//3. 没有label的break
//4. range 循环的是变量的副本

func IfTest() {
	var sl = []int{5, 19, 6, 3, 8, 12, 36}
	var firstEven int = -1

	//
	for i := 0; i < len(sl); i++ {
		switch sl[i] % 2 {
		case 0:
			firstEven = sl[i]
			break //没有label的break 表示跳出当前循环  for switch是双层循环，没有完全调出
		case 1:
			// do nothing
		}
	}
	println(firstEven)
}

func IfTest3() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a)

	for i, v := range a { //rang a 实际是在循环a的副本
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after for range loop, r =", r)
	fmt.Println("after for range loop, a =", a)

	//预期值与实际输出只有啥用 ？rang a 实际是在循环a的副本
	//original a = [1 2 3 4 5]
	//after for range loop, r = [1 2 3 4 5]
	//after for range loop, a = [1 12 13 4 5]
}

// IfTest2  代码中隐私代码块
func IfTest2() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		println(a, b, c) //为什么能输出ABC的值？？就是隐私代码的作用
	}
}

// IfTest1 函数中上面和下面两个代码块差异就是隐私代码中作用域不一样行
func IfTest1() {

	//if a, ok := foo(); a < 10 && ok{ //使用if表达式自用变量
	//
	//}

	//TODO vs.

	//a, ok := foo()
	//if a < 10 && ok {
	//
	//}
}
