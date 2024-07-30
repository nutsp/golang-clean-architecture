package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nutsp/golang-clean-architecture/internal/models"
	"github.com/nutsp/golang-clean-architecture/internal/usecase"
	appError "github.com/nutsp/golang-clean-architecture/pkg/apperror"
	"github.com/nutsp/golang-clean-architecture/pkg/response"
	"go.uber.org/dig"
)

type IUserHandler interface {
	CreateUserHandler(c echo.Context) error
	UpdateUserHandler(c echo.Context) error
	GetUserInfoHandler(c echo.Context) error
}

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

type UserHandlerDependencies struct {
	dig.In
	UserUsecase usecase.IUserUsecase `name:"UserUsecase"`
}

func NewUserHandler(deps UserHandlerDependencies) *UserHandler {
	return &UserHandler{
		userUsecase: deps.UserUsecase,
	}
}

func (h *UserHandler) CreateUserHandler(c echo.Context) error {
	// Get user name from request
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return response.ErrorBuilder(appError.BadRequest(err)).Send(c)
	}

	// Call user use case method
	err := h.userUsecase.CreateUser(c.Request().Context(), user)

	// Return response
	if err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(nil).Send(c)
}

func (h *UserHandler) UpdateUserHandler(c echo.Context) error {
	// Get user name from request
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return response.ErrorBuilder(appError.BadRequest(err)).Send(c)
	}

	// Call user use case method
	err := h.userUsecase.UpdateUserInfo(c.Request().Context(), user)

	// Return response
	if err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(nil).Send(c)
}

func (h *UserHandler) GetUserInfoHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.ErrorBuilder(appError.BadRequest(err)).Send(c)
	}

	user, err := h.userUsecase.GetUserInfo(c.Request().Context(), uint(id))
	if err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(user).Send(c)
}
