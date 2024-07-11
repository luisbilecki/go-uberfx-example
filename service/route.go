package service

import (
	"net/http"

	"github.com/luisbilecki/go-uberfx-example/handler"
)

func NewServeMux(route1, route2 handler.Route) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(route1.Pattern(), route1)
	return mux
}
