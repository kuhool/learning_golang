package basicLevel

import "fmt"

//参考文案：https://www.cnblogs.com/zsy/p/5370052.html

// iota规则：
// 1. iota只能在常量的表达式中使用。
// 2. const中首次出现iota初始化为0
const a = iota // a=0
const (
	b = iota //b=0
	c        //c'=1
)

// 3. 自增长常量经常包含一个自定义枚举类型，允许你依靠编译器完成自增设置。

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 4. 可跳过的值

type AudioOutput int

const (
	OutMute   AudioOutput = iota // 0
	OutMono                      // 1
	OutStereo                    // 2
	_
	_
	OutSurround // 5
)

// 4. 位掩码表达式

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

//5 定义数量级

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

// 7、定义在一行的情况
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

// Apple: 1
// Banana: 2
// Cherimoya: 2
// Durian: 3
// Elderberry: 3
// Fig: 4

//8、中间插队

const (
	i = iota
	j = 3.14
	k = iota
	l
)

func IotaTest() {
	fmt.Println(a, b, c)
}
