package oneDayOneLibray

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

//https://github.com/qax-os/excelize

func Excel() {
	f, err := excelize.OpenFile("oneDayOneLibray/test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取单元格的值
	cellValue, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("A1: ", cellValue)
}
