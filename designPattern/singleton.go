package designPattern

import (
	"fmt"
	"sync"
)

// init 单独包中初始化变量
func init() {
	fmt.Println("singleTom ")
}

// 定义一个指针变量，存放实例化后的对象地址
var singleTon *SingleTon

var once sync.Once

const parCount = 100

// 定义一个类:SingleTon
type SingleTon struct {
}

// 定义类SingleTon的方法
func (*SingleTon) OnlySay() {
	fmt.Println("singleTom OnlySayn!!!")
}

// 实例化类SingleTon
// 使用once.Do保证只实例化一次类SingleTon
func GetInstance() *SingleTon {
	once.Do(func() { //必须要使用once.Do保证只实例化一次
		singleTon = &SingleTon{}
	})
	return singleTon
}

func SingletonT() {
	//I := GetInstance()
	//fmt.Println(I)
	//I.OnlySay()

	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*SingleTon{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(instances)
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			fmt.Println("instance is not equal!!")
		}
	}

	fmt.Println("GetInstance finished")
}
