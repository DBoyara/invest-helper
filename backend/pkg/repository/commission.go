package repository

import (
	"context"
	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/jackc/pgx/v4"
)

func GetCommissions() ([]*models.Commission, error) {
	var res []*models.Commission

	rows, err := connPool.Query(
		context.Background(),
		"select value, type from commissions order by 1",
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	for rows.Next() {
		tmp := &models.Commission{}

		if err = rows.Scan(
			&tmp.Value,
			&tmp.Type); err != nil {
			return res, err
		}

		res = append(res, tmp)
	}

	return res, nil
}
