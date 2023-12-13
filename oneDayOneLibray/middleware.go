package oneDayOneLibray

import "fmt"

func finalHandler(request Request) {
	// 打印请求内容到控制台
	fmt.Println("请求处理: ", request)
}

func middleware1(next HandleFunc) HandleFunc {
	return func(request Request) {
		// 在请求到达最终处理函数之前执行的逻辑
		fmt.Println("中间件1 - 请求前处理: ", request)
		// 调用下一个处理函数
		next(request)
	}
}

func middleware2(next HandleFunc) HandleFunc {
	return func(request Request) {
		// 在请求到达最终处理函数之后执行的逻辑
		next(request)
		fmt.Println("中间件2 - 请求后处理: ", request)
	}
}

// 定义最终处理函数类型
type HandleFunc func(request Request)

// 定义请求类型
type Request struct {
	// 请求内容
	Content string
}

// 定义 Middleware 类型
type Middleware func(next HandleFunc) HandleFunc

// 剥洋葱函数
func OnionHandler(finalHandler HandleFunc, middlewares ...Middleware) HandleFunc {
	// 将最终处理函数赋值给 nextHandler
	nextHandler := finalHandler

	// 逆序遍历中间件切片
	for i := len(middlewares) - 1; i >= 0; i-- {
		// 获取当前中间件
		middleware := middlewares[i]
		// 将当前中间件和上一个中间件的处理函数嵌套， 返回新的处理函数给 nextHandler
		nextHandler = middleware(nextHandler)
	}

	// 返回最终处理函数
	return nextHandler
}

func MiddleWareTest() {
	// 定义最终处理函数
	finalHandler := func(request Request) {
		fmt.Println("最终处理函数: ", request)
	}

	// 定义中间件函数
	middleware1 := func(next HandleFunc) HandleFunc {
		return func(request Request) {
			fmt.Println("中间件1 - 进入中间件1")
			next(request)
			fmt.Println("中间件1 - 离开中间件1")
		}
	}

	middleware2 := func(next HandleFunc) HandleFunc {
		return func(request Request) {
			fmt.Println("中间件2 - 进入中间件2")
			next(request)
			fmt.Println("中间件2 - 离开中间件2")
		}
	}

	middleware3 := func(next HandleFunc) HandleFunc {
		return func(request Request) {
			fmt.Println("中间件3 - 进入中间件3")
			next(request)
			fmt.Println("中间件3 - 离开中间件3")
		}
	}

	// 使用剥洋葱函数组合中间件和最终处理函数
	handler := OnionHandler(finalHandler, middleware1, middleware2, middleware3)

	// 创建一个请求对象
	request := Request{
		Content: "这是一个请求",
	}

	// 调用最终的处理函数
	handler(request)
}
