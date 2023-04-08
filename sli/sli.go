package sli

import "fmt"

func Test() {

	//slice 是有3个字段的数据结构，指向底层数组的指针，切片的长度和切片的容量
	arr := []int{1, 2, 3, 4, 5}

	slice := arr[0:4]
	fmt.Println("Original array:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Length:", len(arr), "Capa:", cap(arr))
	fmt.Println("Length:", len(slice), "Capa:", cap(slice))
	fmt.Printf("slice 的内存地址：%p\n", slice)   //指针指向切片的第一个访问元素
	fmt.Printf("arr[0]的内存地址：%p\n", &arr[0]) //获取底层数组中的首个元素地址

	//1.使用make新建一个切片
	s1 := make([]int, 2, 4)
	fmt.Printf("s1 的长度：%d，容量：%d\n", len(s1), cap(s1))
	s2 := make([]int, 2)
	fmt.Printf("s2 的长度：%d，容量：%d\n", len(s2), cap(s2))
	//常规创建切片
	s3 := []int{0, 1, 2, 3, 4}
	fmt.Printf("s3 的长度：%d，容量：%d\n", len(s3), cap(s3))
	s4 := []int{9: 9} //[0 0 0 0 0 0 0 0 0 9]
	fmt.Printf("s4 的长度：%d，容量：%d\n", len(s4), cap(s4))
	fmt.Println("Slice_s4:", s4)

	//创建nil切片值为nil
	var sliceNil []int
	fmt.Printf("sliceNil 的长度：%d，容量：%d，类型：%T\n", len(sliceNil), cap(sliceNil), sliceNil)
	fmt.Println("sliceNil = ", sliceNil)
	if sliceNil == nil {
		fmt.Println("sliceNil is nil ") //打印这个
	} else {
		fmt.Println("sliceNil is not  nil ")
	}

	//创建空切片值不为nil
	s5 := make([]int, 0)
	fmt.Printf("s5 的长度：%d，容量：%d\n", len(s5), cap(s5))

	s6 := []int{}
	fmt.Printf("s6 的长度：%d，容量：%d\n", len(s6), cap(s6))
	fmt.Println("s5,s6 = ", s5, s6)
	if s6 == nil {
		fmt.Println("s6 is nil ")
	} else {
		fmt.Println("s6 is not  nil ") //打印这个
	}
	//基于切片创建新切片。
	s7 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	s8 := s7[1:3]
	fmt.Printf("s7 的长度：%d，容量：%d\n", len(s7), cap(s7))
	fmt.Printf("s8 的长度：%d，容量：%d\n", len(s8), cap(s8))
	s8[0] = 21
	fmt.Println("s8 = ", s8)
	fmt.Println("s7 = ", s7)

	//多维切片
	s9 := [][]int{{1, 2, 3}, {11, 12, 13, 14}}
	fmt.Printf("s9 的长度：%d，容量：%d\n", len(s9), cap(s9))
	fmt.Printf("s9[0] 的长度：%d，容量：%d\n", len(s9[0]), cap(s9[0]))
	fmt.Printf("s9[1] 的长度：%d，容量：%d\n", len(s9[1]), cap(s9[1]))
	fmt.Println(s9)
	for index, value := range s9[0] {
		fmt.Printf("s9[0][%d] = %d\n", index, value)
	}

	fmt.Println("sli.Test")
}
