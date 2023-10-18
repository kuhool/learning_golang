package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"html"
	"learning_golang/oneDayOneLibray"
	"learning_golang/todo/atomic"
	"learning_golang/todo/grt"
	"learning_golang/todo/log"
	"learning_golang/todo/selectT"
	"learning_golang/todo/timer"
	"os"
	reflect "reflect"
	"regexp"
	"runtime/trace"
	"strings"
	"time"
	"xorm.io/core"
)

var engine *xorm.Engine

func implode(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}
	b := make([]byte, n)
	slen := copy(b, a[0])
	for _, s := range a[1:] {
		slen += copy(b[slen:], sep)
		slen += copy(b[slen:], s)
	}
	return string(b)
}
func demo(paramDemo []int32) ([]int32, error) {
	fmt.Printf("main.demo.paramDemo pointer: %p \n", &paramDemo)
	paramDemo[0] = 2

	return paramDemo, nil
}

func contains(source []interface{}, obj interface{}) bool {
	for _, e := range source {
		if e == obj {
			return true
		}
	}
	return false
}
func StructToMapByTag(_input interface{}, _tag string) (rMap map[string]interface{}) {
	rMap = make(map[string]interface{})
	//反射结构体
	mapTyp := reflect.TypeOf(_input).Elem()
	mapVal := reflect.ValueOf(_input).Elem()
	fmt.Println(mapTyp, mapVal)
	num := mapVal.NumField() //获取字段数量
	for i := 0; i < num; i++ {
		//如果是个指针类型，则不能为nil
		if mapVal.Field(i).Kind() == reflect.Ptr && mapVal.Field(i).IsNil() {
			continue
		}
		//获取标签,需要通过reflect.type来获取标签
		fieldName := mapTyp.Field(i).Tag.Get(_tag)
		if fieldName == "" {
			fieldName = mapTyp.Field(i).Name
		}
		rMap[fieldName] = mapVal.Field(i)
	}
	return
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (y *User) aa() bool {

	return true
}

func separateContent(richText string) map[string]interface{} {
	imgPattern := `<\s*img\s+[^>]*?src\s*=\s*(\'|\")(.*?)\\1[^>]*?\/?\s*>`
	r := regexp.MustCompile(imgPattern)
	imgMatches := r.FindAllStringSubmatch(richText, -1)
	fmt.Println(imgMatches)
	img := make([]string, 0)
	for _, match := range imgMatches {
		img = append(img, match[2])
	}
	text := r.ReplaceAllString(richText, "")
	text = stripTags(text)
	return map[string]interface{}{
		"rich_text_img": img,
		"rich_text":     text,
	}
}
func stripTags(text string) string {
	re := regexp.MustCompile("<[^>]+>")
	return re.ReplaceAllString(text, "")
}

//	func separateContent(richText string) map[string]interface{} {
//		pattern := `<\s*img\s+[^>]*?src\s*=\s*('|")(.*?)\1[^>]*?\/?\s*>`
//		re := regexp.MustCompile(pattern)
//		imgMatches := re.FindAllStringSubmatch(richText, -1)
//		var img []string
//		for _, match := range imgMatches {
//			img = append(img, match[2])
//		}
//		text := re.ReplaceAllString(richText, "")
//		text = stripTags(text)
//		return map[string]interface{}{
//			"rich_text_img": img,
//			"rich_text":     text,
//		}
//	}
//
//	func stripTags(s string) string {
//		return strings.TrimSpace(regexp.MustCompile(`(?s)<[^>]*?>`).ReplaceAllString(s, ""))
//	}
func extractImageUrls(richText string) []string {
	imgPattern := `<img[^>]+src="([^"]+)"[^>]*>`
	r := regexp.MustCompile(imgPattern)
	matches := r.FindAllStringSubmatch(richText, -1)
	imgUrls := make([]string, 0)
	for _, match := range matches {
		imgUrls = append(imgUrls, match[1])
	}
	return imgUrls
}
func separateContent1(richText string) map[string]interface{} {
	pattern := regexp.MustCompile(`<img.*?src=['"](.*?)['"].*?>`)
	matches := pattern.FindAllStringSubmatch(richText, -1)

	var img []string
	for _, match := range matches {
		img = append(img, match[1])
	}

	text := pattern.ReplaceAllString(richText, "")
	text = strings.TrimSpace(strings.ReplaceAll(text, "<", ""))
	text = strings.TrimSpace(strings.ReplaceAll(text, ">", ""))

	return map[string]interface{}{
		"rich_text_img": img,
		"rich_text":     text,
	}
}

func SeparateContent(richText string) (string, []string) {
	pattern := regexp.MustCompile(`<img.*?src=['"](.*?)['"].*?>`)
	matches := pattern.FindAllStringSubmatch(richText, -1)

	var img []string
	for _, match := range matches {
		img = append(img, match[1])
	}

	stripped := html.UnescapeString(richText) // 解码HTML实体
	richText = html.UnescapeString(stripped)  // 额外解码一次，以处理双重转义
	text := pattern.ReplaceAllString(richText, "")
	regex := regexp.MustCompile("<[^>]*>")
	text = strings.TrimSpace(regex.ReplaceAllString(text, ""))
	regex1 := regexp.MustCompile(`[\s\n]+`)
	text = regex1.ReplaceAllString(text, "")

	return text, img
}

type Sku struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
}

