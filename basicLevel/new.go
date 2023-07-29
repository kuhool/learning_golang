package basicLevel

import "fmt"

func NewTest() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	//if sm1 == sm2 { //TODO 编译错误map不能直接比较
	//	fmt.Println("sm1 == sm2")
	//}

}
func NewTest2() {

	//s1 := []int{1, 2, 3}
	//s2 := []int{4, 5}
	//s1 = append(s1, s2...) //elems ...Type ,append只能添加元素，想添加多个元素必须使用多值
	//fmt.Println(s1)

}

// NewTest1 错误情况演示：New不能初始化，slice,map ,channel
func NewTest1() {

	//list := new([]int)
}
