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

}
type RoutHandler struct {
	Pattern string
	http.HandlerFunc
	Redirect bool
}
type RoutServeMux struct {
	Handlers map[string][]*RoutHandler
}

func (rout RoutServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (rout RoutServeMux) HEAD(pattern string, hf http.HandlerFunc) {
	panic("implement me")
}

func (rout RoutServeMux) GET(pattern string, hf http.HandlerFunc) {
	panic("implement me")
}

func (rout RoutServeMux) POST(pattern string, hf http.HandlerFunc) {
	panic("implement me")
}

func (rout RoutServeMux) PUT(pattern string, hf http.HandlerFunc) {
	panic("implement me")
}

func (rout RoutServeMux) DEL(pattern string, hf http.HandlerFunc) {
	panic("implement me")
}

func (rout RoutServeMux) OPTIONS(pattern string, hf http.HandlerFunc) {
	panic("implement me")
}

func (rout RoutServeMux) ServeStaticFiles(folderName string) {
	panic("implement me")
}

func (rout RoutServeMux) GetKeyInt(r *http.Request, key string) (id int) {
	panic("implement me")
}

func (rout RoutServeMux) GetKeyStr(r *http.Request, param string) string {
	panic("implement me")
}


