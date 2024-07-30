package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/nutsp/golang-clean-architecture/pkg/observability"
	"go.uber.org/dig"
)

type IMiddleware interface {
	LoggingMiddleware() echo.MiddlewareFunc
}

type Middleware struct {
	logger observability.Logger
}

type MiddlewareDependencies struct {
	dig.In
	Logger observability.Logger `name:"Logger"`
}

func NewMiddleware(deps MiddlewareDependencies) *Middleware {
	return &Middleware{
		logger: deps.Logger,
	}
}
