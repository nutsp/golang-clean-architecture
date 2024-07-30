package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nutsp/golang-clean-architecture/internal/mocks"
	"github.com/nutsp/golang-clean-architecture/internal/models"
	"github.com/nutsp/golang-clean-architecture/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	ctrl           *gomock.Controller
	mockUserRepo   *mocks.MockIUserRepository
	mockMailerRepo *mocks.MockIMailerRepository
	userService    *usecase.UserUsecase
}

func (s *UserServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.mockUserRepo = mocks.NewMockIUserRepository(s.ctrl)
	s.mockMailerRepo = mocks.NewMockIMailerRepository(s.ctrl)

	userDeps := usecase.UserUsecaseDependencies{
		UserRepository:   s.mockUserRepo,
		MailerRepository: s.mockMailerRepo,
	}

	s.userService = usecase.NewUserUsecase(userDeps)
}

func (s *UserServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (s *UserServiceTestSuite) TestCreateUser() {
	user := &models.User{
		Name:     "John Doe",
		Email:    "test@example.com",
		Password: "password",
	}
	ctx := context.Background()

	tests := []struct {
		name           string
		emailAvailable bool
		saveErr        error
		expectedErr    error
	}{
		{
			name:           "success_case_create_user",
			emailAvailable: true,
			saveErr:        nil,
			expectedErr:    nil,
		},
		{
			name:           "failure_case_create_user_email_not_available",
			emailAvailable: false,
			saveErr:        nil,
			expectedErr:    errors.New("email is already in use"),
		},
		{
			name:           "failure_case_save_user_error",
			emailAvailable: true,
			saveErr:        errors.New("save user error"),
			expectedErr:    errors.New("save user error"),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockMailerRepo.EXPECT().CheckEmailAvailability(ctx, user.Email).Return(tt.emailAvailable, nil)
			if tt.emailAvailable {
				s.mockUserRepo.EXPECT().Save(ctx, user).Return(tt.saveErr)
			}

			err := s.userService.CreateUser(ctx, user)

			if tt.expectedErr != nil {
				assert.EqualError(s.T(), err, tt.expectedErr.Error())
			} else {
				assert.NoError(s.T(), err)
			}
		})
	}
}
func BenchmarkCreateUser(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepository(ctrl)
	mockMailerRepo := mocks.NewMockIMailerRepository(ctrl)

	userDeps := usecase.UserUsecaseDependencies{
		UserRepository:   mockUserRepo,
		MailerRepository: mockMailerRepo,
	}

	userService := usecase.NewUserUsecase(userDeps)

	user := &models.User{
		Name:     "John Doe",
		Email:    "test@example.com",
		Password: "password",
	}
	ctx := context.Background()

	// Setup the mock expectations
	mockMailerRepo.EXPECT().CheckEmailAvailability(ctx, user.Email).Return(true, nil).AnyTimes()
	mockUserRepo.EXPECT().Save(ctx, user).Return(nil).AnyTimes()

	// Reset the timer before running the benchmark
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Call the CreateUser method
		err := userService.CreateUser(ctx, user)
		if err != nil {
			b.Fatalf("CreateUser failed: %v", err)
		}
	}
}
