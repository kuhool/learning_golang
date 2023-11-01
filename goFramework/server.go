package goFramework

import "net/http"

type Server interface {
	http.Server
	start()
}

type HttpServer struct {
}

func (h HttpServer) start() {
	//TODO implement me
	panic("implement me")
}
