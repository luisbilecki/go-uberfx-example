package service

import (
	"net/http"

	"github.com/luisbilecki/go-uberfx-example/handler"
)

func NewServeMux(route handler.Route) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(route.Pattern(), route)
	return mux
}
