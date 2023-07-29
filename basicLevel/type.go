package basicLevel

import (
	"fmt"
)

func TypeTest3() {
	var var1 int = 7
	var2 := float32(var1)
	var3 := int64(var1)
	//var4 := []int8(var1)//上面的var4就会报错，因为类型不兼容。
	fmt.Println(var2, var3)
}

func TypeTest2() {
	var x interface{} = 7 //x的动态类型为int,值为7
	i := x.(int)          // i的类型为int ,值为7

	fmt.Println(i)

	//type I interface{ m() }
	//var y I
	//
	//s := y.(string)    //非法: string 没有实现接口 I (missing method m)
	//r := y.(io.Reader) //y如果实现了接口io.Reader和I的情况下，  r的类型则为io.Reader

}

func TypeTest1() {
	//var i int = 10
	//错误 i不是接口类型
	//v := i.(int)

}

func TypeTest() {
	type Agg struct {
		ID   int
		Name string
	}

	type Agg1 struct {
		ID   int
		Name string
	}
	var a1 = Agg{1, "foo"}
	a2 := (Agg1)(a1)
	fmt.Println(a1, a2)
}
