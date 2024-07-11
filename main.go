package main

import (
	"net/http"

	"github.com/luisbilecki/go-uberfx-example/handler"
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
			fx.Annotate(
				service.NewServeMux,
				fx.ParamTags(`name:"echo"`, `name:"hello"`),
			),
			fx.Annotate(
				handler.NewEchoHandler,
				fx.As(new(handler.Route)),
				fx.ResultTags(`name:"echo"`),
			),
			fx.Annotate(
				handler.NewHelloHandler,
				fx.As(new(handler.Route)),
				fx.ResultTags(`name:"hello"`),
			),
			zap.NewProduction,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
