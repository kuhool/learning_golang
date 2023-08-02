package basicLevel

import "C"
import "fmt"

func ap(a []int) {
	fmt.Printf("%p\n", a)
	//b := append(a, 10)

	fmt.Printf("%p\n", a)
}

func app(a []int) {
	fmt.Printf("%p\n", a)
	a[0] = 1
	fmt.Printf("%p\n", a)
}

func SliceTest() {
	//var x = nil
	//var x interface{} = nil
	//var x string = nil
	var x error = nil
	fmt.Println(x)
	//str := 'abc' + '123'
	//str := "abc" + "123"
	//str := '123' + "abc"
	//str := fmt.Sprintf("abc%d", 123)

	//fmt.Println(str)
}
func SliceTest1() {
	a := []int{7, 8, 9}
	fmt.Printf("%p\n", a)
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}
