package interview

import "fmt"

//1.通过指针变量 p 访问其成员变量 name，有哪几种方式？

type Person struct {
	Name string
}

func QuestionTwo2() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}
func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}

func QuestionTwo() {
	p := &Person{"张三"}

	fmt.Println(p.Name)    //指针变量对应的name的值
	fmt.Println(&p.Name)   //取name的地址
	fmt.Println((*p).Name) //获取p指针对应值的name

}
