package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nutsp/golang-clean-architecture/config"
	"github.com/nutsp/golang-clean-architecture/internal/handlers"
	"github.com/nutsp/golang-clean-architecture/internal/middlewares"
	"github.com/nutsp/golang-clean-architecture/pkg/observability"
	"go.uber.org/dig"
)

type App struct {
	echo        *echo.Echo
	config      *config.Config
	logger      observability.Logger
	middleware  middlewares.IMiddleware
	userHandler handlers.IUserHandler
}

type AppDependencies struct {
	dig.In
	Config      *config.Config
	Logger      observability.Logger    `name:"Logger"`
	Middleware  middlewares.IMiddleware `name:"Middleware"`
	UserHandler handlers.IUserHandler   `name:"UserHandler"`
}

func NewApp(deps AppDependencies) {
	app := &App{
		echo:        middlewares.NewEchoServer(deps.Config, deps.Middleware),
		config:      deps.Config,
		middleware:  deps.Middleware,
		userHandler: deps.UserHandler,
	}
	app.Start()
}

func (app *App) Start() error {
	app.InitRoute()

	return app.Run()
}

func (app *App) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	go func() {
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		app.echo.Shutdown(ctx)
	}()

	return app.echo.Start(fmt.Sprintf(":%s", app.config.Server.Port))
}
