package goFramework

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinServer() {
	r := gin.Default()

	//请求地址：http://127.0.0.1:8080/ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	//请求地址：http://127.0.0.1:8080/hello
	//fdfdf1123456789

	r.GET("/hello", Hello)
	r.Run()

}

func Hello(c *gin.Context) {

	c.String(200, "hello %s", "world")
	c.JSON(http.StatusOK, gin.H{ //以json格式输出
		"name": "tom",
		"age":  "20",
	}) //200代表请求成功,http.StatusOK代表请求成功,gin.H是map的一个快捷方式
}
