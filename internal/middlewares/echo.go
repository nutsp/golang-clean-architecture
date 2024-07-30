package middlewares

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/nutsp/golang-clean-architecture/config"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (v *CustomValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

var configCors = echoMiddleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
}

func NewEchoServer(cfg *config.Config, mw IMiddleware) *echo.Echo {
	e := echo.New()

	e.Use(echoMiddleware.CORSWithConfig(configCors))
	e.Use(mw.LoggingMiddleware())

	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = errorHandler
	e.Debug = cfg.Server.Debug
	e.HideBanner = true

	return e
}
