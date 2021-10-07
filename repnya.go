package Repnya

import "net/http"

type RoutServe interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	HEAD(pattern string, hf http.HandlerFunc)
	GET(pattern string, hf http.HandlerFunc)
	POST(pattern string, hf http.HandlerFunc)
	PUT(pattern string, hf http.HandlerFunc)
	DEL(pattern string, hf http.HandlerFunc)
	OPTIONS(pattern string, hf http.HandlerFunc)
	ServeStaticFiles(folderName string)
	GetKeyInt(r *http.Request, key string) (id int)
	GetKeyStr(r *http.Request, param string) string
	assign(method string, pattern string, hf http.HandlerFunc)
}
type RoutHandler struct {
	Pattern string
	http.HandlerFunc
}
type RoutServeMux struct {
	Handlers map[string][]*RoutHandler
}

func (rout *RoutServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (rout *RoutServeMux) HEAD(pattern string, hf http.HandlerFunc) {
	rout.assign("HEAD", pattern , hf)
}

func (rout RoutServeMux) GET(pattern string, hf http.HandlerFunc) {
	rout.assign("GET", pattern , hf)
}

func (rout *RoutServeMux) POST(pattern string, hf http.HandlerFunc) {
	rout.assign("POST", pattern , hf)
}

func (rout RoutServeMux) PUT(pattern string, hf http.HandlerFunc) {
	rout.assign("PUT", pattern , hf)
}

func (rout *RoutServeMux) DEL(pattern string, hf http.HandlerFunc) {
	rout.assign("DELETE", pattern , hf)
}

func (rout *RoutServeMux) OPTIONS(pattern string, hf http.HandlerFunc) {
	rout.assign("Options", pattern , hf)
}

func (rout *RoutServeMux) ServeStaticFiles(folderName string) {
	panic("implement me")
}

func (rout RoutServeMux) GetKeyInt(r *http.Request, key string) (id int) {
	panic("implement me")
}

func (rout *RoutServeMux) GetKeyStr(r *http.Request, param string) string {
	panic("implement me")
}

func (rout *RoutServeMux) assign(method, pattern string, hf http.HandlerFunc) {
	handlers := rout.Handlers[method]
	for _, handler := range handlers {
		if handler.Pattern == pattern {
			return
		}
	}
	handler := &RoutHandler{
		Pattern:     pattern,
		HandlerFunc: hf,
	}
	rout.Handlers[method]=append(handlers, handler)

}
