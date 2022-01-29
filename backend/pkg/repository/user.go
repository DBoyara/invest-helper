package repository

import (
	"context"

	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/jackc/pgx/v4"
)

func CreateUser(user *models.User) (int, error) {
	err := connPool.BeginFunc(context.Background(), func(t pgx.Tx) error {
		return t.QueryRow(
			context.Background(),
			`insert into instruments 
				(username, password) 
			values 
				($1, $2) 
			returning id`,
			user.Username,
			user.Password,
		).Scan(&user.ID)
	})

	return user.ID, err
}

func GetUser() (*models.User, error) {
	var user *models.User

	err := connPool.QueryRow(
		context.Background(),
		"select id, username, password from users limit 1",
	).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
