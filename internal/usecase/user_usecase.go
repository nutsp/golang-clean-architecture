package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/nutsp/golang-clean-architecture/internal/models"
	"github.com/nutsp/golang-clean-architecture/internal/repositories"
	appError "github.com/nutsp/golang-clean-architecture/pkg/apperror"
	"github.com/nutsp/golang-clean-architecture/pkg/observability"
	"go.uber.org/dig"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUserInfo(ctx context.Context, user *models.User) error
	GetUserInfo(ctx context.Context, id uint) (*models.User, error)
}

type UserUsecase struct {
	logger           observability.Logger
	userRepository   repositories.IUserRepository
	mailerRepository repositories.IMailerRepository
}

type UserUsecaseDependencies struct {
	dig.In
	Logger           observability.Logger           `name:"Logger"`
	UserRepository   repositories.IUserRepository   `name:"UserRepository"`
	MailerRepository repositories.IMailerRepository `name:"MailerRepository"`
}

func NewUserUsecase(deps UserUsecaseDependencies) *UserUsecase {
	return &UserUsecase{
		logger:           deps.Logger,
		userRepository:   deps.UserRepository,
		mailerRepository: deps.MailerRepository,
	}
}

// CreateUser method creates a new user in the database.
// It checks for email availability and hashes the password before saving the user.
func (s *UserUsecase) CreateUser(ctx context.Context, user *models.User) error {
	// Call the third-party API to check email availability
	emailAvailable, err := s.mailerRepository.CheckEmailAvailability(ctx, user.Email)
	if err != nil {
		return appError.InternalServerError(err)
	}

	if !emailAvailable {
		return appError.InternalServerError(errors.New("email is already in use"))
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return appError.InternalServerError(err)
	}

	// Set the hashed password in the user object
	user.Password = string(hashedPassword)
	if err := s.userRepository.Save(ctx, user); err != nil {
		return appError.InternalServerError(err)
	}

	return nil
}

func (s *UserUsecase) UpdateUserInfo(ctx context.Context, user *models.User) error {
	err := s.userRepository.Atomic(ctx, nil, func(tx repositories.IUserRepository) error {
		users, err := tx.GetAll(ctx)
		if err != nil {
			return err
		}

		for _, u := range users {
			fmt.Println("User: ", u)
			u.Name = "Test Atomic"
			if err := tx.UpdateByID(ctx, u); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (s *UserUsecase) GetUserInfo(ctx context.Context, id uint) (*models.User, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, appError.InternalServerError(err)
	}
	return user, nil
}
