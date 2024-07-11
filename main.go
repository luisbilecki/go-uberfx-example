package main

import (
	"net/http"

	"github.com/luisbilecki/go-uberfx-example/service"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(service.NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
