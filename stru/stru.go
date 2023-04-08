package stru

import (
	"encoding/json"
	"fmt"
)

//自定义类型
type NewInt int  //定义类型
type MyInt = int //定义类型别名
var user struct {
	Name string
	Age  int
}

type person struct {
	name string
	city string
	age  int8
}

//遍历所有同学信息
type student struct {
	name   string
	age    int8
	gender string
}

type student_1 struct {
	id   int
	name string
	age  int
}

type Student_2 struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

//Student 学生
type Student_3 struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student_3
}

func Test() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a) //%T就是打印类型
	fmt.Printf("type of b:%T\n", b)

	fmt.Println(a, b)

	// user.Name = "pprof.cn"
	// user.Age = 18
	fmt.Printf("%#v\n", user) //user为一个实例
	var p2 = new(person)
	fmt.Printf("%#v\n", p2) //p2为一个指针地址（和user做对比）
	p2.name = "Jonah"       //给结构体属性赋值
	p2.age = 20
	fmt.Printf("%#v\n", p2)

	//使用结构体键值对初始化
	p5 := person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5)
	stu := []student{
		{name: "pprof.cn", age: 18, gender: "男"},
		{name: "测试", age: 23, gender: "男"},
		{name: "博客", age: 28, gender: "男"},
	}

	fmt.Printf("p5=%#v\n", stu)

	m := make(map[string]*student)
	for _, ss := range stu {
		// fmt.Println(ss.name, &ss)
		// sss := ss
		m[ss.name] = &ss //错误示例，因为ss指向相同一块内存
	}

	fmt.Println(m)
	//打印的结果不是我们想要的
	//map[pprof.cn:0xc00010e1b0 博客:0xc00010e1b0 测试:0xc00010e1b0]

	// fmt.Printf("p5=%#v\n", m)
	p9 := newPerson("pprof.cn", "测试", 90)
	fmt.Printf("%#v\n", p9)

	var ce []student_1 //定义一个切片类型的结构体
	ce = []student_1{
		student_1{1, "xiaoming", 22},
		student_1{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)

	s1 := Student_2{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}

	fmt.Printf("json str:%s\n", data)

	c := &Class{
		Title:    "101",
		Students: make([]*Student_3, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student_3{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	fmt.Printf("json:%s\n", c)
	// //JSON序列化：结构体-->JSON格式的字符串
	data_1, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	fmt.Printf("json:%s\n", data_1)

	// //JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("c1:%#v\n", c1)
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<stru.test>>>>>>>>>>>>>>>>>>>>>")
}

func demo(ce []student_1) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}

// 构造函数，自己实现结构体的构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
