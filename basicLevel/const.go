package basicLevel

import "fmt"

func init() {
	fmt.Println("const 常量")
}
func ConstTest() {
	//申明单个常量
	const pi = 3.1415
	fmt.Println("pi:", pi)
	//申明多个常量
	const (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a, b, c)
	//iota的使用，就是累加1
	const (
		n1 = iota
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

	const (
		//_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println("1KB =", KB, "bytes")
	fmt.Println("1MB =", MB, "bytes")
	fmt.Println("1GB =", GB, "bytes")
	fmt.Println("1TB =", TB, "bytes")
	fmt.Println("1PB =", PB, "bytes")
}
