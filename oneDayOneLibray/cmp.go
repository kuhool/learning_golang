package oneDayOneLibray

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

type Property struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Country []int
}

func CmpTest() {

	Property1 := Property{
		Country: []int{1, 2},
	}
	Property2 := Property{
		Country: []int{1, 2},
	}

	bool := cmp.Equal(Property1, Property2)
	fmt.Println(bool)
}
