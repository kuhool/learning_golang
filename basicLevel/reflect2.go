package basicLevel

import (
	"fmt"
	"reflect"
)

// 参考 https://www.topgoer.cn/docs/golang//1082
type Student struct {
	Name string `json:"name1" db:"name2" haha:"name3"`
}

func ReflectGetTag() {
	var s Student
	v := reflect.ValueOf(&s)
	// 类型
	t := v.Type()
	fmt.Println(v.Type())

	t1 := reflect.TypeOf(v)
	fmt.Println(t1.Kind()) //struct
	//if t1.Kind() == reflect.Ptr {
	//	t1 = t1.Elem()
	//
	//}

	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
	fmt.Println(f.Tag.Get("haha"))
}

// ReflectShowStructField 结构体与反射
func ReflectShowStructField(s interface{}) {
	t := reflect.TypeOf(s)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())

	// 获取值
	v := reflect.ValueOf(s)

	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}

	fmt.Println("=================方法====================")

	fmt.Println(t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}

}
func ReflectSetValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 反射修改值,必须使用地址传递
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}

	fmt.Println(k)
}

func ReflectValue2(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
	}
}
func ReflectType2(a interface{}) {
	fmt.Println("basicLevel.ReflectType2")
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)

	k := t.Kind() //获取底层类型
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}

	fmt.Println("basicLevel.ReflectType2")
}

// ReflectKind2 获取变量的底层数据类型
func ReflectKind2() {
	type MyInt int
	var num MyInt = 1
	t := reflect.TypeOf(num)
	kind := t.Kind()
	fmt.Println(t, kind) // 输出 basicLevel.MyInt int
}
