package repository

import (
	"context"
	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/jackc/pgx/v4"
)

func GetCommissions() (*models.Commission, error) {
	var commission *models.Commission

	err := connPool.QueryRow(
		context.Background(),
		"select value, type from commissions",
	).Scan(&commission.Value, &commission.Type)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return commission, nil
}
