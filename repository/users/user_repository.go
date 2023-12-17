package users

import (
	"context"

	"github.com/dhimweray222/test-BE-uninet/model/domain"
)

type UserRepository interface {
	CreateUserTx(ctx context.Context, user domain.User) error
	FindUserByQueryTx(ctx context.Context, query, value string) (domain.User, error)
	FindAllNotDeletedUserTx(ctx context.Context) ([]domain.User, error)
	LoginTx(ctx context.Context, id string) (domain.User, error)
}
