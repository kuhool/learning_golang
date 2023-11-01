package goFramework

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {

	http.Handle("/index", &CustomerHandler{})
	http.ListenAndServe(":8080", nil)
}

// CustomerHandler 实现了Handler中的ServeHTTP方法
type CustomerHandler struct {
}

func (c *CustomerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("implement http server by self")
	writer.Write([]byte("server is running"))
}
