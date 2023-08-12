package oneDayOneLibray

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
)

type Car struct {
	Color string
}

func PieTest() {
	//slice1 := []int{1, 3, 4, 5}
	slice2 := []int{1, 5, 2}
	slice := pie.Sort(slice2)
	fmt.Println(slice)
}
func PieTest3() {
	//chars := pie.AreSorted([]float64{1, 2, 3, 4, 5})
	//chars := pie.AreUnique([]float64{1, 2, 3, 4, 5})
	//chars := pie.Average([]float64{1, 2, 3, 4, 5})
	//chars := pie.Bottom([]float64{1, 2, 3, 4, 5}, 3)
	//chars := pie.Chunk([]float64{1, 2, 3, 4, 5}, 3)

	// 输出结果

	// 定义切片
	//cars := []Car{
	//	{Color: "red"},
	//	{Color: "green"},
	//	{Color: "blue"},
	//}
	//
	//// 遍历切片并输出每个元素的颜色
	//pie.Each(cars, func(car Car) {
	//	fmt.Printf("Car color is: %s\n", car.Color)
	//})

	//chars := pie.GroupBy(
	//	[]int{23, 76, 37, 11, 23, 47},
	//	func(num int) int {
	//		return num % 5
	//	},
	//)
	// 输出结果
	//fmt.Println(chars)

}

func PieAllTest() {
	age := []int{1, 2, 99, 2}
	abs := pie.All(age, func(age int) bool {
		if age < 100 {
			return true
		} else {
			return false
		}
	})

	fmt.Println(abs)
}

func PieAbsTest1() {
	abs := pie.Abs(0.5)
	abs1 := pie.Abs(-0.5)
	fmt.Println(abs, abs1)
}

func PieUniqTest1() {
	arr := []int{1, 2, 3, 4, 5, 5, 6, 7}
	arrUni := pie.Unique(arr)
	fmt.Println(arrUni)
}
