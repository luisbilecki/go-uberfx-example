package routes

import (
	"net/http"

	"github.com/luisbilecki/go-uberfx-example/handler"
)

func NewEchoRoute(echo *handler.EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}
