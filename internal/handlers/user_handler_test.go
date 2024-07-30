package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/nutsp/golang-clean-architecture/internal/handlers"
	"github.com/nutsp/golang-clean-architecture/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	mockUsecase *mocks.MockIUserUsecase
	handler     *handlers.UserHandler
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockUsecase = mocks.NewMockIUserUsecase(s.ctrl)

	handlerDeps := handlers.UserHandlerDependencies{
		UserUsecase: s.mockUsecase,
	}

	s.handler = handlers.NewUserHandler(handlerDeps)
}

func (s *UserHandlerTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) TestCreateUserHandler() {
	e := echo.New()

	tests := []struct {
		name          string
		requestBody   interface{}
		expectedCode  int
		expectedError bool
		setupMocks    func()
	}{
		{
			name: "success",
			requestBody: map[string]string{
				"name":     "John Doe",
				"email":    "john.doe@example.com",
				"password": "password123",
			},
			expectedCode:  http.StatusOK,
			expectedError: false,
			setupMocks: func() {
				s.mockUsecase.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		// {
		// 	name:          "bad_request",
		// 	requestBody:   "invalid body",
		// 	expectedCode:  http.StatusBadRequest,
		// 	expectedError: true,
		// 	setupMocks:    func() {},
		// },
		// {
		// 	name: "internal_server_error",
		// 	requestBody: map[string]string{
		// 		"name":     "John Doe",
		// 		"email":    "john.doe@example.com",
		// 		"password": "password123",
		// 	},
		// 	expectedCode:  http.StatusInternalServerError,
		// 	expectedError: true,
		// 	setupMocks: func() {
		// 		s.mockUsecase.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
		// 	},
		// },
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := s.handler.CreateUserHandler(c)
			if tt.expectedError {
				assert.Error(s.T(), err)
			} else {
				assert.NoError(s.T(), err)
			}
			assert.Equal(s.T(), tt.expectedCode, rec.Code)
		})
	}
}

func (s *UserHandlerTestSuite) TestUpdateUserHandler() {
	e := echo.New()

	tests := []struct {
		name          string
		requestBody   interface{}
		expectedCode  int
		expectedError bool
		setupMocks    func()
	}{
		// {
		// 	name: "success",
		// 	requestBody: map[string]string{
		// 		"id":       "1",
		// 		"name":     "John Doe",
		// 		"email":    "john.doe@example.com",
		// 		"password": "password123",
		// 	},
		// 	expectedCode:  http.StatusOK,
		// 	expectedError: false,
		// 	setupMocks: func() {
		// 		s.mockUsecase.EXPECT().UpdateUserInfo(gomock.Any(), gomock.Any()).Return(nil)
		// 	},
		// },
		// {
		// 	name:          "bad_request",
		// 	requestBody:   "invalid body",
		// 	expectedCode:  http.StatusBadRequest,
		// 	expectedError: true,
		// 	setupMocks:    func() {},
		// },
		// {
		// 	name: "internal_server_error",
		// 	requestBody: map[string]string{
		// 		"id":       "1",
		// 		"name":     "John Doe",
		// 		"email":    "john.doe@example.com",
		// 		"password": "password123",
		// 	},
		// 	expectedCode:  http.StatusInternalServerError,
		// 	expectedError: true,
		// 	setupMocks: func() {
		// 		s.mockUsecase.EXPECT().UpdateUserInfo(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
		// 	},
		// },
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := s.handler.UpdateUserHandler(c)
			if tt.expectedError {
				assert.Error(s.T(), err)
			} else {
				assert.NoError(s.T(), err)
			}
			assert.Equal(s.T(), tt.expectedCode, rec.Code)
		})
	}
}
