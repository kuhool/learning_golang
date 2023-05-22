package designPattern

import "fmt"

func init() {
	fmt.Println("simplefactory ")
}

// 定义一个接口是API
type API interface {
	Say(name string) string
}

// 实现API
type HiAPI struct {
}

func (*HiAPI) Say(name string) string {
	return fmt.Sprintf("Hi %s ", name)
}

type HelloAPI struct {
}

func (*HelloAPI) Say(name string) string {
	return fmt.Sprintf("Hello %s ", name)
}

// 工厂模式New一个对象
// https://www.topgoer.cn/docs/golang-design-pattern/SimpleFactory
func NewAPI(t int) API {
	if t == 1 {
		return &HiAPI{}
	} else if t == 2 {
		return &HelloAPI{}
	} else {
		return nil
	}
}

func DesignPatternT() {

	//工厂生产的对象调用
	api := NewAPI(1)
	s := api.Say("Tom")
	fmt.Println(s)
}
