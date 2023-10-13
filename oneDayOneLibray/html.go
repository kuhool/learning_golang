package oneDayOneLibray

import (
	"fmt"
	"html"
	"strings"
)

func TestHtml() {
	s := "编辑卡项信息;并审核通过;"
	fmt.Println(s)

	s = strings.TrimRight(s, ";")
	fmt.Println(s)
}
func TestHtml1() {
	// 转义HTML标签
	original := `<h1>Hello,  World! </h1>`
	escaped := html.EscapeString(original)
	fmt.Println("Escaped HTML: ", escaped)

	// 反转义HTML标签
	unescaped := html.UnescapeString(escaped)
	fmt.Println("Unescaped HTML: ", unescaped)

	unescaped = html.UnescapeString(unescaped)
	fmt.Println("Unescaped HTML: ", unescaped)
}
