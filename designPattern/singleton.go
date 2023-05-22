package designPattern

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("singleTom ")
}

var singleTon *SingleTon
var once sync.Once

const parCount = 100

// 定义一个类：singeton
type SingleTon struct {
}

func (*SingleTon) OnlySay() {
	fmt.Println("singleTom OnlySayn!!")
}

// 单例模式
func GetInstance() *SingleTon {
	once.Do(func() {
		singleTon = &SingleTon{}
	})
	return singleTon
}

func SingletonT() {
	// I := GetInstance()
	// fmt.Println(I)
	// I.OnlySay()

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
