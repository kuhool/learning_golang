package basicLevel

import "fmt"

// variable 变量
//func init() {
//	fmt.Printf("-------variable 变量-------\n")
//}

func VariableTest() {
	//申明单个变量
	var hello string
	hello = "打印：hello world"
	fmt.Println(hello)

	//申明多个变量
	var (
		name   string
		age    int
		gender string
	)
	//每个变量会被初始化成其类型的默认值
	//比如：整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。
	//布尔型变量默认为false。 切片、函数、指针变量的默认为nil。
	fmt.Println(name, age, gender)

	//短变量
	n := 10
	//类型推到申明
	var country = "China"
	fmt.Println(n, country)

	//匿名变量 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。
	intVal, _ := foo() //忽略第二个参数
	fmt.Println(intVal)

	//注意点：
	//函数外的每个语句都必须以关键字开始（var、const、func等）
	//:=不能使用在函数外。
	//_多用于占位，表示忽略值。
}

func foo() (int, string) {
	return 10, "Q1mi"
}

func hello() []string {
	return nil

}

func VarTest1() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

}

func GetValue() int {
	return 1
}

func VarTest() {
	//i := GetValue()
	//switch i.(type) { //i (variable of type int) is not an interface
	//case int:
	//	println("int")
	//case string:
	//	println("string")
	//case interface{}:
	//	println("interface")
	//default:
	//	println("unknown")
	//}
}
