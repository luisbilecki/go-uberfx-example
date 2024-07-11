package service

import (
	"net/http"

	"github.com/luisbilecki/go-uberfx-example/handler"
)

func NewServeMux(routes []handler.Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}
