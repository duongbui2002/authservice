package repositories

import (
	"context"
	"github.com/duongbui2002/myblog-authservice/internal/auth/models"
)

type UserCacheRepository interface {
	PutUser(ctx context.Context, key string, product *models.User) error
	GetUserById(ctx context.Context, key string) (*models.User, error)
	DeleteUser(ctx context.Context, key string) error
	DeleteAllUsers(ctx context.Context) error
}