func CheckString(input string) bool {
	regex := regexp.MustCompile(`^[\d, ]+$`)
	return regex.MatchString(input)
}

func main() {

	//oneDayOneLibray.TestOs()
	//oneDayOneLibray.Excel()
	//oneDayOneLibray.ReadInit()
	oneDayOneLibray.TestHtml()
	//basicLevel.SliceTest4()
	//oneDayOneLibray.PieTest()
	//oneDayOneLibray.CmpTest()
	//crossingLevels.ProtocolBufferTest()
	//intSlice := []int{1, 2, 3, 4, 5, 6}
	//
	//strSlice := make([]string, len(intSlice))
	//for i, num := range intSlice {
	//	strSlice[i] = strconv.Itoa(num)
	//}
	//str := "[" + strconv.Quote(strings.Join(strSlice, ",  ")) + "]"
	//
	//fmt.Println(str)

	// 获取商品信息
	//sku := &Sku{
	//	Id:       1,
	//	Name:     "sku1",
	//	Price:    100,)
	//str1 := "123,456,134343,545464"
	//str2 := "abc,def"
	//result1 := CheckString(str1)
	//result2 := CheckString(str2)
	//fmt.Println(str1, ":", result1)
	//fmt.Println(str2, ":", result2)

	//yfkCardName := "小绿卡"
	//yfkType := 1
	//yfkCardName = "%" + yfkCardName + "%"
	//whereStr := fmt.Sprintf("status ! = -1 AND type = %d AND name LIKE %s ", yfkType, yfkCardName)
	//fmt.Println(whereStr)
	//
	//query := strings.Join([]string{whereStr}, " ")
	//fmt.Println(query)
	//oneDayOneLibray.PieTest()
	//oneDay
	//interview.QuestionTwo()
	//basicLevel.IfTest()
	//oneDayOneLibray.NanoidTest()
	//oneDayOneLibray.PieAbsTest1()
	//oneDayOneLibray.PieUniqTest()

	//oneDayOneLibray.XormGetTest()

	//oneDayOneLibray.ContextTimeOutTest()
	//oneDayOneLibray.CtxTest()
	//oneDayOneLibray.ContextTest()
	//goFramework.GinServer()
	//$a = 1;
	//func b(){
	//	echo $a.PHP_EOL;
	//}

	//richText := `<p>测试代码<img src='https://img2.soyoung.com/product/20211213/4/8e36ebd8b3bb876ca1f90900701efe4a.jpg' watermark="0" width="750" height="1268"><img src="https://img2.soyoung.com/product/20211213/7/3d367c885817de5d8590a96c3c4e33e0.jpg" watermark="0" width="750" height="904"><img src="https://img2.soyoung.com/product/20211213/8/14e4235eb528c576190a6d6267215ac8.jpg" watermark="0" width="750" height="901"><img src="https://img2.soyoung.com/product/20211213/3/aa04ab1a0b7416b6b0ff927d89fe4fb8.jpg" watermark="0" width="750" height="911"><br></p><p><br></p>`
	//richText := `<p> <span style="font-size:14px;">我院李志海博士研发的智能韩式微创长曲线下颌角截骨术，是在口腔内采取微创切口的方式，利用智能颌面全景扫描机器人观察骨轮廓，运用德国颌面整形系统完成一次性截骨术，安全高效，弧形，塑造自然惊艳小脸。</span> </p> <p> <span style="font-size:14px;"><img src="https://img1.soyoung.com/post/20140629/0/20140629150331621.jpg" width="730" height="465" alt="" /><br /> </span> </p> <p> <span style="font-size:14px;"><img src="https://img1.soyoung.com/post/20140629/1/20140629150342651.jpg" width="730" height="347" alt="" /><br /> </span> </p> <p> <span style="font-size:14px;"><img src="https://img2.soyoung.com/post/20140629/1/20140629150351144.jpg" width="730" height="303" alt="" /><br /> </span> </p> <p> <span style="font-size:14px;"><img src="https://img2.soyoung.com/post/20140629/8/20140629150401735.jpg" width="730" height="311" alt="" /><br /> </span> </p> <p> <span style="font-size:14px;"><img src="https://img2.soyoung.com/post/20140629/2/20140629150410605.jpg" width="730" height="378" alt="" /><br /> </span> </p>	`
	//result := separateContent1(richText)
	////result := extractImageUrls(richText)
	//fmt.Println(result)

	//date := time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println(date)
	// 将返回现在的datetime 如 2019-04-27 22:37:23
	//basicLevel.ReflectGetTag()
	//u := User{1, "zs", 20}
	//basicLevel.ReflectShowStructField(u)
	//var x float64 = 3.4
	//basicLevel.ReflectSetValue(&x)
	//basicLevel.ReflectValue2(x)
	//basicLevel.ReflectKind2()
	//var x float64 = 3.4
	//basicLevel.ReflectType2(x)
	//var s []interface{}
	//s = append(s, "a")
	//s = append(s, "b")
	//s = append(s, 1)
	//fmt.Println(contains(s, 1))
	//os.Exit(0)
	//str := "杨先森test"
	//fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
	//basicLevel.TestReflect()
	//crossingLevels.TestCopier()
	//designPattern.T()
	//paramDemo := []int32{1, 2, 3}
	//fmt.Printf("a:%d ,ptr:%p &ptr:%p\n", paramDemo, paramDemo, &paramDemo)
	//fmt.Printf("main.paramDemo 1 %v, pointer: %p \n", paramDemo, &paramDemo)
	//// 浅拷贝
	//demo(paramDemo)
	//fmt.Printf("main.paramDemo 2 %v, pointer: %p \n", paramDemo, &paramDemo)

	//--------------------------------设计模式-----------------------
	//designPattern.FactorMethodT()
	//designPattern.SingletonT()
	//designPattern.SimpleFactoryT()

	//--------------------------------设计模式-----------------------
	// basicLevel.GoroutineT()
	// basicLevel.PanicTest()
	//basicLevel.AnonFuncTest()
	//basicLevel.FuncTest()
	//basicLevel.VariableTypeTest()
	//basicLevel.ConstTest()
	//basicLevel.VariableTest()
	//test_log()
	// test_trace()
	// test_atomic()
	// test_select()
	// test_timer()
	// test_grt()
	// test_xrom()
	// stru.Test()
	// sli.Test()
	// pointer.Test()
	// arr.Test()
	// test_strings()
	// test_slice_string()
	// test_slice()
	// test_time() //练习使用time包
	// test_iota()
	fmt.Println("\r\n-------finished!!!------- ")

}

