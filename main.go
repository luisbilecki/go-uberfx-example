package main

import (
	"net/http"

	"github.com/luisbilecki/go-uberfx-example/handler"
	routes "github.com/luisbilecki/go-uberfx-example/route"
	"github.com/luisbilecki/go-uberfx-example/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			service.NewHTTPServer,
			routes.NewEchoRoute,
			handler.NewEchoHandler,
			zap.NewProduction,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
