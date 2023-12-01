package oneDayOneLibray

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func Time() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2018-04-24 10:11:20
	fmt.Println(time.Now().Format(time.UnixDate))
}

func Time_2() {
	fmt.Println(time.Now()) //2023-12-01 14:53:27.856524 +0800 CST m=+0.003546751

}
func Time_1() {
	layout := "2006-01-02 15:04:05"
	startTime, err := time.Parse(layout, cast.ToString("2023-11-11 15:04:05"))
	//startTime.
	fmt.Println(startTime, err)
	fmt.Println(startTime.Format("2006-01-02"))
	//start_time := time.Unix(ist[content_id]["start_time"].(int64), 0)
	//formatted_date := start_time.Format("2006-01-02")
}
