package basicLevel

import (
	"fmt"
	"math"
)

//func init() {
//	fmt.Println("const 常量")
//}

func VariableTypeTest() {
	//整型
	//int8 int16 int32 int64
	//uint8 uint16 uint32 uint64

	//uint8就是我们熟知的byte型
	//rune 就是unicode字符，4个字节，占32位
	//une是int32 的别名

	//字符串的常用操作
	//len(str)	求长度
	//+或fmt.Sprintf	拼接字符串
	//strings.Split	分割
	//strings.Contains	判断是否包含
	//strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	//strings.Index(),strings.LastIndex()	子串出现的位置
	//strings.Join(a[]string, sep string)	join操作}

	//强制类型转换的基本语法如下：
	//runeS2 := []rune(s2)

	// math.Sqrt()接收的参数是float64类型，需要强制转换
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
