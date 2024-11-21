package repositories

import (
	"github.com/duongbui2002/core-package/utils"
	"github.com/duongbui2002/myblog-authservice/internal/auth/models"
	"golang.org/x/net/context"
)

type UserRepository interface {
	GetAllUsers(
		ctx context.Context,
		listQuery *utils.ListQuery,
	) (*utils.ListResult[*models.User], error)
	SearchUsers(
		ctx context.Context,
		searchText string,
		listQuery *utils.ListQuery,
	) (*utils.ListResult[*models.User], error)
	GetUserById(ctx context.Context, uuid string) (*models.User, error)
	CreateUser(ctx context.Context, product *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, product *models.User) (*models.User, error)
	DeleteUserByID(ctx context.Context, uuid string) error
}
