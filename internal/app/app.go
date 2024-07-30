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
	"go.uber.org/dig"
)

type App struct {
	echo        *echo.Echo
	config      *config.Config
	userHandler handlers.IUserHandler
}

type AppDependencies struct {
	dig.In
	Config      *config.Config
	UserHandler handlers.IUserHandler `name:"UserHandler"`
}

func NewApp(deps AppDependencies) {
	fmt.Println("NewApp")
	app := &App{
		echo:        middlewares.NewEchoServer(deps.Config),
		config:      deps.Config,
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
