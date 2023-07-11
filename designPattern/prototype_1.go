package designPattern

import (
	"fmt"
	time "time"
)

// Cloneable 是一个原型对象需要实现的接口
//type Cloneable interface {
//	Clone() Cloneable
//}

// Person 是一个实现了 Cloneable 接口的结构体
type PPerson struct {
	Name string
	Age  int
}

// Clone 实现了克隆方法，返回一个新的 Person 对象
func (p *PPerson) Clone() *PPerson {
	tmp := *p
	return &tmp
}

func T() {
	// 创建一个原型对象
	original := &PPerson{
		Name: "John",
		Age:  30,
	}

	// 克隆原型对象来创建新对象
	clone := original.Clone()
	// 修改克隆对象的属性
	clone.Name = "Jane"
	clone.Age = 25

	// 打印原型对象和克隆对象的属性
	fmt.Println("Original:", original.Name, original.Age)
	fmt.Println("Clone:", clone.Name, clone.Age)
	fmt.Println("Clone:", original, clone)

	fmt.Println("time:", time.Time{})

}
