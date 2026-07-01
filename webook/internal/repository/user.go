package repository

import (
	"context"
	"rewebook/internal/domain"
	"rewebook/internal/repository/dao"

	"github.com/redis/go-redis/v9"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao         *dao.UserDAO
	redisClient redis.Cmdable
}

func NewUserRepository(dao *dao.UserDAO, redisClient redis.Cmdable) *UserRepository {
	return &UserRepository{
		dao:         dao,
		redisClient: redisClient,
	}
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
func (r *UserRepository) FindById(int64) {}
