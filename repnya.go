package Repnya

import (
	"net/http"
	"net/url"
)

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
	getMapKey(path string) (url.Values, bool)
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
func (rh *RoutHandler) getMapKey(path string) (url.Values, bool) {
	mapValues := make(url.Values)
	var  j int
	for i:=j; i < len(path); i++ {
		switch {
		case j >= len(rh.Pattern):
			if rh.Pattern != "/" && len(rh.Pattern) > 0 && rh.Pattern[len(rh.Pattern)-1] == '/' {
				return mapValues, true
			}
			return nil, false
		case rh.Pattern[j] == ':':
			var name, val string
			var nextSymbol byte
			name, nextSymbol, j = match(rh.Pattern, isStrOrInt, j+1)
			val, _, i = match(path, matchPart(nextSymbol), i)
			escapedVal, err := url.QueryUnescape(val)
			if err != nil {
				return nil, false
			}
			mapValues.Add(":"+name, escapedVal)
		case path[i] == rh.Pattern[j]:
			j++
		default:
			return nil, false
		}
	}
	if j != len(rh.Pattern) {
		return nil, false
	}
	return mapValues, true
}

func matchPart(b byte) func(byte) bool {
	return func(c byte) bool {
		return c != b && c != '/'
	}
}

func match(s string, f func(byte) bool, i int) (slice string, next byte, j int) {
	for j=i; j < len(s) && f(s[j]); {
		j++
	}
	if j < len(s) {
		next = s[j]
	}
	return s[i:j], next, j
}

func isString(byte byte) bool {
	return 'a' <= byte && byte <= 'z' || 'A' <= byte && byte <= 'Z' || byte == '_'
}

func isInt(byte byte) bool {
	return '0' <= byte && byte <= '9'
}

func isStrOrInt(byte byte) bool {
	return isString(byte) || isInt(byte)
}