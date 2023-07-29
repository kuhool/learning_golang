package oneDayOneLibray

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"math/rand"
)

func GoEchartsTest() {
	//实验失败
	generateBarItems()
}

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(500)})
	}
	return items
}
