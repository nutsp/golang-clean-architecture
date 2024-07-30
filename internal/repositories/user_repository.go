package repositories

import (
	"context"
	"database/sql"

	"github.com/nutsp/golang-clean-architecture/internal/infastructure/database"
	"github.com/nutsp/golang-clean-architecture/internal/models"
	"github.com/nutsp/golang-clean-architecture/pkg/datasource"
	"go.uber.org/dig"
)

type IUserRepository interface {
	Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx IUserRepository) error) error
	Save(ctx context.Context, user *models.User) error
	UpdateByID(ctx context.Context, user *models.User) error
	GetAll(ctx context.Context) ([]*models.User, error)
	GetByID(ctx context.Context, id uint) (*models.User, error)
}

type UserRepository struct {
	db   database.IDatabase
	conn datasource.DB
}

type UserRepositoryDependencies struct {
	dig.In
	DB database.IDatabase `name:"Database"`
}

func NewUserRepository(deps UserRepositoryDependencies) *UserRepository {
	return &UserRepository{db: deps.DB, conn: deps.DB.OllamaDB()}
}

// Atomic implements Repository Interface for transaction query
func (r *UserRepository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx IUserRepository) error) error {
	tx := r.db.OllamaDB().Begin()
	if tx.Error() != nil {
		return tx.Error()
	}

	newRepository := &UserRepository{db: r.db, conn: tx}
	err := repo(newRepository)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Save(ctx context.Context, user *models.User) error {
	return r.conn.Debug().WithContext(ctx).Create(user).Error()
}

func (r *UserRepository) UpdateByID(ctx context.Context, user *models.User) error {
	return r.conn.Debug().WithContext(ctx).Where("id =?", user.ID).Updates(user).Error()
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := r.conn.Debug().WithContext(ctx).Find(&users).Error()
	if err != nil {
		return nil, err
	}
	return users, err
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user *models.User
	err := r.conn.Debug().WithContext(ctx).Where("id =?", id).First(&user).Error()
	if err != nil {
		return nil, err
	}

	return user, nil
}
