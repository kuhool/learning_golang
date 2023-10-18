package oneDayOneLibray

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"html"
	"log"
	"net/http"
	"strings"
)

func TestHtml() {
	// 发送 HTTP 请求并获取响应
	response, err := http.Get("https://www.sy.soyoung.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// 使用 goquery 解析 HTML 文档
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(doc.Text())
	// 查询并处理 HTML 元素
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// 获取链接文本和链接地址
		linkText := s.Text()
		linkHref, _ := s.Attr("href")

		fmt.Printf("Link %d:  %s - %s\n", i, linkText, linkHref)
	})
}
func TestHtml2() {
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

}
