package users

import (
	"context"
	"fmt"
	"log"

	"github.com/dhimweray222/test-BE-uninet/model/domain"
	"github.com/jackc/pgx/v5"
)

func (repository *UserRepositoryImpl) CreateUser(ctx context.Context, db pgx.Tx, user domain.User) error {
	query := fmt.Sprintf(`
	INSERT INTO %s (
		id,
		name,
		email,
		password,
		created_at)
	VALUES($1,$2,$3,$4,$5)`, "users")

	if _, err := db.Prepare(context.Background(), "create_user", query); err != nil {
		log.Println(err, "ini")
		return err
	}

	if _, err := db.Exec(context.Background(), "create_user",
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
	); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) FindUserByQuery(ctx context.Context, db pgx.Tx, query, value string) (domain.User, error) {
	queryStr := fmt.Sprintf(`SELECT id, name, email, password, created_at  FROM %s 
	WHERE %s = $1`, "users", query)

	user, err := db.Query(context.Background(), queryStr, value)

	if err != nil {
		log.Println("log err:", err)
		return domain.User{}, err
	}

	defer user.Close()

	data, err := pgx.CollectOneRow(user, pgx.RowToStructByPos[domain.User])

	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	return data, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, db pgx.Tx) ([]domain.User, error) {
	queryStr := fmt.Sprintf(`SELECT id, name, email, password, phone, r.name, created_at, updated_at  FROM %s`, "users")

	user, err := db.Query(context.Background(), queryStr)
	if err != nil {
		log.Println(err)
		return []domain.User{}, err
	}

	defer user.Close()
	data, err := pgx.CollectRows(user, pgx.RowToStructByPos[domain.User])
	if err != nil {
		log.Println(data, err)
		return []domain.User{}, err
	}

	return data, nil
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, db pgx.Tx, nik string) (domain.User, error) {
	queryStr := fmt.Sprintf("SELECT * FROM %s WHERE email = $1 ", "users")

	user, err := db.Query(context.Background(), queryStr, nik)
	log.Println(queryStr)
	if err != nil {
		return domain.User{}, err
	}

	defer user.Close()

	data, err := pgx.CollectOneRow(user, pgx.RowToStructByPos[domain.User])
	log.Println(data)
	if err != nil {
		return domain.User{}, err
	}

	return data, nil
}
