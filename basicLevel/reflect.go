package basicLevel

import (
	"fmt"
	"reflect"
)

type Agg struct {
	Name string
}
type CSpuAgg struct {
	Id    int
	Title string
	Arr   []string
	Name1 Agg
}

func ReflectTest() {
	src := CSpuAgg{
		Id:    1,
		Title: "343r",
		Arr:   []string{"222", "333", "444"},
		Name1: Agg{Name: "111"},
	}

	dest := CSpuAgg{
		Id:    1,
		Title: "343r",
		Arr:   []string{"222", "333", "444"},
		Name1: Agg{Name: "11"},
	}
	Type := reflect.TypeOf([]int{10, 20, 32})
	//src := reflect.ValueOf([]int{10, 20, 32})
	//dest := reflect.ValueOf([]int{1, 2, 3})
	//cnt := reflect.Copy(dest, src)
	//cnt += 1
	fmt.Println("===", src, dest, Type)
	// DeepEqual is used to check
	// two interfaces are eual or not
	res1 := reflect.DeepEqual(src, dest)
	fmt.Println("Is dest is equal to src:", res1)
}