func test_log() {
	log.Test()
}
func test_trace() {
	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
}
func test_atomic() {
	atomic.Test()
}

func test_select() {
	selectT.Test()
}

func test_timer() {
	timer.Test()
}

func odd(block chan struct{}) {
	for i := 1; i <= 100; i++ {
		<-block
		if i%2 == 1 {
			fmt.Println("奇数：", i)
		}
	}
}

func even(block chan struct{}) {
	for i := 1; i <= 100; i++ {
		block <- struct{}{}
		if i%2 == 0 {
			fmt.Println("偶数：", i)
		}
	}
}
func test_grt() {
	grt.Test()
}
func test_xrom() {
	engine, err := xorm.NewEngine("mysql", "lvweb1:lavion#2013@tcp(polar-dev.rwlb.rds.aliyuncs.com:3306)/db_saas?charset=utf8")
	fmt.Println(engine, err)
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	fmt.Println(engine.Ping())

}

func test_strings() {
	s := "hello world hello world"
	str := "wl"

	//返回字符串str中的任何一个字符在字符串s中第一次出现的位置。
	//如果找不到或str为空则返回-1
	index := strings.IndexAny(s, str)
	fmt.Println(index)

	a := "qwert"
	fmt.Println(strings.Split(a, ""))
}

func test_slice_string() {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:] //6为索引位子
	fmt.Println(s2)

	str_1 := "hello world"
	s := []rune(str_1) //字符串转换为byte类型,如果是中文使用[]rune(str_1)
	fmt.Println(str_1, s)
	s[6] = 'G'
	s_1 := s[:8]
	fmt.Println(string(s), string(s_1))

	// a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x [重点关注]

	// 数组or切片转字符串：
	array_or_slice := [...]int{1, 2, 3}
	str_9 := strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
	fmt.Println(fmt.Sprint(array_or_slice), str_9)

}

