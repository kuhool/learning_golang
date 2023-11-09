package goFramework

import (
	"fmt"
	"net/http"
	"testing"
)

type Controller struct{}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !!!")
}

func TestServer(t *testing.T) {
	ctrl := &Controller{}
	// http.HandleFunc("/", ctrl.Index)
	http.HandleFunc("/hello", Logger(ctrl.Index))
	http.Handle("/beego/", http.StripPrefix("/beego/", http.FileServer(http.Dir("beego"))))

	http.ListenAndServe(":8080", nil)
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request: ", r.Method, r.URL.Path)
		next(w, r)
	}
}
