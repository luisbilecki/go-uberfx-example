package main

import (
	"net/http"

	"github.com/luisbilecki/go-uberfx-example/handler"
	routes "github.com/luisbilecki/go-uberfx-example/route"
	"github.com/luisbilecki/go-uberfx-example/service"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(service.NewHTTPServer, routes.NewEchoRoute, handler.NewEchoHandler),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
