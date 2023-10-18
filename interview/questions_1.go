package interview

import "fmt"

//这个代码有什么缺陷？？
//答案:在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。这里的第一个返回值有命名 sum，第二个没有命名，所以错误。
//func funcMui(x, y int) (sum int, error) {
//	return x + y, nil
//}

func Questions() {

}
func Questions3() {
	//a := 5
	//b := 8.1
	//fmt.Println(a + b)//类型不一致不能相加
}

func Questions2() {
	//Total, err := funcMui(1, 2)
	//fmt.Println(Total, err)
}

func Questions1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s1 := make([]int, 0)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)
}
