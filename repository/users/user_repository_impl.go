package users

import (
	"context"

	"github.com/dhimweray222/test-BE-uninet/model/domain"
	"github.com/dhimweray222/test-BE-uninet/repository"
	"github.com/jackc/pgx/v5"
)

type UserRepositoryImpl struct {
	DB repository.Store
}

func NewUserRepository(db repository.Store) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) CreateUserTx(ctx context.Context, user domain.User) error {
	var err error
	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {

		err := repository.CreateUser(ctx, tx, user)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (repository *UserRepositoryImpl) FindUserByQueryTx(ctx context.Context, query, value string) (domain.User, error) {

	var data domain.User
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {

		data, err = repository.FindUserByQuery(ctx, tx, query, value)
		if err != nil {
			return err
		}

		return nil
	})

	return data, err
}

func (repository *UserRepositoryImpl) FindAllNotDeletedUserTx(ctx context.Context) ([]domain.User, error) {

	var data []domain.User
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {

		data, err = repository.FindAll(ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})

	return data, err
}

func (repository *UserRepositoryImpl) LoginTx(ctx context.Context, id string) (domain.User, error) {

	var data domain.User
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {

		data, err = repository.Login(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	})

	return data, err
}
