package oneDayOneLibray

import "fmt"

func Test() {

}

func Test1() {
	ch := '5'
	ch -= '0' //如果 ch 的值是字符 '5'，那么 '5' - '0' 的结果是整数 5，因为字符 '5' 的ASCII码是53，字符 '0' 的ASCII码是48，所以计算结果是53 - 48 = 5。
	fmt.Println(ch)
}
