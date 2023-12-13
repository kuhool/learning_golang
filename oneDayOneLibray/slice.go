package oneDayOneLibray

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func TestSlice() {
	isOk := slices.Contains([]int{1, 2, 4}, 1)
	fmt.Println(isOk)
}
