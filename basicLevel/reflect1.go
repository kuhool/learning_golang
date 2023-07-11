package basicLevel

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func (p Person) SayHello() bool {
	fmt.Println("Hello!")
	return true
}

func (p Person) SayName() bool {
	fmt.Printf("My name is %s\n", p.name)
	return true
}

func TestReflect() {
	person := Person{name: "John", age: 30}

	// 获取 Person 结构体类型的 reflect.Type 对象
	personType := reflect.TypeOf(person)

	// 遍历结构体的方法集合，并输出方法名
	for i := 0; i < personType.NumMethod(); i++ {
		method := personType.Method(i)

		// 获取方法的 reflect.Value 对象，并调用方法
		methodValue := reflect.ValueOf(person).MethodByName(method.Name)

		result := methodValue.Call(nil)
		fmt.Println(methodValue, result)
		// 处理方法的返回值
		//for _, value := range result {
		//	ret := value.Interface()
		//	if ret != true {
		//
		//	}
		//
		//}
		//fmt.Println(len(method.Name), method.Name[:3])

	}
}
