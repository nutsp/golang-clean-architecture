package repositories_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_database "github.com/nutsp/golang-clean-architecture/internal/infastructure/database/mock"
	"github.com/nutsp/golang-clean-architecture/internal/models"
	"github.com/nutsp/golang-clean-architecture/internal/repositories"
	mock_datasource "github.com/nutsp/golang-clean-architecture/pkg/datasource/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	ctrl           *gomock.Controller
	mockDB         *mock_database.MockIDatabase
	mockConn       *mock_datasource.MockDB
	userRepository *repositories.UserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockDB = mock_database.NewMockIDatabase(s.ctrl)
	s.mockConn = mock_datasource.NewMockDB(s.ctrl)

	// Set up expectations for GORM methods
	ctx := context.Background()
	s.mockDB.EXPECT().OllamaDB().Return(s.mockConn).AnyTimes()
	s.mockConn.EXPECT().Debug().Return(s.mockConn).AnyTimes() // Expect Debug to be called
	s.mockConn.EXPECT().WithContext(ctx).Return(s.mockConn)   // Expect WithContext to be called

	userRepoDeps := repositories.UserRepositoryDependencies{
		DB: s.mockDB,
	}
	s.userRepository = repositories.NewUserRepository(userRepoDeps)
}

func (s *UserRepositoryTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

// Happy Case: Save succeeds
func (s *UserRepositoryTestSuite) TestUserRepositorySaveThenSuccess() {
	ctx := context.Background()

	s.Run("success", func() {
		user := &models.User{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
		}

		// Set up expectations for GORM methods
		s.mockConn.EXPECT().Create(user).Return(s.mockConn) // Expect Create to be called
		s.mockConn.EXPECT().Error().Return(nil)             // Expect Error to return nil

		// Call the Save method
		err := s.userRepository.Save(ctx, user)

		// Assert no error is returned
		assert.Nil(s.T(), err)
	})
}

// Fail Case: Save fails due to an error
func (s *UserRepositoryTestSuite) TestUserRepositorySaveThenFail() {
	ctx := context.Background()

	s.Run("failure_when_database_error", func() {
		user := &models.User{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
		}

		// Set up expectations for GORM methods
		s.mockConn.EXPECT().Create(user).Return(s.mockConn)              // Expect Create to be called
		s.mockConn.EXPECT().Error().Return(errors.New("database error")) // Expect Error to return an error

		// Call the Save method
		err := s.userRepository.Save(ctx, user)

		// Assert the expected error is returned
		assert.NotNil(s.T(), err)
		assert.Equal(s.T(), "database error", err.Error())
	})
}
