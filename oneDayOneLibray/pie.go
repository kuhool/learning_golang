package oneDayOneLibray

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
)

func PieUniqTest() {
	arr := []int{1, 2, 3, 4, 5, 5, 6, 7}
	arrUni := pie.Unique(arr)
	fmt.Println(arrUni)
}
