package repositories

import (
	"context"
	"emperror.dev/errors"
	"fmt"
	"github.com/duongbui2002/core-package/core/data"
	"github.com/duongbui2002/core-package/logger"
	"github.com/duongbui2002/core-package/mongodb"
	"github.com/duongbui2002/core-package/mongodb/repository"
	"github.com/duongbui2002/core-package/utils"
	"github.com/duongbui2002/myblog-authservice/internal/auth/contracts/repositories"
	"github.com/duongbui2002/myblog-authservice/internal/auth/models"
	uuid2 "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollection = "users"
)

type mongoUserRepository struct {
	log                    logger.Logger
	mongoGenericRepository data.GenericRepository[*models.User]
}

func NewMongoUserRepository(
	log logger.Logger,
	db *mongo.Client,
	mongoOptions *mongodb.MongoDbOptions,
) repositories.UserRepository {
	mongoRepo := repository.NewGenericMongoRepository[*models.User](
		db,
		mongoOptions.Database,
		userCollection,
	)
	return &mongoUserRepository{
		log:                    log,
		mongoGenericRepository: mongoRepo,
	}
}

func (p *mongoUserRepository) GetAllUsers(
	ctx context.Context,
	listQuery *utils.ListQuery,
) (*utils.ListResult[*models.User], error) {
	//ctx, span := p.tracer.Start(ctx, "mongoUserRepository.GetAllUsers")
	//defer span.End()

	// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/
	result, err := p.mongoGenericRepository.GetAll(ctx, listQuery)
	if err != nil {
		return nil, errors.WrapIf(
			err,
			"error in the paginate",
		)
	}

	p.log.Infow(
		"users loaded",
		logger.Fields{"UsersResult": result},
	)

	//span.SetAttributes(attribute.Object("UsersResult", result))

	return result, nil
}

func (p *mongoUserRepository) SearchUsers(
	ctx context.Context,
	searchText string,
	listQuery *utils.ListQuery,
) (*utils.ListResult[*models.User], error) {
	//ctx, span := p.tracer.Start(ctx, "mongoUserRepository.SearchUsers")
	//span.SetAttributes(attribute2.String("SearchText", searchText))
	//defer span.End()

	result, err := p.mongoGenericRepository.Search(ctx, searchText, listQuery)
	if err != nil {
		return nil,
			errors.WrapIf(
				err,
				"error in the paginate",
			)

	}

	p.log.Infow(
		fmt.Sprintf(
			"users loaded for search term '%s'",
			searchText,
		),
		logger.Fields{"UsersResult": result},
	)

	//span.SetAttributes(attribute.Object("UsersResult", result))

	return result, nil
}

func (p *mongoUserRepository) GetUserById(
	ctx context.Context,
	uuid string,
) (*models.User, error) {
	//ctx, span := p.tracer.Start(ctx, "mongoUserRepository.GetUserById")
	//span.SetAttributes(attribute2.String("Id", uuid))
	//defer span.End()

	id, err := uuid2.FromString(uuid)
	if err != nil {
		return nil, err
	}

	user, err := p.mongoGenericRepository.GetById(ctx, id)
	if err != nil {
		return nil,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"can't find the user with id %s into the database.",
					uuid,
				),
			)
	}

	//span.SetAttributes(attribute.Object("User", user))

	p.log.Infow(
		fmt.Sprintf("user with id %s laoded", uuid),
		logger.Fields{"User": user, "Id": uuid},
	)

	return user, nil
}

func (p *mongoUserRepository) GetUserByID(
	ctx context.Context,
	uuid string,
) (*models.User, error) {
	userId := uuid
	//ctx, span := p.tracer.Start(
	//	ctx,
	//	"mongoUserRepository.GetUserByID",
	//)
	//span.SetAttributes(attribute2.String("ID", userId))
	//defer span.End()

	user, err := p.mongoGenericRepository.FirstOrDefault(
		ctx,
		map[string]interface{}{"userId": uuid},
	)
	if err != nil {
		return nil,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"can't find the user with userId %s into the database.",
					uuid,
				),
			)
	}

	//span.SetAttributes(attribute.Object("User", user))

	p.log.Infow(
		fmt.Sprintf(
			"user with userId %s laoded",
			userId,
		),
		logger.Fields{"User": user, "ID": uuid},
	)

	return user, nil
}

func (p *mongoUserRepository) CreateUser(
	ctx context.Context,
	user *models.User,
) (*models.User, error) {
	//ctx, span := p.tracer.Start(ctx, "mongoUserRepository.CreateUser")
	//defer span.End()

	err := p.mongoGenericRepository.Add(ctx, user)
	if err != nil {
		return nil,
			errors.WrapIf(
				err,
				"error in the inserting user into the database.",
			)

	}

	//span.SetAttributes(attribute.Object("User", user))

	p.log.Infow(
		fmt.Sprintf(
			"user with id '%s' created",
			user.ID,
		),
		logger.Fields{"User": user, "Id": user.ID},
	)

	return user, nil
}

func (p *mongoUserRepository) UpdateUser(
	ctx context.Context,
	updateUser *models.User,
) (*models.User, error) {
	//ctx, span := p.tracer.Start(ctx, "mongoUserRepository.UpdateUser")
	//defer span.End()
	//
	err := p.mongoGenericRepository.Update(ctx, updateUser)
	// https://www.mongodb.com/docs/manual/reference/method/db.collection.findOneAndUpdate/
	if err != nil {
		return nil,
			errors.WrapIf(
				err,
				fmt.Sprintf(
					"error in updating user with id %s into the database.",
					updateUser.ID,
				),
			)
	}

	//span.SetAttributes(attribute.Object("User", updateUser))
	p.log.Infow(
		fmt.Sprintf(
			"user with id '%s' updated",
			updateUser.ID,
		),
		logger.Fields{"User": updateUser, "Id": updateUser.ID},
	)

	return updateUser, nil
}

func (p *mongoUserRepository) DeleteUserByID(
	ctx context.Context,
	uuid string,
) error {
	//ctx, span := p.tracer.Start(ctx, "mongoUserRepository.DeleteUserByID")
	//span.SetAttributes(attribute2.String("Id", uuid))
	//defer span.End()

	id, err := uuid2.FromString(uuid)
	if err != nil {
		return err
	}

	err = p.mongoGenericRepository.Delete(ctx, id)
	if err != nil {
		return errors.WrapIf(err, fmt.Sprintf(
			"error in deleting user with id %s from the database.",
			uuid,
		))

	}

	p.log.Infow(
		fmt.Sprintf("user with id %s deleted", uuid),
		logger.Fields{"User": uuid},
	)

	return nil
}
