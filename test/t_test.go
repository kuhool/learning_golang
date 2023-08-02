package test

import (
	"fmt"
	"testing"
)

func TestArr1(t *testing.T) {
	t.Run("开始测试第一个函数:", TestArray)
}

func TestArray(t *testing.T) {
	var arr1 [5]int
	arr2 := [3]int{3, 8, 33}
	arr3 := [...]int{3, 13, 41, 214, 3141, 324}
	fmt.Println(arr1, arr2, arr3)
}
