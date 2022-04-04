package repository

import (
	"context"
	"fmt"
	"github.com/DBoyara/invest-helper/pkg/models"
	"github.com/jackc/pgx/v4"
	"time"
)

const (
	selectedFieldsFutures string = "id, datetime, tiker, is_open, warranty_provision, count, amount, margin, commission"
)

func scanFutures(rows pgx.Row, model *models.Futures) error {
	return rows.Scan(
		&model.Id,
		&model.Datetime,
		&model.Tiker,
		&model.IsOpen,
		&model.WarrantyProvision,
		&model.Count,
		&model.Amount,
		&model.Margin,
		&model.Commission,
	)
}

func constructFuturesQuery(dateStart, dateEnd, showOpen string) (string, error) {
	selectQuery := fmt.Sprintf("select %s from futures", selectedFieldsFutures)
	filterQuery := fmt.Sprintf("where is_open = %s", showOpen)
	dateFilter := ""

	if dateStart != "" && dateEnd != "" {
		ds, err := time.Parse(layout, dateStart)
		if err != nil {
			return "", err
		}
		de, err := time.Parse(layout, dateEnd)
		if err != nil {
			return "", err
		}
		dateFilter = fmt.Sprintf(" and datetime between '%v' and '%v'", ds, de)
	}
	query := fmt.Sprintf("%s %s%s order by datetime desc", selectQuery, filterQuery, dateFilter)

	return query, nil
}

func CreateFuture(futures *models.Futures) (*models.Futures, error) {
	err := connPool.BeginFunc(context.Background(), func(t pgx.Tx) error {
		return t.QueryRow(
			context.Background(),
			`insert into futures 
				(tiker, is_open, warranty_provision, count, amount, commission) 
			values 
				($1, $2, $3, $4, $5, $6) 
			returning id`,
			futures.Tiker,
			futures.IsOpen,
			futures.WarrantyProvision,
			futures.Count,
			futures.Amount,
			futures.Commission,
		).Scan(&futures.Id)
	})

	return futures, err
}

func GetFutures(dateStart, dateEnd, showOpen string) ([]*models.Futures, error) {
	var res []*models.Futures

	query, err := constructFuturesQuery(dateStart, dateEnd, showOpen)
	if err != nil {
		return nil, err
	}

	rows, err := connPool.Query(
		context.Background(),
		query,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	for rows.Next() {
		tmp := &models.Futures{}

		if err = scanFutures(rows, tmp); err != nil {
			return res, err
		}

		res = append(res, tmp)
	}

	return res, nil
}

func GetSingleFutures(futuresId int64) (*models.Futures, error) {
	res := &models.Futures{}

	err := scanFutures(connPool.QueryRow(
		context.Background(),
		fmt.Sprintf("select %s from futures where id = $1", selectedFieldsFutures),
		futuresId,
	), res)

	if err == pgx.ErrNoRows {
		return res, nil
	}

	return res, err
}

func UpdateFutures(futures *models.Futures) error {
	query := fmt.Sprintf("update futures set margin = %f, is_open = %t, commission = %f where id = %d",
		futures.Margin,
		futures.IsOpen,
		futures.Commission,
		futures.Id)

	_, err := connPool.Exec(
		context.Background(),
		query,
	)

	if err == pgx.ErrNoRows {
		return nil
	}

	return err
}
