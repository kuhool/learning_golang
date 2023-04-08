package pointer

import (
	"fmt"
)

func Test() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%s\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
	// 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储
	//而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间
	// make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作
	// var b map[string]int
	// b["测试"] = 100
	// fmt.Println(b)
	fmt.Println("pointer.test")
}
