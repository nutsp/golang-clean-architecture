package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nutsp/golang-clean-architecture/internal/models"
	"github.com/nutsp/golang-clean-architecture/pkg/datasource"
	"go.uber.org/dig"
)

type IUserRedisRepository interface {
	SetUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id uint) (*models.User, error)
}

type UserRedisRepository struct {
	client datasource.IRedisClient
}

type UserRedisRepositoryDependencies struct {
	dig.In
	Client datasource.IRedisClient `name:"RedisClient"`
}

func NewUserRedisRepository(deps UserRedisRepositoryDependencies) *UserRedisRepository {
	return &UserRedisRepository{
		client: deps.Client,
	}
}

func (r *UserRedisRepository) SetUser(ctx context.Context, user *models.User) error {
	key := r.client.GetKeyName("users", fmt.Sprint(user.ID))
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, val, 15*time.Minute)
}

func (r *UserRedisRepository) GetUser(ctx context.Context, id uint) (*models.User, error) {
	var user *models.User
	key := r.client.GetKeyName("users", fmt.Sprint(id))
	userStr, err := r.client.Get(ctx, key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	if err := json.Unmarshal([]byte(userStr), &user); err != nil {
		return nil, err
	}

	return user, nil
}