func test_slice() {
	var s1 []int //定义空切片，对应的值为nil
	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}
	s2 := []int{} //初始化空切片，对应的值为空[]，而不是nil
	if s2 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}
	fmt.Println(s2)

	var s3 []int = make([]int, 0)
	fmt.Println(s3)

	//注意判断slice是否为空不能使用nil或[],只能使用len(s1)

	//初始化并赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	arr := [5]int{11, 22, 33, 44, 55}
	var s6 []int
	s6 = arr[1:4] //前包后不包（前开后闭）
	fmt.Println(s6)

	// var slice00 []int = make([]int, 10)
	// var slice11 = make([]int, 10)
	// var slice22 = make([]int, 10, 10)
	// fmt.Printf("make全局slice0 ：%v\n", slice00)
	// fmt.Printf("make全局slice0 ：%v\n", slice11)
	// fmt.Printf("make全局slice0 ：%v\n", slice22)
	// fmt.Println(slice00)

	//切片的内存布局
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice_1 := data[1:6]
	fmt.Println(data)
	fmt.Println(slice_1)
	slice_1[1] = 100 //实际修改的为内存中数据
	fmt.Println(data)
	p := &data[1]
	fmt.Println(p)

	var d = [5]struct { //定义结构体数组
		x int
	}{}
	s := d[:] //结构体数据赋值给切片
	s[1].x = 100
	fmt.Println(d)
	fmt.Println(s)

	//append使用
	var a = []int{1, 2, 3}
	var b = []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println(c)
	cc := append(c, 7)
	fmt.Println(cc)

	ccc := append(cc, 8, 9, 10)
	fmt.Println(ccc)

	//append ：向 slice 尾部添加数据，返回新的 slice 对象。
	slice_2 := []int{1, 2, 3}
	slice_3 := append(slice_2, 4)
	// fmt.Println(&slice_2, &slice_3)
	fmt.Printf("%p\n", &slice_2)
	fmt.Printf("%p\n", &slice_3) //返回不一样的指针地址

	arr_1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(arr_1)
	slice_4 := data[:2:3]
	fmt.Println(slice_4, len(slice_4), cap(slice_4))
	slice_5_2 := []int{0, 1, 2, 3, 4}
	fmt.Println(slice_5_2)
	// slice中cap重新分配规律：2倍分配

	s_1 := []int{1, 2, 3, 4, 5}
	s_2 := make([]int, 10)
	s_2[0] = 1000
	fmt.Println(s_1, s_2)
	copy(s_2, s_1) //copy是两个slice直接数据拷贝，索引区域覆盖对应的值
	fmt.Println(s_1, s_2)

	s_3 := []int{1, 2, 3}
	s_3 = append(s_3, s_2...) //append 合并两个切片，返回新的 (合并后) 切片对象。
	fmt.Println(s_3)
	s_3 = append(s_3, 4, 5, 6) //append讲多个元素合并到切片中
	fmt.Println(s_3)

	s_5 := []int{5, 6, 7, 8, 9}
	fmt.Println(s_5)
	s_6 := s_5[1:2]
	fmt.Println(s_6)
	s_7 := s_6[1:4] //前包后不包（前开后闭）
	fmt.Println(s_7)

	//字符串切片，string底层就是一个byte的数组，因此，也可以进行切片操作。

}

func test_iota() {
	//巧妙使用iota,让数组索引自增
	const (
		name int = iota
		email
		age
	)

	var arr4 = [...]string{name: "lucy", email: "lucy@gmail.com", age: "10"}
	fmt.Println(name, email, age)
	fmt.Println(arr4)
}

func test_time() {
	var arr []int = []int{9: 10, 8: 10}
	fmt.Println(arr) // [9 10]

	t := time.Now() //Now()格式不能直接使用需要转换
	fmt.Println(t)
	t1 := t.Unix() //获取时间戳
	fmt.Println(t1)
	t = t.Add(time.Hour * 24) //时间相加一天
	fmt.Println(t)
	t2 := t.Format("2006-01-02 15:04:05") //格式化Y-m-d H:i:s
	fmt.Println(t2)

	//指定时间戳转换成时间
	timeStamp := 1680602716
	t3 := time.Unix(int64(timeStamp), 0)
	fmt.Println(t3)

	timeStamp = 1569826535             // 2019-09-30 14:55:35
	t = time.Unix(int64(timeStamp), 0) // time.Time
	year, month, day := t.Date()

	hour, minute, second := t.Clock()

	fmt.Println(year, month, day, hour, minute, second)

	// 同样可以根据时间格式提供的方法
	// 可以获取星期几，还可以获取当前是一年的多少天
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Weekday(), t.YearDay())

	fmt.Println(time.Now().Format("1"))

	// 判断某个时间是不是在某个时间之前/之后/相等，传入的参数必须是时间格式
	t4 := t.After(time.Now())
	t5 := t.Before(time.Now())
	t6 := t.Equal(t)
	// 返回的是bool值
	fmt.Println(t4, t5, t6) // false true true
}
