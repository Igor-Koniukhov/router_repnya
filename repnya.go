package Repnya

import "net/http"

type RoutServe interface {

}
type RoutHandler struct {
	Pattern string
	http.HandlerFunc
	Redirect bool
}
type RoutServeMux struct {
	handlers map[string][]*RoutHandler
}
