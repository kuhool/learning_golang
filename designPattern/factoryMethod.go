package designPattern

/**
设计模式之工厂方法模式
将生产的类封装到工厂类统一的方法中

A 需要被生产出来
B 也要被生产出来
AA 中的方法封装A的类并实例化
BB 中的方法封装B的类并实例化
新增一个C后
原来代码不需要调整，直接新增类CC封装C并实例化

*/
import (
	"fmt"
)

/*
*
简单工厂模式
*/
func init() {
	fmt.Println("工厂方法模式")
}

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

// SetA 设置 A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// SetB 设置 B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
// PlusOperatorFactory 实现了接口OperatorFactory
type PlusOperatorFactory struct{}

// 实现了接口的方法
func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase //指向OperatorBase结构体的指针
}

// Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func FactorMethodT() {
	var (
		factory OperatorFactory
	)

	factory = PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		fmt.Println("error with factory method pattern")
	}

	factory = MinusOperatorFactory{}
	if compute(factory, 4, 2) != 2 {
		fmt.Println("error with factory method pattern")
	}

}
